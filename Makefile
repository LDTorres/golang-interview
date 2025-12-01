.PHONY: run test clean

run:
	go run .

test:
	go test -v ./...

clean:
	go clean
