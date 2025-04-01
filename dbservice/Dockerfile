# Используем базовый образ devcontainer с Go
ARG VARIANT="1.24"
FROM mcr.microsoft.com/vscode/devcontainers/go:${VARIANT}

###############################################################################
# 1) Системные пакеты и инструменты (от root)
###############################################################################

USER root

RUN apt-get update

# Пример установки act (GitHub Actions local runner)
RUN curl -s https://raw.githubusercontent.com/nektos/act/master/install.sh | bash

# Установка утилиты buf:
RUN BIN="/usr/local/bin" && \
    VERSION="1.26.1" && \
    curl -sSL \
    "https://github.com/bufbuild/buf/releases/download/v${VERSION}/buf-$(uname -s)-$(uname -m)" \
    -o "${BIN}/buf" && \
    chmod +x "${BIN}/buf"

# Установка golang-migrate
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-arm64.tar.gz \
    | tar xvz && \
    mv migrate /usr/local/bin/migrate

###############################################################################
# 2) Установка Go-инструментов (от vscode), чтобы кэш Go принадлежал vscode
###############################################################################

USER vscode

# Устанавливаем нужные Go-бинари
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.55.2 && \
    go install github.com/google/wire/cmd/wire@latest

###############################################################################
# 3) Если в базе были установлены модули от root — чиним права
###############################################################################

USER root
RUN chown -R vscode:vscode "$(go env GOMODCACHE)"

# Возвращаемся к пользователю vscode
USER vscode
