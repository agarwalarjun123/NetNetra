package model

import (
	"log"

	"github.com/olivere/elastic"
)

var client *elastic.Client

//establishes connection to ElasticSearch
func Connect() (*elastic.Client, error) {

	client, err := elastic.NewClient()
	if err != nil {
		return &elastic.Client{}, err

	}
	log.Println("ElasticSearch running on localhost:9200")
	return client, nil
}
