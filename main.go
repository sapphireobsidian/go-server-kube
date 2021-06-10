package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go/server/context"
	"github.com/gorilla/mux"
)

var appContext = context.ApplicationContext{}

func main() {

	var parentRouter = setupRoutes()

	parentRouter.HandleFunc("/", ProcessRoot).Methods("GET")

<<<<<<< HEAD
	http.ListenAndServe(":8080", parentRouter)
=======
	http.ListenAndServe(":10000", parentRouter)
>>>>>>> fa17b32824163e69b394241027f381806621e5ea

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

}

func ProcessRoot(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{Root: "%v"}`, "OK")))
}

func ProcessTestSql(w http.ResponseWriter, r *http.Request) {

	{
		err := appContext.Load()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf(`{Message: "%v"}`, err.Error())))
			return
		}
	}

<<<<<<< HEAD
	results, err := appContext.Database().Query("SELECT User from user;")
=======
	results, err := appContext.Database().Query("SELECT CURRENT_TIMESTAMP")
>>>>>>> fa17b32824163e69b394241027f381806621e5ea
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{Message: "%v"}`, err.Error())))
		return
	}

<<<<<<< HEAD
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
	w.Write([]byte(fmt.Sprintf(`{Message: "%v", Next=%v}`, "OK", unames)))
=======
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{Message: "%v", Next=%v}`, "OK", results.Next())))
>>>>>>> fa17b32824163e69b394241027f381806621e5ea

}
