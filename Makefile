APP_NAME=quotes

.PHONY: build run test clean fmt

## Сборка бинарника
build:
	go build -o bin/$(APP_NAME) ./cmd/server

## Запуск приложения
run:
	go run ./cmd/server

## Тесты всех пакетов
test:
	go test -v ./internal/...

## Форматирование кода
fmt:
	go fmt ./...

## Удаление бинарников
clean:
	rm -rf bin
