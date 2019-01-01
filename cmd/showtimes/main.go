package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/irahardianto/monorepo-microservices/showtimes/router"
	"github.com/irahardianto/monorepo-microservices/showtimes/storage/mongodb"
	"github.com/spf13/viper"

	mgo "gopkg.in/mgo.v2"
)

func init() {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
}

func main() {
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{viper.GetString("database.mongoDbHost")},
		Username: viper.GetString("database.mongoDbUser"),
		Password: viper.GetString("database.mongoDbPassword"),
		Timeout:  60 * time.Second,
	})
	if err != nil {
		log.Fatalf("[createDbSession]: %s\n", err)
	}

	s := &mongodb.Storage{session.DB(viper.GetString("database.mongoDbName"))}
	defer session.Close()

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	router := router.InitRouter(r, s)

	log.Fatalf("%s", http.ListenAndServe(fmt.Sprintf(":%s", viper.GetString("server.port")), router))
}
