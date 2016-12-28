$(GOPATH)/bin/kraken-tick: $(wildcard *.go)
	go build -o $(GOPATH)/bin/kraken-tick main.go
	$(GOPATH)/bin/kraken-tick --help-long | sed 's/^/    /' > README.md

clean:
	rm -f $(GOPATH)/bin/kraken-tick
