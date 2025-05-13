package main

import (
	"go-fiber-app/configs"
	"go-fiber-app/pkg"
	s "go-fiber-app/server"
)

func main() {
	configs.NewAppInitTime()
	cfg := configs.NewConfig()
	db := pkg.MongoBuilder(cfg)
	// cache := pkg.NewRedis(cfg)
	log := pkg.NewAppLogsZap()
	// server := servers.NewServer(cfg, db, cache, log)
	server := s.NewServer(cfg, db , log)
	server.Start()
}
