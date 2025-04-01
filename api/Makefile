.PHONY: run
run:
	go run ./cmd/web/.

.PHONY: wire
wire:
	wire ./internal/app/

.PHONY: gen
gen:  
	for tag in ouser oroute; do \
		rm -rf ./internal/web/$$tag/; \
		mkdir -p ./internal/web/$$tag; \
	done
 
	oapi-codegen -config openapi/.openapi  -include-tags user -package ouser openapi/openapi.yaml > ./internal/web/ouser/api.gen.go
	oapi-codegen -config openapi/.openapi  -include-tags route -package oroute openapi/openapi.yaml > ./internal/web/oroute/api.gen.go

.PHONY: lint
lint:
	golangci-lint run --out-format=colored-line-number
