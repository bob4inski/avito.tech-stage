// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"github.com/go-redis/redis"
// )

// type Server struct {
// 	service Service
// }

// func (s *Server) GetHandler(w http.ResponseWriter, r *http.Request) {

// 	key := r.URL.Query().Get("key")
// 	if key == "" {
// 		http.Error(w, "Отсутствует аргумент 'key'", http.StatusBadRequest)
// 		return
// 	}

// 	value, err := s.service.Get(key)
// 	if err != nil {
// 		if err.Error() == "Key not found" {
// 			w.WriteHeader(http.StatusNotFound)
// 			fmt.Fprint(w, "404 Page Not Found")
// 		} else {
// 			log.Println(err)
// 			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		}
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	fmt.Fprintf(w, "%s : %s", key, value)
// }

// func (s *Server) DelHandler(w http.ResponseWriter, r *http.Request) {
// 	var del_payload map[string]string

// 	err := json.NewDecoder(r.Body).Decode(&del_payload)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	if len(del_payload) == 1 {
// 		for key, _ := range del_payload {
// 			_, err := s.service.Del(key)
// 			if err == redis.Nil {
// 				w.WriteHeader(http.StatusNotFound)
// 				fmt.Fprint(w, "key does not exist")
// 				return
// 			} else if err != nil {
// 				log.Println(err)
// 				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 			}
// 		}

// 	} else {
// 		w.WriteHeader(http.StatusBadRequest)
// 		fmt.Fprint(w, "Wrong request")
// 		// fmt.Fprint(w, "цеа")
// 	}

// 	// w.WriteHeader(http.StatusOK)
// 	// fmt.Fprint(w, "DELETE request processed")

// }


// func (s *Server) SetHandler(w http.ResponseWriter, r *http.Request) {

// 	var payload map[string]string

// 	err := json.NewDecoder(r.Body).Decode(&payload)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	if len(payload) == 1 {
// 		for key, value := range payload {

// 			set_result, err := s.service.Set(key, value)

// 			if err != nil {
// 				log.Println(err)
// 				http.Error(w, "Internal Server Error while set key", http.StatusInternalServerError)
// 				return
// 			} else {
// 				w.WriteHeader(http.StatusOK)
// 				fmt.Fprintf(w, "%s : %s", "Запись в базу данных прошла успешно", set_result)
// 			}
// 		}

// 	} else {
// 		w.WriteHeader(http.StatusBadRequest)
// 		fmt.Fprint(w, "Wrong request")
// 	}
// }