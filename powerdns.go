package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

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
