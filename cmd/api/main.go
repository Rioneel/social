package main

import (
	"log"
	"os"
		"github.com/rioneel/social/internal/env"
		"github.com/rioneel/social/internal/store"
		"github.com/rioneel/social/internal/db"
	
)

func main(){
	cfg := config{
		addr : env.GetString("ADDR", ":3000"),
		db : dbConfig{
			addr : env.GetString("DB_ADDR", os.Getenv("POSTGRES_URL")),
		maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
		maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
		maxIdleTime: env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	db , err := db.New(
		cfg.db.addr ,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	 )
	 if err != nil{
		log.Panic(err)
	 }
	 defer db.Close()
	 log.Println("database connection pool established")
	store := store.NewStorage(db)
	
	app:= &application{
		config : cfg,
		store : store,
	}

	


	mux := app.mount()
	
	log.Fatal(app.run(mux))
}