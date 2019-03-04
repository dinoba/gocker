package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/olivere/elastic"
)

//ElasticsearchHandler ..
type ElasticsearchHandler struct {
	latestI string
	prefix  string
	*elastic.Client
}

//NewElasticsearchHandler creates new handler
func NewElasticsearchHandler(elasticSearchServer string, indexPrefix string) (*ElasticsearchHandler, error) {
	// Create a client
	esh, err := elastic.NewClient(
		elastic.SetURL(elasticSearchServer),
	)
	if err != nil {
		// Handle error
		return nil, err
	}

	return &ElasticsearchHandler{
			Client:  esh,
			prefix:  indexPrefix,
			latestI: indexPrefix + time.Now().Format("2006-01-02"),
		},
		err
}

//Store save log to ES
func (handler *ElasticsearchHandler) Store(log DockerLog) (errReturn error) {
	ctx := context.Background()
	currentIndex := handler.prefix + "-" + time.Now().Format("2006-01-02")

	if handler.latestI != currentIndex {
		handler.latestI = currentIndex
		_, err := handler.CreateIndex(handler.latestI).Do(ctx)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("New index created ", handler.latestI)
		}
	}

	_, err := handler.Index().
		Index(handler.latestI).
		Type("_doc").
		BodyJson(log).
		Refresh("true").
		Do(ctx)

	if err != nil {
		errReturn = err
	}
	return
}
