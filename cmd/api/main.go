package main

import (
	"database/sql"
	"flag"
	"go/mydelivery/cmd/api/handler"
	"go/mydelivery/data/sqlite"
	"go/mydelivery/domain/cliente"
	"log/slog"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
)

type config struct {
	dsn string
}

func main() {
	var cfg config
	flag.StringVar(&cfg.dsn, "dsn", "file:database_dev.db", "Datasource url")

	db, err := openDB(cfg.dsn)
	if err != nil {
		slog.Error("db connection failed")
		os.Exit(1)
	}
	defer db.Close()

	router := httprouter.New()

	clienteRepo := sqlite.NewClienteRepository(db)
	handler.NewClienteHandler(clienteRepo, cliente.NewSalvaClienteService(clienteRepo)).InitRoutes(router)

	addr := ":8080"

	slog.Info("server started", "addr", addr)
	http.ListenAndServe(addr, router)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
