package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"main/service"
	"net/http"
)

type Server struct {
	service service.Service
}


// NewServer creates a new instance of the Server struct.
func NewServer(service service.Service) *Server {
	return &Server{
		service: service,
	}
}

// GetHandler handles the GET request for retrieving a value from the Redis database.
func (s *Server) GetHandler(w http.ResponseWriter, r *http.Request) {

	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "Отсутствует аргумент 'key'\n", http.StatusBadRequest)
		return
	}

	value, err := s.service.Get(key)
	if err != nil {
		if err == redis.Nil {
			log.Println("Ключа '",key ,"' нет")
			http.Error(w, "В таблице отсутсвует запись об этом ключе",http.StatusNotFound)
			return
		} else {
			http.Error(w, "Произошла ошибка при получении данных",http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Key: %s\nValue: %s\n", key, value)
	return 
}

func (s *Server) DelHandler(w http.ResponseWriter, r *http.Request) {
	
	var payload map[string]string

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if len(payload) == 1 {
		for key, _ := range payload {
			response, err := s.service.Del(key)

			if err != nil {
				log.Println(err)
				w.Header().Set("Content-Type", "text/plain")
				http.Error(w, "Internal Server Error\n", http.StatusInternalServerError)
			}
			if response == 1 {
				w.WriteHeader(http.StatusOK)
				w.Header().Set("Content-Type", "text/plain")
				fmt.Fprintf(w, "Удаление прошло успешно\n")
			} else if response == 0 {
				w.WriteHeader(http.StatusOK)
				w.Header().Set("Content-Type", "text/plain")
				fmt.Fprintf(w, "Такого ключа нет\n")
			}
		}
	} else {
		log.Println(payload)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Неправильный запрос\n")
	}
}

func (s *Server) SetHandler(w http.ResponseWriter, r *http.Request) {

	var payload map[string]string

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(payload) == 1 {
		for key, value := range payload {

			err := s.service.Set(key, value)

			if err != nil {
				log.Println(err)
				http.Error(w, "Произошла ошибка при записи", http.StatusInternalServerError)
				return
			} else {
				log.Println("Записано в бд")
				w.WriteHeader(http.StatusOK)
				fmt.Fprintf(w,"%s", "Запись в базу данных прошла успешно\n")
			}
		}

	} else  {
		log.Println("Неправильный запрос", payload)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Неправильный запрос\n")
	}
}