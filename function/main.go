package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	el "github.com/olivere/elastic"

	"github.com/valentyn88/presentation/function/handlers"
	"github.com/valentyn88/presentation/function/storage"
)

func main() {
	log.SetOutput(os.Stdout)

	client, err := el.NewClient(el.SetURL("http://elasticsearch:9200"))
	if err != nil {
		log.Fatalf("init elastic client error: %s", err.Error())
	}

	elasticSt := storage.ElasticStorage{ElasticClient: client}
	if err := elasticSt.InitFixtures(); err != nil {
		log.Fatalf("couldn't init fixtures error: %s", err.Error())
	}

	var h handlers.Handler
	h.Storage = elasticSt

	r := chi.NewRouter()
	r.Use(h.AuthMiddleware)
	r.Route("/v1", func(r chi.Router) {
		r.With(h.QueryParamsMiddleware).Get("/products", h.Products)
	})

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("couldn't start search server error: %s", err.Error())
	}
}
