# golang-test-task
Тестовая задача для Golang разработчика

## Сборка и запуск

1. Установить [gb](https://getgb.io/) и [docker-compose](https://docs.docker.com/compose/)
1. Восстановить зависимости: `gb vendor restore`
1. Изменить кофигурацию (src/configstore/main.go)
1. Собрать проект: `gb build`
1. Опционально перегенерировать файлы (нужен gomock в $PATH): `gb generate`
1. Запустить тесты: `gb test -v`
1. Запустить базу данных: `docker-compose up -d`
1. Запустить миграции: `./bin/configstore migrate`
1. Запустить сервер: `./bin/configstore run`

## Сборка и запуск для "продакшона"

1. Недавно делал аналогичное, смысла дублировать код не вижу, в задаче этого не было: https://github.com/derlaft/golang-test-task

## Замечания

1. Микросервисы хороши тем, что можно писать их на разных (т.е. подходящих) языках с использованием разных (т.е. подходящих инструментов). По ряду причин, для подобного сервиса я бы не выбрал использовать связку go + postgres:
    1) Данные плохо нормализуются (у каждого конфига свои столбцы, операции по отдельным столбцам операций никаких не производится). MongoDB или даже redis тут были бы гораздо уместнее. Фактически, тут key-value хранилище.
    1) В go неудобно работать с динамическими структурами. Написание на python убрало бы кучу лишнего кода.

```
 % curl localhost:8078/get_config -d '{"type": "Develop.mr_robot", "data": "YHySKtEhYm"}'
{"Data":"YHySKtEhYm","Host":"WKyh","Port":46836,"Database":"vkNMTN","User":"zmPBh","Password":"teBuZLhZ","Schema":"YUEZgVvd"}

```


