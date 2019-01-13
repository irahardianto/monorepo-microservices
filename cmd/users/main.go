package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/irahardianto/monorepo-microservices/package/log"
	"github.com/irahardianto/monorepo-microservices/users/router"
	"github.com/irahardianto/monorepo-microservices/users/storage/mongodb"
	//authmiddleware "github.com/irahardianto/monorepo-microservices/users/middleware"
	"github.com/spf13/viper"

	mgo "gopkg.in/mgo.v2"
)

func init() {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading config file", err)
	}
}

func main() {
	//URI without ssl=true
	var mongoURI = viper.GetString("database.atlasConnectionString")
	dialInfo, err := mgo.ParseURL(mongoURI)
	if err != nil {
		panic(err)
	}

	tlsConfig := &tls.Config{}
	tlsConfig.InsecureSkipVerify = true

	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		return conn, err
	}
	session, err := mgo.DialWithInfo(dialInfo)

	if err != nil {
		log.Fatal("error while creating session", err)
	}

	s := &mongodb.Storage{session.DB(viper.GetString("database.mongoDbName"))}
	defer session.Close()

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(middleware.URLFormat)
	//r.Use(authmiddleware.ValidateToken)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	router := router.InitRouter(r, s)

	// if err := http.ListenAndServe(fmt.Sprintf(":%s", viper.GetString("server.port")), router); err != nil {
	// 	log.Fatal("error while serve http server", err)
	// }
	if err := http.ListenAndServe(fmt.Sprintf(":%s", "8081"), router); err != nil {
		log.Fatal("error while serve http server", err)
	}
}
