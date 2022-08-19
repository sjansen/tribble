.PHONY: default
default: generate

.PHONY: generate
generate:
	go generate ./...
