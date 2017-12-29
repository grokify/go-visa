package main

import (
	"flag"
	"fmt"

	"github.com/grokify/gotilla/time/timeutil"
)

var (
	URL        = "https://sandbox.api.visa.com/travelnotificationservice/v1/travelnotification/itinerary"
	MerchatURL = "https://sandbox.api.visa.com/merchantsearch/v1/search"
	certFile   = flag.String("cert", "someCertFile", "A PEM encoded certificate file.")
	keyFile    = flag.String("key", "someKeyFile", "A PEM encoded private key file.")
)

type TravelNotificationRequest struct {
	AddTravelItinerary TravelItinerary `json:"addTravelItinerary,omitempty"`
}

type TravelItinerary struct {
	PartnerBid            string                  `json:"partnerBid,omitempty"`
	UserId                string                  `json:"userId,omitempty"`
	PrimaryAccountNumbers []AccountNumber         `json:"primaryAccountNumbers,omitempty"`
	Destinations          []Destination           `json:"destinations,omitempty"`
	DepartureDate         timeutil.RFC3339YMDTime `json:"departureDate,omitempty"`
	ReturnDate            timeutil.RFC3339YMDTime `json:"returnDate,omitempty"`
}

type AccountNumber struct {
	CardAccountNumber string `json:"cardAccountNumber,omitempty"`
}

type Destination struct {
	Country string `json:"country,omitempty"`
	State   string `json:"state,omitempty"`
}

func main() {
	fmt.Println("DONE")
}
