runworker:
	go run worker/main.go

runstarter:
	go run starter/main.go

format:
	golangci-lint fmt

lint: format
	golangci-lint run

.PHONY: runworker runstarter format lint
