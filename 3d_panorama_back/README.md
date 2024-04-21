# Инструкция по сборке  


 - Установить Postgresql
 - Развернуть бд файл миграции находиться в ```migration/migrate.sql```
 - Установить модули go: ```go mod tidy```
 - Настроить конфиг в ```internals/cfg/config.go```
 - Запуск: ```go run cmd/main.go```


 