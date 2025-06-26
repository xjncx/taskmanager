# Переменные
PORT ?= 8080
GOFILES := $(shell find . -name '*.go' -not -path "./vendor/*")

# Цель: собрать бинарник
build:
	go build -o bin/taskmanager ./cmd/taskmanager

# Цель: запустить
run:
	go run ./cmd/taskmanager

# Цель: запуск с .env
dev:
	PORT=$(PORT) go run ./cmd/taskmanager

# Цель: тестирование
test:
	go test ./...

# Цель: удалить бинарники
clean:
	rm -rf bin/

# Цель по умолчанию
.PHONY: build run dev test mock clean
