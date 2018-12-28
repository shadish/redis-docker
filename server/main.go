package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func GetClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

/*func ExampleNewClient() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
}*/

// func ExampleClient() {

// 	// Output: key value
// 	// key2 does not exist
// }

func main() {
	client := GetClient()
	err := client.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	// val, err := client.Get("key").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("key", val)

	val, err := client.Get("mykey").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("mykey", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}
