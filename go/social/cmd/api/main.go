package main

import (
	"log"

	"github.com/belone0/stack-trial/go/social/internal/env"
	"github.com/belone0/stack-trial/go/social/internal/store"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8081"),
	}
	store := store.NewStorage(nil)
	
	app := &application{
		config: cfg,
		store: store,
	}


	mux := app.mount()

	log.Fatal(app.run(mux))

}
