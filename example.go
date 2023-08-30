package test

import (
	"fmt"
	"log"
	"github.com/go-redis/redis"
)

// Интерфейс Сервис
type Service interface {
	Set(key string, value interface{}) error
	Get(key string) (interface{}, error)
	Del(key string) error
}

// Структура СервисImpl
type ServiceImpl struct {
	db *redis.Client
	}

// Реализация метода GetData() интерфейса Service
func (s *ServiceImpl) GetData() string {

	return "Данные из базы данных"
}

// Структура Сервер
type Server struct {
	service Service
}

// Обработчик запроса
func (s *Server) HandleRequest() {
	data := s.service.GetData()
	fmt.Println("Полученные данные:", data)
	// Здесь можно продолжить обработку полученных данных
}

func test() {
// Создание подключения к базе данных
	db, err := sql.Open("driverName", "dataSourceName")
	if err != nil {
		log.Fatal(err)
	}

	// Инициализация сервиса с подключением к базе данных
	service := &ServiceImpl{
		db: db,
	}

	// Создание сервера с сервисом
	server := &Server{
		service: service,
	}

	// Обработка запроса
	server.HandleRequest()
}