all: mjolnir

mjolnir: *.go
	go build -v
	go install mjolnir/midgard
	go install mjolnir/helheim
	go install

.PHONY: clean
clean:
	go clean
	go clean helheim
	go clean midgard
