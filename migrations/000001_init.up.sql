create database if not exists monitoring;

create table if not exists monitoring.users
(
    id       UUID,
    login    String,
    password String,
    name     String
) engine = MergeTree
      primary key (login, password);

-- Добавить пользователя вручную !!!!
-- insert into monitoring.users(id, login, password, name) values (generateUUIDv4(), 'alisa', '67626b6c6d6264666c4a4c4b356b356c34333439664c4b4b6c6c643030343033393238363976626b69646779434740bd001563085fc35165329ea1ff5c5ecbdbbeef', 'Alisa Ganz')