package main

import (
	"context"
	"encoding/json"
	"github.com/fabiotavarespr/goexpert-client-server-api/quote"
	"io"
	"log"
	"net/http"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	urlQuote        = "/cotacao"
	urlFindQuote    = "https://economia.awesomeapi.com.br/json/last/USD-BRL"
	httpPort        = ":8080"
	requestCurrency = "USDBRL"
	timeoutAPI      = time.Millisecond * 200
	timeoutDB       = time.Millisecond * 10
	databaseQuote   = "quote.db"
)

func main() {
	http.HandleFunc(urlQuote, handler)
	log.Printf("Starting server at port %s\n", httpPort)
	log.Fatal(http.ListenAndServe(httpPort, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	if r.URL.Path != urlQuote {
		log.Println("404 not found.")
		w.WriteHeader(http.StatusNotFound)
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	quote, err := FindQuote(ctx)
	if err != nil {
		log.Println(error.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "500 Internal server error.", http.StatusInternalServerError)
		return
	}

	if quote != nil && quote.Bid != "" {
		err = SaveQuote(ctx, quote)
		if err != nil {
			log.Println(error.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			http.Error(w, "500 Internal server error.", http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(quote)
}

func FindQuote(ctx context.Context) (*quote.Quote, error) {
	ctx, cancel := context.WithTimeout(ctx, timeoutAPI)
	defer cancel()

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, urlFindQuote, nil)
	if err != nil {
		log.Println(error.Error(err))
		return nil, err
	}
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Println(error.Error(err))
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(error.Error(err))
		return nil, err
	}

	var jsonResp map[string]interface{}
	json.Unmarshal(body, &jsonResp)

	var c quote.Quote
	jsonStr, _ := json.Marshal(jsonResp[requestCurrency])
	err = json.Unmarshal(jsonStr, &c)
	if err != nil {
		log.Println(error.Error(err))
		return nil, err
	}

	return &c, nil
}

func SaveQuote(ctx context.Context, quote *quote.Quote) error {
	ctx, cancel := context.WithTimeout(ctx, timeoutDB)
	defer cancel()

	db, err := initDatabase()
	if err != nil {
		log.Println(error.Error(err))
		return err
	}

	result := db.WithContext(ctx).Create(&quote)
	if result.Error != nil {
		log.Println(error.Error(result.Error))
		return result.Error
	}

	return nil
}

func initDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(databaseQuote), &gorm.Config{})
	if err != nil {
		log.Println(error.Error(err))
		return nil, err
	}

	err = db.AutoMigrate(&quote.Quote{})
	if err != nil {
		log.Println(error.Error(err))
		return nil, err
	}
	return db, nil
}
