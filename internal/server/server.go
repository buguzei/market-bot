package server

import (
	"fmt"
	"log"
	"net/http"
)

const (
	bitcoin  = "btc"
	usdt     = "trc20/usdt"
	litecoin = "itc"
)

type ReqBody struct {
	Callback string `json:"callback"`
	Address  string `json:"address"`
}

var Amount int

func StartServer() {
	//url := fmt.Sprintf("https://api.cryptapi.io/%d/create/", bitcoin)
	url1 := "https://api.cryptapi.io/info/"

	client := http.DefaultClient

	req, err := http.NewRequest(http.MethodGet, url1, nil)
	if err != nil {
		log.Println(err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(resp.Body)
}

func MakeWalletAddress(userID int64) {
	mux := http.NewServeMux()
	mux.HandleFunc("/makeaddress", func(w http.ResponseWriter, r *http.Request) {

	})

	log.Println(http.ListenAndServe("localhost:8080", mux))
}
