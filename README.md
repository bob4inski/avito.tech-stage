# avito.tech-stage
Задание 

## how to start
```docker compose up -d```
### Что где лежит
- [nginx.conf](./nginx/nginx.conf) конфигурация nginx 
- [redis.conf](./redis/redis.conf) конфигурация redis (пароль для подключения также указывается [тут](https://github.com/bob4inski/avito.tech-stage/blob/e54357d23a1acf3c9a60c05a162cd57aeeb837bc/app/code/main.go#L19-L23) )
- [main.go](./app/code/main.go)

## Запросы которые можно делать
- set value  //not working
    ```bash
    curl -X POST   -H "Content-Type: application/json"   -d '{ "name": "Robert"}'   http://localhost:8089/set
    ```
- get value
    ```bash
    curl -X GET  http://localhost:8089/get?key=name
    ```
- delete value
    ```bash
    curl -X DELETE -H "Content-Type: application/json" -d '{"name": "Robert"}' http://localhost:8089/del
    ```