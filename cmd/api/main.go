package main

import (
	"flag"
	"go/mydelivery/cmd/api/handler"
	"go/mydelivery/domain/cliente"
	"go/mydelivery/persistence"
	"go/mydelivery/persistence/sqlite"
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

	db, err := persistence.NewConnection(cfg.dsn)
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
