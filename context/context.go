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

	log.Printf("Loading Application Context. %v %p", ctx, ctx)

	log.Printf("Opening Database Connection.")

	db, err := sql.Open("mysql", "xxx:yyyy@tcp(localhost:3307)/mysql")

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
