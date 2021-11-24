
### Запуск postgres в docker

```bash
./docker run --name=dispatcher-db -e POSTGRES_PASSWORD=123 -p 5436:5432 -d postgres
```
### Запуск clickhouse в docker

```bash
./docker run -d --name clickhouse --ulimit nofile=262144:262144 -p 9000:9000 -p 8123:8123  yandex/clickhouse-server
```

### Установка утилиты миграций

```bash
https://github.com/golang-migrate/migrate/tree/master/cmd/migrate#with-go-toolchain
```

### Создание миграции

```bash
./migrate create -ext sql -dir ./migrations -seq init
```

### Запуск миграций

```bash
./migrate -path ./migrations -database postgres://postgres:123@localhost:5436/postgres?sslmode=disable up
./migrate -database 'clickhouse://localhost:9000?x-multi-statement=true' -path ./migrations up
```

### Откат миграций

```bash
./migrate -path ./migrations -database postgres://postgres:123@localhost:5436/postgres?sslmode=disable down
./migrate -database 'clickhouse://localhost:9000?x-multi-statement=true' -path ./migrations down
```

### Билд docker-compose

```bash
docker-compose up --build dispatcher
```

### Запуск docker-compose

```bash
docker-compose up dispatcher
```