package main

import (
	"log"
	"os"
		"github.com/rioneel/social/internal/env"
		"github.com/rioneel/social/internal/store"
)

func main(){
	cfg := config{
		addr : env.GetString("ADDR", ":3000"),
	}
	store := store.NewStorage(nil)
	
	app:= &application{
		config : cfg,
		store : store,
	}

	

	os.LookupEnv("")
	mux := app.mount()
	
	log.Fatal(app.run(mux))
}