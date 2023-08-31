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
		http.Error(w, "Отсутствует аргумент 'key'", http.StatusBadRequest)
		return
	}

	value, err := s.service.Get(key)
	if err != nil {
		if err == redis.Nil {
			// w.WriteHeader(http.StatusNotFound)
			http.Error(w, "В таблице отсутсвует запись об этом ключе",http.StatusNotFound)
			return
		} else {

			// w.WriteHeader(http.StatusInternalServerError)
			http.Error(w, "Произошла ошибка при получении данных",http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Key: %s\nValue: %s\n", key, value)

}

func (s *Server) DelHandler(w http.ResponseWriter, r *http.Request) {
	var delPayload map[string]string

	err := json.NewDecoder(r.Body).Decode(&delPayload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(delPayload) == 1 {
		for key := range delPayload {
			value, err := s.service.Del(key)
			if err != nil {
				log.Println(err)
				http.Error(w, "Internal Server Error\n", http.StatusInternalServerError)
			}
			if value == 0 {
				w.WriteHeader(http.StatusOK)
				fmt.Fprintf(w, "Удаление прошло успешно\n")
			} else if value == 1 {
				w.WriteHeader(http.StatusOK)
				fmt.Fprintf(w, "Такого ключа нет\n")
			}
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Неправильный запрос\n")
	}

	

}

func (s *Server) SetHandler(w http.ResponseWriter, r *http.Request) {

	var payload map[string]string

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(payload) == 1 {
		for key, value := range payload {

			setResult, err := s.service.Set(key, value)

			if err != nil {
				log.Println(err)
				http.Error(w, "Internal Server Error while set", http.StatusInternalServerError)
				return
			} else {
				w.WriteHeader(http.StatusOK)
				fmt.Fprintf(w,"%s : %s", "Запись в базу данных прошла успешно", setResult)
			}
		}

	} else {
		
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Неправильный запрос\n")
	}
}