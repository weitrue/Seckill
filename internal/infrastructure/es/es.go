package es

import "github.com/olivere/elastic/v7"

func NewEsClient() {
	elastic.NewClient()
}
