.PHONY: test coverage run

test:
	go test ./... -tags=unit

coverage:
	go test ./... -tags=unit -coverprofile=coverage.out
	go tool cover -html=coverage.out

run:
	go run ./... $(args)