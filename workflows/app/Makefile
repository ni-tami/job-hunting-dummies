# Configuration
PROTO_BASE_DIR   := pb
OUT_BASE_DIR     := pb/out
GO_OUT_DIR       := $(OUT_BASE_DIR)/go
PYTHON_OUT_DIR	 := $(OUT_BASE_DIR)/python

# Find all .proto files recursively
PROTO_FILES := $(shell find $(PROTO_BASE_DIR) -name "*.proto")

# Derive output directories from proto paths
PROTO_DIRS := $(sort $(dir $(PROTO_FILES)))

## Generate all proto files
proto: proto-go proto-py

## Generate Go gRPC code for all services
proto-go: $(PROTO_FILES)
	@echo "Generating Go gRPC code..."
	@protoc \
		--proto_path=$(PROTO_BASE_DIR) \
		--go_out=$(GO_OUT_DIR) \
		--go_opt=paths=source_relative \
		--go-grpc_out=$(GO_OUT_DIR) \
		--go-grpc_opt=paths=source_relative \
		$(PROTO_FILES)
	@echo "Done."

## Generate Python gRPC code for all services
proto-py: $(PROTO_FILES)
	@echo "Generating Python gRPC code..."
	@python -m grpc_tools.protoc \
		-I$(PROTO_BASE_DIR) \
		--python_out=$(PYTHON_OUT_DIR) \
		--pyi_out=$(PYTHON_OUT_DIR) \
		--grpc_python_out=$(PYTHON_OUT_DIR) \
		$(PROTO_FILES)
	@echo "Done."

## Clean generated files
clean:
	rm -rf $(OUT_BASE_DIR)

runworker:
ifndef wf
	$(error workflow is not set. Usage: make runworker wf=<use_case_name, e.g. populate_user>)
endif
	go run $(wf)/worker/main.go

runpyworker:
ifndef wf
	$(error workflow is not set. Usage: make runworker wf=<use_case_name, e.g. populate_user>)
endif
	uv run -m $(wf).worker.$(wf)_worker

runstarter:
ifndef wf
	$(error workflow is not set. Usage: make runstarter wf=<use_case_name, e.g. populate_user>)
endif
	go run starters/$(wf)/main.go

format:
	golangci-lint fmt

lint: format
	golangci-lint run

.PHONY: proto proto-go clean runworker runstarter format lint
