# Subscriptions Service

REST-сервис для управления онлайн-подписками пользователей.

## Функциональность

- CRUDL-операции над подписками (создание, чтение, обновление, удаление, список)
- Подсчёт общей стоимости подписок по фильтрам (`user_id`, `дата`, `название`)
- Swagger-документация
- Миграции БД (PostgreSQL)
- Конфигурация через `.env`
- Docker Compose запуск

## Технологии

- Go + GORM
- PostgreSQL
- Swagger + swaggo
- Docker / docker-compose
- GitHub + .env конфиги

## Запуск проекта

1. Клонировать репозиторий:

```bash
git clone https://github.com/sorrrrow/subscriptions-service.git
cd subscriptions-service
```

2. Создать файл .env:

Создать файл .env:

```ini
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=subscriptions_db
SERVER_PORT=8080
```

3. Запуск с помощью Docker:

```bash
docker-compose up --build
```

4. Swagger доступен по адресу:

```bash
http://localhost:8080/swagger/index.html
```

## Миграции (если нужно вручную)
```bash
./migrate.exe -path migrations -database "postgres://postgres:your_password@localhost:5432/subscriptions_db?sslmode=disable" up
```

## Пример тела запроса
```json
{
  "service_name": "Yandex Plus",
  "price": 400,
  "user_id": "60601fee-2bf1-4721-ae6f-7636e79a0cba",
  "start_date": "07-2025"
}
```

Выполнено в рамках тестового задания на позицию Junior Golang Developer.
