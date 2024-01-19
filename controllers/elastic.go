package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/CS559-CSD-IITBH/search-service/models"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

var esClient *elasticsearch.Client

func init() {
	config := elasticsearch.Config{
		Addresses: []string{os.Getenv("ELASTICSEARCH_URL")},
	}

	es, err := elasticsearch.NewClient(config)
	if err != nil {
		log.Fatalf("Error creating Elasticsearch client: %s", err)
	}

	esClient = es
}

func ElasticQuery(query map[string]interface{}) ([]models.Store, error) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, err
	}

	res, err := esClient.Search(
		esClient.Search.WithIndex(os.Getenv("ELASTICSEARCH_INDEX")),
		esClient.Search.WithBody(&buf),
		esClient.Search.WithContext(context.Background()),
	)

	if err != nil {
		log.Printf("Error executing search query: %s", err)
		return nil, err
	}
	defer res.Body.Close()

	var results []models.Store
	if res.IsError() {
		log.Printf("Error response from Elasticsearch: %s", res.String())
		return nil, fmt.Errorf("error response from Elasticsearch")
	}

	var searchResult map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&searchResult); err != nil {
		log.Printf("Error parsing the search result: %s", err)
		return nil, err
	}

	hits, found := searchResult["hits"].(map[string]interface{})["hits"].([]interface{})
	if !found {
		log.Printf("Error getting hits from search result")
		return nil, fmt.Errorf("error getting hits from search result")
	}

	for _, hit := range hits {
		source, found := hit.(map[string]interface{})["_source"]
		if !found {
			log.Printf("Error getting source from search result hit")
			continue
		}

		store := models.Store{}
		if err := mapstructure.Decode(source, &store); err != nil {
			log.Printf("Error decoding source to Store model: %s", err)
			continue
		}

		results = append(results, store)
	}

	return results, nil
}

func GlobalSearch(c *gin.Context) {
	searchTerm := c.Param("keyword")

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"Items.Name": searchTerm,
			},
		},
	}

	results, err := ElasticQuery(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error executing search query",
		})
		return
	}

	c.JSON(http.StatusOK, results)
}

func StoreSearch(c *gin.Context) {
	storeID := c.Param("storeID")
	searchTerm := c.Query("keyword")

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []map[string]interface{}{
					{
						"match": map[string]interface{}{
							"Items.Name": searchTerm,
						},
					},
					{
						"match": map[string]interface{}{
							"StoreID": storeID,
						},
					},
				},
			},
		},
	}

	results, err := ElasticQuery(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error executing store-specific search query",
		})
		return
	}

	c.JSON(http.StatusOK, results)
}
