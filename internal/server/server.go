package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	url2 "net/url"
)

const (
	email = "acmkantemir@gmail.com"

	Bitcoin  = "btc"
	USDT     = "bep20/usdt"
	Litecoin = "itc"
)

type RespBody struct {
	AddressIn   string `json:"address_in"`
	AddressOut  string `json:"address_out"`
	CallbackURL string `json:"callback_url"`
	Priority    string `json:"priority"`
	Status      string `json:"status"`
}

var Amount int

func MakeWalletAddress(currency string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("server: %s /\n", r.Method)

	})

	//3E1DuCVKdfddYg8PrQjMhrgZ3ohv9RcWMU      3Jq9z8vDr8fFjSPQBnHPy3FiQXTVdg4s3N

	server := http.Server{
		Addr:    ":8888/address",
		Handler: mux,
	}

	go server.ListenAndServe()

	url := fmt.Sprintf("https://api.cryptapi.io/%s/create/", Bitcoin)

	client := http.DefaultClient

	params := url2.Values{}
	params.Add("callback", "http://localhost:8888/address")
	params.Add("address", "1HELGKbKdyQCAEYwPVBqMd3xiZgLTqwahT")
	params.Add("pending", "0")
	params.Add("confirmations", "1")
	params.Add("email", "string")
	params.Add("post", "0")
	params.Add("priority", "default")
	params.Add("multi_token", "0")
	params.Add("multi_chain", "0")
	params.Add("convert", "0")

	req, err := http.NewRequest(http.MethodGet, url+"?"+params.Encode(), nil)
	if err != nil {
		log.Println(err, 1)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err, 2)
		return
	}
	defer resp.Body.Close()

	respbody := RespBody{
		AddressIn:   "",
		AddressOut:  "",
		CallbackURL: "",
		Priority:    "",
		Status:      "",
	}

	respJSON, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err, 3)
		return
	}
	fmt.Println(resp.Status)

	err = json.Unmarshal(respJSON, &respbody)
	if err != nil {
		log.Println(err, 4)
		return
	}

	fmt.Println("Response body:", respbody)
}
