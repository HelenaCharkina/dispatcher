
### Запуск postgres в docker

```bash
./docker run --name=dispatcher-db -e POSTGRES_PASSWORD='123' -p 5436:5432 -d postgres
```

### Создание миграции

```bash
./migrate create -ext sql -dir ./migrations -seq init
```

### Запуск миграций

```bash
./migrate -path ./migrations -database 'postgres://postgres:123@localhost:5436/monitoring?sslmode=disable' up
```

### Откат миграций

```bash
./migrate -path ./migrations -database 'postgres://postgres:123@localhost:5436/monitoring?sslmode=disable' down
```