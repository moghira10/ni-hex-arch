package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
	"github.com/ni/checkin"
	"github.com/ni/storage/cassandra"
)

func main() {

	var checkinRepo checkin.Repository

	checkinRepo = cassandra.NewCassandraCheckinRepository(cassandraConnection())

	checkinService := checkin.NewService(checkinRepo)
	checkinHandler := checkin.NewHandler(checkinService)

	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	logger.Println("Server is starting...")

	router := mux.NewRouter().StrictSlash(true)
	router.Handle("/addCheckIn", logging(logger, http.HandlerFunc(checkinHandler.AddCheckIn))).Methods("POST")
	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}

func cassandraConnection() *gocql.Session {

	fmt.Println("Connecting to Cassandra")

	cluster := gocql.NewCluster(os.Getenv("CQL_HOST"))
	cluster.ProtoVersion = 4
	pass := gocql.PasswordAuthenticator{Username: os.Getenv("CQL_USER"), Password: os.Getenv("CQL_PASS")}
	cluster.Authenticator = pass
	cluster.Keyspace = os.Getenv("CQL_KEYSPACE")
	cluster.Timeout = 60 * time.Second
	cluster.ConnectTimeout = 10 * time.Second
	cql, err := cluster.CreateSession()

	if err != nil {
		log.Fatalf("%s", err)
		return nil
	}
	return cql

}

func logging(logger *log.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)

		bodyReq := ioutil.NopCloser(bytes.NewBuffer(body))
		r.Body = bodyReq

		defer func() {
			logger.Println(r.Method, r.URL.Path, r.RemoteAddr)
		}()

		next.ServeHTTP(w, r)
	})
}
