package main

import (
	"stad_projekt/config"
	"stad_projekt/handler"
	"stad_projekt/routes"
	"stad_projekt/service"
	"stad_projekt/storage/postgres"
)
// @contact.name   API Jasurbek
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io
// @license.name  @Suyunov_Jasurbek
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	cfg := config.Load()
	strg := postgres.NewPostgres(cfg)
	s := service.NewService(strg)
	h := handler.Newhandler(s)
	routes.RouteSetup(h, cfg)
}