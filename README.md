## ПОЛНАЯ ИНСТРУКЦИЯ ПО ЗАПУСКУ (для разработчика)
```bash
1. Установить утилиту миграций
go install -tags 'clickhouse' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

2. Запустить создание контейнера в докере
docker-compose up --build dispatcher

3. Запустить миграции (удалить комментарий в файле migrations/000001_init.up.sql)
migrate -database 'clickhouse://localhost:9000?x-multi-statement=true' -path ./migrations up

4. Добавить пользователя в базу вручную
insert into monitoring.users(id, login, password, name) values (generateUUIDv4(), 'alisa', '67626b6c6d6264666c4a4c4b356b356c34333439664c4b4b6c6c643030343033393238363976626b69646779434740bd001563085fc35165329ea1ff5c5ecbdbbeef', 'Alisa Ganz')

5. Установить настройки в файле conf.ini. Например, адрес клиента.

4. Запустить приложение 
go run dispatcher
```

## Памятка команд

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

