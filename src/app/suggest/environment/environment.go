package environment

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Variables struct {
	ServerHost           string
	ServerPort           string
	DadataKey            string
	DadataSuggestCount   int
	CacheSize            int
	CacheClearPercentage int
}

func LoadVariables() (error, Variables) {
	loadError := godotenv.Load()

	if loadError == nil {
		suggestCount, _ := strconv.Atoi(os.Getenv("DADATA_SUGGEST_COUNT"))
		cacheSize, _ := strconv.Atoi(os.Getenv("CACHE_SIZE"))
		cacheClearPercentage, _ := strconv.Atoi(os.Getenv("CACHE_CLEAR_PERCENTAGE"))

		return nil, Variables{
			ServerHost:           os.Getenv("SERVER_HOST"),
			ServerPort:           os.Getenv("SERVER_PORT"),
			DadataKey:            os.Getenv("DADATA_KEY"),
			DadataSuggestCount:   suggestCount,
			CacheSize:            cacheSize,
			CacheClearPercentage: cacheClearPercentage,
		}
	} else {
		log.Fatal("Error while loading environment .env file not found")

		return nil, Variables{}
	}
}
