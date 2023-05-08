package main

import (
	"stad_projekt/config"
	"stad_projekt/handler"
	"stad_projekt/routes"
	"stad_projekt/service"
	"stad_projekt/storage/postgres"
)

func main() {
	cfg := config.Load()
	strg := postgres.NewPostgres(cfg)
	s := service.NewService(strg)
	h := handler.Newhandler(s)
	routes.RouteSetup(h,cfg)
}