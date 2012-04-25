all: mjolnir

mjolnir: *.go
	go tool 6g -o mjolnir.6 *.go
	go tool 6l -o mjolnir mjolnir.6
	go install

.PHONY: clean
clean:
	go clean
