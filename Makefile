$(GOPATH)/bin/kraken: $(wildcard *.go)
	go build -o $(GOPATH)/bin/kraken main.go
	$(GOPATH)/bin/kraken --help-long > README.md

clean:
	rm -f $(GOPATH)/bin/kraken
