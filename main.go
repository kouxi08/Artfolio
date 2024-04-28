package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type API struct {
	Endpoint string
	ApiKey   string
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}

	api := API{
		Endpoint: os.Getenv("ENDPOINT"),
		ApiKey:   os.Getenv("APIKEY"),
	}

	// showZones(api)
	resp, err := addRecords(api, "hoge.hoge.com.", "CNAME", "3600", "huga.hoge.com.")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("Response Status:", resp.Status)
}

func showZones(api API) {

	req, err := http.NewRequest("GET", api.Endpoint, nil)
	if err != nil {
		fmt.Print("1", err)
		return
	}

	req.Header.Set("X-API-Key", api.ApiKey)
	//リクエスト処理
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print("2", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Print(string(body))
}

// レコード追加処理
func addRecords(api API, name, recordType, ttl, content string) (*http.Response, error) {
	params := map[string]interface{}{
		"rrsets": []map[string]interface{}{
			{
				"name":       name,
				"type":       recordType,
				"ttl":        ttl,
				"changetype": "REPLACE",
				"comments":   []interface{}{},
				"records": []map[string]interface{}{
					{
						"content":  content,
						"disabled": false,
					},
				},
			},
		},
	}
	//json形式にするやつ
	payloadBytes, _ := json.Marshal(params)
	//zoneを指定してエンドポイントにつける
	zone := os.Getenv("ZONE")
	endpoint := api.Endpoint + zone
	//リクエスト処理
	req, err := http.NewRequest(http.MethodPatch, endpoint, bytes.NewReader(payloadBytes))
	req.Header.Set("X-API-Key", api.ApiKey)
	req.Header.Set("Content-Type", "application/json")

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
