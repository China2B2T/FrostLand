package main

import (
	"encoding/json"
	"frostland/frostland"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

type Config struct {
	BindAddress  string `json:"bindAddress"`
	MongoAddress string `json:"mongoAddress"`
	MongoDb      string `json:"mongoDb"`
	MongoCol     string `json:"mongoCol"`
}

func loadConfig(fileName string, v interface{}) error {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, v)
	if err != nil {
		return err
	}

	return nil
}

func main() {

	r := gin.Default()

	v := Config{}
	err := loadConfig("./config.json", &v)
	if err != nil {
		panic(err)
	}

	frostland.IUpdateConfig(v.MongoAddress, v.MongoDb, v.MongoCol)
	println("Starting up FrostLand on " + v.BindAddress)

	// API
	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", frostland.Ping)
		v1.POST("/user/create", frostland.CreateUser)
		v1.GET("/user/query/nickname/:id", frostland.QueryUser)
		// v1.GET("/user/query/uuid/:uuid", frostland.QueryUUID)
	}

	r.Run(v.BindAddress)
}
