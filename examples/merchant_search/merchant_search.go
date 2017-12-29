package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/grokify/gotilla/fmt/fmtutil"
	"github.com/grokify/gotilla/net/httputilmore"
	"github.com/grokify/gotilla/time/timeutil"
	"github.com/joho/godotenv"

	"github.com/grokify/go-visa"
	visautil "github.com/grokify/oauth2more/visa"
)

func DemoMerchantSearchRequest(merchantName string) visa.MerchantSearchRequest {
	t, err := time.Parse(timeutil.ISO8601NoTzMilli, "2017-11-04T17:40:50.903")
	if err != nil {
		log.Fatal(err)
	}
	tt := timeutil.ISO8601NoTzMilliTime{Time: t}

	return visa.MerchantSearchRequest{
		SearchAttrList: visa.MerchantSearchAttrList{
			MerchantName:        merchantName,
			MerchantCountryCode: "840",
		},
		SearchOptions: visa.MerchantSearchOptions{
			WildCard:        []string{"merchantName"},
			MaxRecords:      5,
			MatchIndicators: true,
			MatchScore:      true,
			Proximity:       []string{"merchantName"},
		},
		ResponseAttrList: []string{"GNSTANDARD"},
		Header: visa.MerchantSearchHeader{
			RequestMessageId: "Request_001",
			StartIndex:       0,
			MessageDateTime:  tt,
		},
	}
}

func merchantSearchDirect(client *http.Client, merchantName string) (*visa.MerchantSearchResponse, error) {
	apiURL := "https://sandbox.api.visa.com/merchantsearch/v1/search"

	reqBody := DemoMerchantSearchRequest(merchantName)

	fmtutil.PrintJSON(reqBody)

	cm := httputilmore.ClientMore{Client: client}
	resp, err := cm.PostToJSON(apiURL, reqBody)
	if err != nil {
		log.Fatal(err)
	}
	httputilmore.PrintResponse(resp, true)

	res := &visa.MerchantSearchResponse{}

	err = httputilmore.UnmarshalResponseJSON(resp, res)
	return res, err
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

	merchantName := "amazon"

	if 1 == 1 {
		res, err := merchantSearchDirect(client, merchantName)
		if err != nil {
			log.Fatal(err)
		}
		fmtutil.PrintJSON(res)
	}

	if 1 == 1 {
		api := visa.NewMerchantsearchApi()
		api.Configuration.BasePath = "https://sandbox.api.visa.com"
		api.Configuration.Transport = client.Transport

		reqBody := DemoMerchantSearchRequest(merchantName)

		item, res, err := api.SearchMerchant(reqBody)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(res.StatusCode)

		httputilmore.PrintResponse(res.Response, true)

		fmtutil.PrintJSON(item)
	}
	fmt.Println("DONE")
}
