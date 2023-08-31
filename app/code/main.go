package main

import (
	// "context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/go-redis/redis"
)

var data = make(map[string]interface{})

func main() {
	http.HandleFunc("/set", setHandler)
	http.HandleFunc("/get", getHandler)
	http.HandleFunc("/del", delHandler)
	
	log.Fatal(http.ListenAndServe(":8080", nil))
	}



func setHandler(w http.ResponseWriter, r *http.Request) {
	var payload map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	if len(payload) == 1 {
		for key, value := range payload {
			rdb := redis.NewClient(&redis.Options{
				Addr:     "redis:6379",
				Password: "verystrongpassword", // Replace with your Redis password
				DB:       1,            // Replace with your Redis database number
			})
		
			_ , err := rdb.Set( key , value, 0).Result()
			if err != nil {
				fmt.Fprintf(w,"%s : error: '%s'", "some  errror while set", err)
				return 
			}
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Неправильный запрос\n")
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Key-value pair(s) set successfully")

}


func getHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	} 

	r.ParseForm()
	key := r.Form.Get("key")
	if key == "" {
		http.Error(w, "Отсутствует аргумент 'key'", http.StatusBadRequest)
		return
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "verystrongpassword", // Replace with your Redis password
		DB:       1,            // Replace with your Redis database number
	})

	value, err := rdb.Get(key).Result()
	if err == redis.Nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "404 Page Not Found")
	} else if err != nil {
		fmt.Println(err)
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s : %s", key, value)
	}
}



func delHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodDelete {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	} 

	var payload map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	if len(payload) == 1 {
		for key, _ := range payload {
			rdb := redis.NewClient(&redis.Options{
				Addr:     "redis:6379",
				Password: "verystrongpassword", // Replace with your Redis password
				DB:       1,            // Replace with your Redis database number
			})

			_ , err := rdb.Del(key).Result()
			if err == redis.Nil {
				fmt.Println("key does not exist")
			} else if err != nil {
				panic(err)
			} 
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Слишком много ключей\n")
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "DELETE request processed")
}
