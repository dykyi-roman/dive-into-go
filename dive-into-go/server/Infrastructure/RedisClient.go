package Infrastructure

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	Domain "github.com/dykyi-roman/dive-into-go/server/Domain/Model"
	"github.com/go-redis/redis"
)

type RedisClient struct{}

func (obj RedisClient) dbConnect() *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := client.Ping().Result()
	if err != nil {
		panic(err.Error())
	}

	return client
}

func (obj RedisClient) Get() []Domain.User {
	db := obj.dbConnect()

	val, err := db.Get("user:1").Result()
	if err != nil {
		panic(err.Error())
	}
	jsonData := []byte(val)

	var model Domain.User
	err2 := json.Unmarshal(jsonData, &model)
	if err2 != nil {
		panic(err.Error())
	}

	res := []Domain.User{}
	res = append(res, model)

	return res
}

func (obj RedisClient) Insert(r *http.Request) bool {
	db := obj.dbConnect()

	if r.Method == "POST" {
		name := r.FormValue("name")
		intAge, _ := strconv.Atoi(r.FormValue("age"))

		model := Domain.User{}
		model.Id = 1
		model.Name = name
		model.Age = int(intAge)

		var jsonData []byte
		jsonData, err := json.Marshal(model)
		if err != nil {
			log.Println(err)
		}

		error := db.Set("user:1", string(jsonData), 0).Err()
		if error != nil {
			panic(error)
		}
	}
	return true
}

func (obj RedisClient) Delete(r *http.Request) bool {

	//TODO:: To be continue...

	return true
}

func (obj RedisClient) Update(r *http.Request) bool {

	//TODO:: To be continue...

	return true
}
