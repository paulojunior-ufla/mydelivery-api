package main

import (
	"database/sql"
	"flag"
	"go/mydelivery/cmd/api/handler"
	"go/mydelivery/model"
	"go/mydelivery/service/cliente"
	"go/mydelivery/service/entrega"
	"go/mydelivery/service/ocorrencia"
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

	clienteRepo := model.NewClienteRepositoryDB(db)
	handler.NewClienteHandler(clienteRepo, cliente.NewCatalogoService(clienteRepo)).InitRoutes(router)

	entregaRepo := model.NewEntregaRepositoryDB(db)
	handler.NewEntregaHandler(entregaRepo, entrega.NewSolicitaEntregaService(clienteRepo, entregaRepo)).InitRoutes(router)

	ocorrenciaRepo := model.NewOcorrenciaRepositoryDB(db)
	handler.NewOcorrenciaHandler(ocorrencia.NewRegistraOcorrenciaService(ocorrenciaRepo, entregaRepo)).InitRoutes(router)

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
