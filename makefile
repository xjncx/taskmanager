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

# Цель: запуск mockgen
mock:
	mockgen -source=internal/service/contract.go -destination=internal/service/mocks/contract_mock.go -package=mocks

# Цель: удалить бинарники
clean:
	rm -rf bin/

# Цель по умолчанию
.PHONY: build run dev test mock clean
