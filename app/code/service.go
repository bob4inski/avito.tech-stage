
// package main

// import (
// 	"fmt"
// 	"github.com/go-redis/redis"
// )

// type Service interface {
// 	Get(key string) (string, error)
// 	Del(key string) (string, error)
// 	Set(key string, value string) (string, error)
// }

// // ServiceImpl is a struct that implements the Service interface.
// type ServiceImpl struct {
// 	redisClient *redis.Client
// }

// // Get retrieves the value for the given key from the Redis database.
// func (s *ServiceImpl) Get(key string) (string, error) {
// 	value, err := s.redisClient.Get(key).Result()
// 	if err == redis.Nil {
// 		return "", fmt.Errorf("Key not found")
// 	} else if err != nil {
// 		return "", err
// 	}
// 	return value, nil
// }

// func (s *ServiceImpl) Del(key string) (string, error) {

// 	_, err := s.redisClient.Del(key).Result()
// 	if err == redis.Nil {
// 		return "", fmt.Errorf("Key not found")
// 	} else if err != nil {
// 		return "", err
// 	}

// 	return "ok", nil
// }

// func (s *ServiceImpl) Set(key string, value string) (string, error) {

// 	set_result, err := s.redisClient.Set(key, value, 0).Result()

// 	if err != nil {
// 		return "", err
// 	}
// 	return set_result, nil
// }