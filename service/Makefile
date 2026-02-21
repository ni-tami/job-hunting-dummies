# Configuration
PROTO_BASE_DIR   := pb
OUT_BASE_DIR     := pb/out
GO_OUT_DIR       := $(OUT_BASE_DIR)/go

# Find all .proto files recursively
PROTO_FILES := $(shell find $(PROTO_BASE_DIR) -name "*.proto")

# Derive output directories from proto paths
PROTO_DIRS := $(sort $(dir $(PROTO_FILES)))

## Generate all proto files
proto: proto-go

## Generate Go gRPC code for all services
proto-go: $(PROTO_FILES)
	@echo "Generating Go gRPC code..."
	@for dir in $(PROTO_DIRS); do \
		echo "  Processing: $$dir"; \
		protoc \
			--proto_path=$(PROTO_BASE_DIR) \
			--go_out=$(GO_OUT_DIR) \
			--go_opt=paths=source_relative \
			--go-grpc_out=$(GO_OUT_DIR) \
			--go-grpc_opt=paths=source_relative \
			$$dir/*.proto; \
	done
	@echo "Done."

## Generate code for a specific use case: make proto-svc SVC=users
proto-svc:
ifndef SVC
	$(error SVC is not set. Usage: make proto-svc SVC=<use_case_name>)
endif
	@echo "Generating code for service: $(SVC)"
	@protoc \
		--proto_path=$(PROTO_BASE_DIR) \
		--go_out=$(GO_OUT_DIR) \
		--go_opt=paths=source_relative \
		--go-grpc_out=$(GO_OUT_DIR) \
		--go-grpc_opt=paths=source_relative \
		$(PROTO_BASE_DIR)/$(SVC)/*.proto
	@echo "Done."

## Clean generated files
clean:
	rm -rf $(OUT_BASE_DIR)

## Show help
help:
	@grep -E '^##' Makefile | sed 's/## //'

gqlgen:
	@echo "Generating GraphQL code..."
	@go tool gqlgen generate
	@echo "Done."

runserver:
	go run server.go

rundev:
	go tool air

.PHONY: proto proto-go clean help gqlgen
