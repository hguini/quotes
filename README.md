# Цитатник

REST API-сервис на Go для хранения и управления цитатами.

## Функциональность

- Добавление новой цитаты (`POST /quotes`)
- Получение всех цитат (`GET /quotes`)
- Получение случайной цитаты (`GET /quotes/random`)
- Фильтрация по автору (`GET /quotes?author=Confucius`)
- Удаление цитаты по ID (`DELETE /quotes/{id}`)

---

## Установка и запуск

### 1. Клонировать репозиторий

```
git clone https://github.com/hguini/quotes.git
cd quotes
```

### 2 Запуск через Makefile

```
make run
```

---

## Примеры запросов (curl)

```
# Добавить цитату
curl -X POST http://localhost:8080/quotes \
  -H "Content-Type: application/json" \
  -d '{"author":"Confucius", "quote":"Life is simple, but we insist on making it complicated."}'

# Получить все цитаты
curl http://localhost:8080/quotes

# Получить случайную цитату
curl http://localhost:8080/quotes/random

# Фильтрация по автору
curl http://localhost:8080/quotes?author=Confucius

# Удалить цитату
curl -X DELETE http://localhost:8080/quotes/1
```

---
## Makefile: команды

```
make build - Сборка бинарника 
make run   - Запуск приложения
make test  - Запуск всех тестов
make fmt   - Форматирование Go-кода
make clean - Очистка собранных файлов
```
