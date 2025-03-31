# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.245.2/containers/go/.devcontainer/base.Dockerfile

ARG VARIANT="1.24"

FROM mcr.microsoft.com/vscode/devcontainers/go:${VARIANT}

# Update packages
RUN apt-get update

# [Package] Install golangci-lint 
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.55.2

# [Package] Install act
RUN curl -s https://raw.githubusercontent.com/nektos/act/master/install.sh | sudo bash

# [Pakage] Install openapi
RUN go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest

RUN chmod -R a+rwX /go/pkg

RUN BIN="/usr/local/bin" && \
    VERSION="1.26.1" && \
    curl -sSL \
    "https://github.com/bufbuild/buf/releases/download/v${VERSION}/buf-$(uname -s)-$(uname -m)" \
    -o "${BIN}/buf" && \
    chmod +x "${BIN}/buf"

# [Package] migrate 

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-arm64.tar.gz | tar xvz && \
    mv migrate /usr/local/bin/migrate

# [Package] wire

RUN go install github.com/google/wire/cmd/wire@latest

# [Package] Proto gen

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
