package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	var parentRouter = setupRoutes()

	parentRouter.HandleFunc("/", ProcessRoot).Methods("GET")

	http.ListenAndServe(":10000", parentRouter)

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
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{Message: "%v"}`, "OK")))
}
