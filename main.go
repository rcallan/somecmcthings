package main

import(
	"fmt"
	// "encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	coinStr := "vechain"

	if len(os.Args) == 2 {
		coinStr = os.Args[1]
	}

	url := fmt.Sprintf("https://api.coinmarketcap.com/v1/ticker/%s/", coinStr)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		fmt.Errorf("bad status: %d", resp.StatusCode)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	veninfo := fmt.Sprintf("%s", b)

	fmt.Printf(veninfo)
}