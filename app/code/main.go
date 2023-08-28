package main

import (
	// "context"
	"fmt"
	"github.com/go-redis/redis"
)


func main() {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "bib", // Replace with your Redis password
		DB:       1,            // Replace with your Redis database number
	})

	
	key := "boba"
	value := "bobas"


	set_result, err := rdb.Set( key, value, 0).Result()
	if err != nil {
		fmt.Println("some  errror", err)
		return 
	}
	fmt.Println(set_result)


	val2, err := rdb.Get(key).Result()
	if err == redis.Nil {
		fmt.Println("key does not exist")
	} else if err != nil {
		panic(err)
		fmt.Println(err)
	} else {
		fmt.Println(key," : ", val2)
	}



	del_result, err := rdb.Del("dsfdfg").Result()
	if err == redis.Nil {
		fmt.Println("key does not exist")
	} else if err != nil {
		panic(err)
		fmt.Println(err)
	} else {
		fmt.Println(del_result)
	}



	// if val != "" {
	// 	fmt.Println("There is no key", a )
	// }
	// fmt.Println("key", val)

}
