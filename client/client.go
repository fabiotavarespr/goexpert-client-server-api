package main

import (
	"encoding/json"
	"fmt"
	"github.com/fabiotavarespr/goexpert-client-server-api/quote"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	urlQuote      = "http://localhost:8080/cotacao"
	timeoutClient = time.Millisecond * 300
	quoteFile     = "cotacao.txt"
	output        = "Dólar: "
)

func main() {
	c := http.Client{Timeout: timeoutClient}
	resp, err := c.Get(urlQuote)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	var jsonResponse quote.Quote
	json.Unmarshal(body, &jsonResponse)

	if string(jsonResponse.Bid) != "" {
		f, err := os.Create(quoteFile)
		if err != nil {
			log.Println(err)
			panic(err)
		}
		defer f.Close()

		_, err = f.Write([]byte(output + string(jsonResponse.Bid)))
		if err != nil {
			log.Println(err)
			panic(err)
		}
		fmt.Println(output + string(jsonResponse.Bid))
	}
}
