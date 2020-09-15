package main

import (
	"leboncoin/handlers/FizzBuzzHandler"
	"leboncoin/handlers/StatisticsHandler"
	"leboncoin/model"

	"github.com/go-akka/configuration"
	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

func main() {
	config := configuration.LoadConfig("configuration.hcon")

	router := routing.New()
	Routes(router, config)
}

func Routes(router *routing.Router, config *configuration.Config) {

	defaultLimit := config.GetInt64("defaults.limit")
	fizzBarHandler := FizzBuzzHandler.New(defaultLimit)

	stats := model.Statistics{make(map[string]int)}
	statsHandler := StatisticsHandler.New(stats)

	router.Get("/fizzbuzz", statsHandler.SaveStats, fizzBarHandler.GetFizzBuzz)
	router.Get("/mostUsed", statsHandler.SaveStats, statsHandler.GetMostUsed)

	panic(fasthttp.ListenAndServe(":8080", router.HandleRequest).Error())
}
