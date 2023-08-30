package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/go-chi/chi"
)

// Service is an interface that defines the required functions for handlers.
type Service interface {
	Get(key string) (string, error)
	Del(key string)  (string, error)
}

// ServiceImpl is a struct that implements the Service interface.
type ServiceImpl struct {
	redisClient *redis.Client
}

// Get retrieves the value for the given key from the Redis database.
func (s *ServiceImpl) Get(key string) (string, error) {
	value, err := s.redisClient.Get(key).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("Key not found")
	} else if err != nil {
		return "", err
	}
	return value, nil
}


func (s *ServiceImpl) Del(key string) (string, error) {

	_ , err := s.redisClient.Del(key).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("Key not found")
	} else if err != nil {
		return "", err
	} 

	return "ok", nil
}

// Server is a struct that represents the HTTP server.
type Server struct {
	service Service
}

// NewServer creates a new instance of the Server struct.
func NewServer(service Service) *Server {
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
		if err.Error() == "Key not found" {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 Page Not Found")
		} else {
			log.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s : %s", key, value)
}

func (s *Server) DelHandler(w http.ResponseWriter, r *http.Request) {
	var payload map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 
	}


	if len(payload) == 1 {
		for key, _ := range payload {
			_ , err := s.service.Del(key)

			if err == redis.Nil {
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprint(w, "key does not exist")
				return
			} else if err != nil {
				log.Println(err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}
	
	} else {
		// fmt.Fprint(w, "цеа")
	}

	// w.WriteHeader(http.StatusOK)
	// fmt.Fprint(w, "DELETE request processed")

}


func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "bib", // Replace with your Redis password
		DB:       0,     // Replace with your Redis database number
	})

	service := &ServiceImpl{
		redisClient: rdb,
	}

	server := NewServer(service)

	r := chi.NewRouter()
	r.Get("/get", server.GetHandler)
	r.Delete("/del", server.DelHandler)

	log.Fatal(http.ListenAndServe(":18080", r))
}