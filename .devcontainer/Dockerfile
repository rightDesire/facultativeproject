ARG VARIANT="1.24"
FROM mcr.microsoft.com/vscode/devcontainers/go:${VARIANT}

# Обновляем пакеты и устанавливаем дополнительные утилиты (если нужны)
RUN apt-get update && apt-get install -y \
    curl \
    git \
    && rm -rf /var/lib/apt/lists/*

# Устанавливаем рабочую директорию. Она должна соответствовать пути, указанному в devcontainer.json
WORKDIR /workspaces/${localWorkspaceFolderBasename}

# Не запускаем приложение – контейнер будет держаться в режиме ожидания,
# т.к. в docker-compose.yaml прописана команда: sleep infinity
CMD ["sleep", "infinity"]
