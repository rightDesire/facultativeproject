.PHONY: run
run:
	go run ./cmd/web/.

.PHONY: wire
wire:
	wire ./internal/app/

# .PHONY: gen
# gen:  
# 	for tag in omain oprofile ofederation oproject otask oreminder ocatalog; do \
# 		rm -rf ./internal/web/$$tag/; \
# 		mkdir -p ./internal/web/$$tag; \
# 	done
 
# 	oapi-codegen -config openapi/.openapi  -include-tags about,health -package omain  openapi/openapi.yaml > ./internal/web/omain/api.gen.go
# 	oapi-codegen -config openapi/.openapi  -include-tags profile -package oprofile  openapi/openapi.yaml > ./internal/web/oprofile/api.gen.go
# 	oapi-codegen -config openapi/.openapi  -include-tags federation -package ofederation  openapi/openapi.yaml > ./internal/web/ofederation/api.gen.go
# 	oapi-codegen -config openapi/.openapi  -include-tags project -package oproject  openapi/openapi.yaml > ./internal/web/oproject/api.gen.go
# 	oapi-codegen -config openapi/.openapi  -include-tags task -package otask openapi/openapi.yaml > ./internal/web/otask/api.gen.go
# 	oapi-codegen -config openapi/.openapi  -include-tags reminder -package oreminder openapi/openapi.yaml > ./internal/web/oreminder/api.gen.go
# 	oapi-codegen -config openapi/.openapi  -include-tags catalog -package ocatalog openapi/openapi.yaml > ./internal/web/ocatalog/api.gen.go

.PHONY: migrate-new
migrate-new:
	migrate create -ext sql -dir ./migrations ${NAME}

.PHONY: migrate
migrate:
	migrate -path ./migrations -database "postgres://rightdesire:secret@postgres:5432/main?sslmode=disable" up 

.PHONY: migrate-down
migrate-down:
	migrate -path ./migrations -database "postgres://rightdesire:secret@postgres:5432/main?sslmode=disable" down  
	