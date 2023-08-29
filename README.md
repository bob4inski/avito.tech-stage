# avito.tech-stage
Задание 

## how to
```docker compose up -d```

``` cd app/code ```

``` go run main.go ```


## Запросы которые можно делать
- set value 
    ```bash
    curl -X POST   -H "Content-Type: application/json"   -d '{ "name": "Robert"}'   http://localhost:8080/set
    ```
- get value
    ```bash
    curl -X GET  http://localhost:8080/get?key=name
    ```
- delete value
    ```bash
    curl -X DELETE -H "Content-Type: application/json" -d '{"name": "Robert"}' http://localhost:8080/del
    ```