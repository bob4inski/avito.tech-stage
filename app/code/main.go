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
				Addr:     "localhost:6379",
				Password: "bib", // Replace with your Redis password
				DB:       1,            // Replace with your Redis database number
			})
		
			set_result, err := rdb.Set( key , value, 0).Result()
			if err != nil {
				fmt.Fprintf(w, "some  errror", err)
				return 
			}
			fmt.Fprintf(w, set_result)
		}
	} else {
		// Обработка ошибки, если передано больше одной пары ключ-значение
	}

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
		Addr:     "localhost:6379",
		Password: "bib", // Replace with your Redis password
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
				Addr:     "localhost:6379",
				Password: "bib", // Replace with your Redis password
				DB:       1,            // Replace with your Redis database number
			})

			_ , err := rdb.Del(key).Result()
			if err == redis.Nil {
				fmt.Println("key does not exist")
			} else if err != nil {
				panic(err)
			} 
			// else {
			// 	fmt.Fprint(w, "Key-value pair(s) deleted successfully")
			// }
		}
	} else {
		// Обработка ошибки, если передано больше одной пары ключ-значение
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "DELETE request processed")
}





// func main() {

	// rdb := redis.NewClient(&redis.Options{
	// 	Addr:     "localhost:6379",
	// 	Password: "bib", // Replace with your Redis password
	// 	DB:       1,            // Replace with your Redis database number
	// })

	
// 	key := "boba"
// 	value := "bobas"


// 	set_result, err := rdb.Set( key, value, 0).Result()
// 	if err != nil {
// 		fmt.Println("some  errror", err)
// 		return 
// 	}
// 	fmt.Println(set_result)


// 	val2, err := rdb.Get(key).Result()
// 	if err == redis.Nil {
// 		fmt.Println("key does not exist")
// 	} else if err != nil {
// 		panic(err)
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(key," : ", val2)
// 	}



// 	del_result, err := rdb.Del("dsfdfg").Result()
// 	if err == redis.Nil {
// 		fmt.Println("key does not exist")
// 	} else if err != nil {
// 		panic(err)
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(del_result)
// 	}

// }
