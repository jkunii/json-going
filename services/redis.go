package services

import (
	"fmt"

	"gopkg.in/redis.v5"
)

func genKey(method, key string) string {

	return method + ":" + key
}

func GetClient() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := client.Ping().Result()
	if err != nil {
		fmt.Println("Redis Connecting Error .... Try again later")
		return nil, err
	}
	return client, nil
}

func SetResponse(method, key, value string) error {
	c, err := GetClient()
	if err != nil {
		fmt.Sprintln("Set Error")
	} else {
		err = c.Set(genKey(method, key), value, 0).Err()
		if err != nil {
			return err
		}
	}
	return nil
}

func GetResponse(method, key string) string {
	c, err := GetClient()
	if err != nil {
		fmt.Sprintln("Get Error")
	} else {
		response, err := c.Get(genKey(method, key)).Result()
		if err != nil {
			panic(err)
		}
		return response
	}
	return ""
}
