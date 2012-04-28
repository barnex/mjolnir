all: mjolnir

mjolnir: *.go
	go install mjolnir/midgard
	go install mjolnir/helheim
	go install

.PHONY: clean
clean:
	go clean
	rm -rf $(GOPATH)/pkg/*/mjolnir
