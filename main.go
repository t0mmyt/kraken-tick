package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/fatih/color"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	debug    = kingpin.Flag("debug", "Enable debug logs").Bool()
	tickRoot = kingpin.Flag("tickerRoot", "URL Endpoint for ticker").Default("https://api.kraken.com/0/public/Ticker").String()
	pair     = kingpin.Flag("pair", "Currency pair to get").Default("XXBTZEUR").String()
	format   = kingpin.Flag("format", "Format for number").Default("%.3f\n").String()
)

type tick struct {
	Error  []string              `json:"error"`
	Result map[string]tickResult `json:"result"`
}

type tickResult struct {
	Ask     []string `json:"a"`
	Bid     []string `json:"b"`
	Last    []string `json:"c"`
	Volume  []string `json:"v"`
	VolWA   []string `json:"p"`
	Trades  []int64  `json:"t"`
	Low     []string `json:"l"`
	High    []string `json:"h"`
	Opening string   `json:"o"`
}

func getTicker(tickRoot, pair string) (tick, error) {
	var t tick

	req, _ := http.NewRequest("GET", tickRoot, nil)
	q := req.URL.Query()
	q.Add("pair", pair)
	req.URL.RawQuery = q.Encode()
	log.Debugf("Built URL: %s", req.URL.String())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return t, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return t, err
	}
	err = json.Unmarshal(b, &t)
	if err != nil {
		return t, err
	}
	if len(t.Error) > 0 {
		return t, fmt.Errorf("Errors in response: %+v", t.Error)
	}
	return t, nil
}

func main() {
	kingpin.Parse()
	if *debug {
		log.SetLevel(log.DebugLevel)
	}

	t, err := getTicker(*tickRoot, *pair)
	if err != nil {
		log.Fatalf("%s", err)
	}
	last, err := strconv.ParseFloat(t.Result[*pair].Last[0], 64)
	if err != nil {
		log.Fatalf("%s", err)
	}
	color.Yellow(fmt.Sprintf(*format, last))
}
