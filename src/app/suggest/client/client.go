package client

import (
	"fmt"
	. "gopkg.in/webdeskltd/dadata.v2"
	"time"
)

type Client struct {
	Key   string
	Count int
}

type SuggestResult struct {
	Storage   []string
	ExpiredAt time.Time
}

func (client *Client) Request(query string) SuggestResult {
	addresses, requestError := NewDaData(client.Key, "").SuggestAddresses(SuggestRequestParams{Query: query, Count: client.Count})
	if nil != requestError {
		fmt.Println(requestError)
	}

	var suggest SuggestResult

	for _, address := range addresses {
		suggest.Storage = append(suggest.Storage, address.UnrestrictedValue)
	}
	suggest.ExpiredAt = time.Now()

	return suggest
}
