package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/grokify/gotilla/net/httputilmore"
	visautil "github.com/grokify/oauth2more/visa"
	"github.com/joho/godotenv"
)

func getHelloWorld(client *http.Client) {
	apiURL := "https://sandbox.api.visa.com/vdp/helloworld"

	resp, err := client.Get(apiURL)
	if err != nil {
		log.Fatal(err)
	}

	httputilmore.PrintResponse(resp, true)
}

func main() {
	if len(os.Getenv("ENV_PATH")) > 0 {
		err := godotenv.Load(os.Getenv("ENV_PATH"))
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err := godotenv.Load()
		if err != nil {
			log.Fatal(err)
		}
	}

	client, err := visautil.NewClientEnv()
	if err != nil {
		log.Fatal(err)
	}

	getHelloWorld(client)

	fmt.Println("DONE")
}
