package main

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"log"
	"net/http"
)

var esClient *elastic.Client

func main() {
	var err error
	// Подключение к Elasticsearch
	esClient, err = elastic.NewClient(elastic.SetURL("http://elasticsearch:9200"), elastic.SetSniff(false))
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	http.HandleFunc("/search", searchHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	// Пример запроса в Elasticsearch
	query := elastic.NewMatchAllQuery()
	searchResult, err := esClient.Search().
		Index("test-index").
		Query(query).
		Do(context.Background())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(searchResult)
}
