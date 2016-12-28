$(GOPATH)/bin/kraken: $(wildcard *.go)
	go build -o $(GOPATH)/bin/kraken main.go
	$(GOPATH)/bin/kraken --help-long | sed 's/^/    /' > README.md

clean:
	rm -f $(GOPATH)/bin/kraken
