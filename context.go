package context

import (
	"database/sql"
	"fmt"
	"log"

	redis "github.com/go-redis/redis/v8"
)

type ApplicationContext struct {
	database *sql.DB
	rdb      *redis.Client
}

// Load Application
func (ctx *ApplicationContext) Load() error {

	log.Printf("Loading Application Context. %v %p", ctx, ctx)

	dberr := ctx.LoadDatabase()
	if dberr != nil {
		return dberr
	}

	// rediserr := ctx.setupRedis()
	// if rediserr != nil {
	// 	return rediserr
	// }

	log.Printf("Application Context Loaded")

	return nil

}

// LoadDatabase
func (ctx *ApplicationContext) LoadDatabase() error {

	if ctx.database != nil {
		return nil
	}

	log.Printf("Opening Database Connection.")
	db, err := sql.Open("mysql", "xxxx:yyyy@tcp(localhost:3307)/mysql")

	if err != nil {
		return fmt.Errorf("%v", err)
	}

	ctx.database = db
	log.Printf("Application Context Loaded")

	return nil

}

// LoadRedis
func (ctx *ApplicationContext) LoadRedis() error {

	if ctx.rdb != nil {
		return nil
	}

	log.Printf("Opening Redis Connection.")
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	ctx.rdb = rdb

	return nil

}

// Database returns databse connection
func (ctx *ApplicationContext) Database() *sql.DB {
	return ctx.database
}

// Redis returns redis connection
func (ctx *ApplicationContext) Redis() *redis.Client {
	return ctx.rdb
}
