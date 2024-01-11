package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	models "github.com/CS559-CSD-IITBH/search-service/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/olivere/elastic.v6"
)

func initClient() (*elastic.Client, error) {
	client, err := elastic.NewClient(elastic.SetURL(os.Getenv("ELASTICSEARCH_URL")))
	if err != nil {
		return nil, err
	}

	return client, nil
}

func ElasticQuery(term string) ([]models.Store, error) {
	client, err := initClient()
	if err != nil {
		log.Printf("Error creating Elasticsearch client: %v", err)
		return nil, err
	}

	nestedQuery := elastic.NewNestedQuery("Items", elastic.NewMatchQuery("Items.Name", term))

	searchResult, err := client.Search().
		Index(os.Getenv("MONGODB_COLLECTION")).
		Query(nestedQuery).
		Do(context.Background())

	if err != nil {
		log.Printf("Error executing search query: %v", err)
		return nil, err
	}

	var results []models.Store
	for _, hit := range searchResult.Hits.Hits {
		var store models.Store
		err := json.Unmarshal(*hit.Source, &store)
		if err != nil {
			log.Printf("Error unmarshaling JSON: %v", err)
			continue
		}
		results = append(results, store)
	}

	return results, nil
}

func SearchHandler(c *gin.Context) {
	searchTerm := c.Param("keyword")

	results, err := ElasticQuery(searchTerm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error executing search query",
		})
		return
	}

	resultsJSON, err := json.Marshal(results)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error marshaling search results",
		})
		return
	}

	c.JSON(http.StatusOK, resultsJSON)
}
