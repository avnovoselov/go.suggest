package server

import (
	"encoding/json"
	. "example.com/suggest/src/app/suggest/cache"
	. "example.com/suggest/src/app/suggest/client"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Server struct {
	Host string
	Port string

	Client *Client
	Cache  *Cache
}

func prepareQuery(query string) string {
	return strings.Trim(strings.ToLower(query), " \t")
}

func extractGetParameter(request *http.Request, parameter string) string {
	values, ok := request.URL.Query()[parameter]

	if !ok || len(values) < 0 {
		log.Panicf("GET parameter %[1]s not found\n", parameter)
		return ""
	}
	return prepareQuery(values[0])
}

func getSuggestResult(server *Server, request *http.Request) (string, SuggestResult) {
	query := extractGetParameter(request, "query")
	log.Printf("Request from %[1]s. Query \"%[2]s\"\n", request.RemoteAddr, query)

	cachedSuggestResult, containsInCache := server.Cache.Get(query)

	if containsInCache {
		return query, cachedSuggestResult
	} else {
		suggestResult := server.Client.Request(query)
		server.Cache.Store(query, suggestResult)

		return query, suggestResult
	}
}

func writeJSONResponse(writer http.ResponseWriter, suggests []string) {
	jsonString, jsonError := json.Marshal(suggests)

	if jsonError != nil {
		http.Error(writer, jsonError.Error(), http.StatusInternalServerError)
		log.Panicln("JSON error")
	} else {
		writer.Header().Set("Content-Type", "application/json")
		_, writeError := writer.Write(jsonString)
		if writeError != nil {
			log.Panicln("Write response error")
		}
	}
}

func (server *Server) suggest(writer http.ResponseWriter, request *http.Request) {
	_, suggestResult := getSuggestResult(server, request)
	writeJSONResponse(writer, suggestResult.Storage)
}

func (server *Server) Run() {
	log.Printf("Server run at %[1]s:%[2]s\n", server.Host, server.Port)

	http.HandleFunc("/suggest", server.suggest)

	err := http.ListenAndServe(fmt.Sprintf("%[1]s:%[2]s", server.Host, server.Port), nil)

	if err != nil {
		log.Fatal(err)
	}
}
