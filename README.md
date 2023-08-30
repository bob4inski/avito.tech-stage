# avito.tech-stage
Задание 

## how to
```docker compose up -d```




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