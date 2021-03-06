package main

import (
	gocontext "context"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go/server/context"
	"github.com/gorilla/mux"
)

var appContext = context.ApplicationContext{}
var ctx = gocontext.Background()

func main() {

	var parentRouter = setupRoutes()

	parentRouter.HandleFunc("/", ProcessRoot).Methods("GET")

	http.ListenAndServe(":8080", parentRouter)

}

func setupRoutes() *mux.Router {

	parentRouter := mux.NewRouter().StrictSlash(true)

	apiRouter := parentRouter.PathPrefix("/api").Subrouter()
	apiRouter.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Unknown Path %s\n", r.RequestURI)
		w.WriteHeader(http.StatusNotFound)
	})
	apiRouter.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("%v %v \n", r.RequestURI, r.Method)
			next.ServeHTTP(w, r)
		})
	})

	setupRoutesApi(apiRouter)

	return parentRouter
}

func setupRoutesApi(apiRouter *mux.Router) {

	myRouter := apiRouter

	myRouter.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	myRouter.HandleFunc("/sql", ProcessTestSql).Methods("GET")
	myRouter.HandleFunc("/redis", ProcessTestRedis).Methods("GET")

}

func ProcessRoot(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{Root: "%v", Time: "%v"}`, "OK", now)))
}

func ProcessTestSql(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	{
		err := appContext.LoadDatabase()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf(`{Message: "%v"}`, err.Error())))
			return
		}
	}

	results, err := appContext.Database().Query("SELECT User from user;")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{Message: "%v"}`, err.Error())))
		return
	}

	var uname string
	var unames []string
	for results.Next() {
		err := results.Scan(&uname)
		if err != nil {
			log.Fatal(err)
		}
		unames = append(unames, uname)
		log.Println(uname)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{Message: "%v", Names=%v}`, now, unames)))

}

func ProcessTestRedis(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	{
		err := appContext.LoadRedis()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf(`{Message: "%v"}`, err.Error())))
			return
		}
	}

	{
		err := appContext.Redis().Set(ctx, "key", "value", 0).Err()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf(`{Message: "%v"}`, err.Error())))
			return
		}
	}

	val, err := appContext.Redis().Get(ctx, "key").Result()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{Message: "%v"}`, err.Error())))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{Message: "%v", Next=%v}`, now, val)))

}
