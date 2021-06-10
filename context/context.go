package context

import (
	"database/sql"
	"fmt"
	"log"
)

type ApplicationContext struct {
	database *sql.DB
}

// Load Application
func (ctx *ApplicationContext) Load() error {

	if ctx.database != nil {
		return nil
	}

<<<<<<< HEAD
	log.Printf("Loading Application Context. %v %p", ctx, ctx)

	log.Printf("Opening Database Connection.")
	db, err := sql.Open("mysql", "b2b:P9173usjs@tcp(localhost:3307)/mysql")
	//db, err := sql.Open("mysql", "root:12345678@tcp(localhost:3307)/mysql")
=======
	log.Printf("Opening Database Connection.")
	db, err := sql.Open("mysql", "xxxx:dddd@tcp(localhost:3307)/demo")
>>>>>>> fa17b32824163e69b394241027f381806621e5ea
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	ctx.database = db
	log.Printf("Application Context Loaded")

	return nil
}

// Database returns databse connection
func (ctx *ApplicationContext) Database() *sql.DB {
	return ctx.database
}
