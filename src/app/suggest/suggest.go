package suggest

import (
	. "example.com/suggest/src/app/suggest/cache"
	. "example.com/suggest/src/app/suggest/client"
	. "example.com/suggest/src/app/suggest/environment"
	. "example.com/suggest/src/app/suggest/server"
)

type Suggest struct {
	client Client
	server Server
	cache  Cache
}

func (suggest *Suggest) Create(variables Variables) *Suggest {
	suggest.cache = Cache{
		Size:            variables.CacheSize,
		ClearPercentage: variables.CacheClearPercentage,
	}

	suggest.client = Client{
		Key:   variables.DadataKey,
		Count: variables.DadataSuggestCount,
	}

	suggest.server = Server{
		Host: variables.ServerHost,
		Port: variables.ServerPort,

		Client: &suggest.client,
		Cache:  &suggest.cache,
	}

	return suggest
}

func (suggest *Suggest) Run() *Suggest {
	suggest.cache.Create()
	suggest.server.Run()

	return suggest
}
