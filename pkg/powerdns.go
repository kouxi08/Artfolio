package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/kouxi08/Artfolio/utils"
)

type API struct {
	Endpoint string
	ApiKey   string
}

func GetAPI() API {
	return API{
		Endpoint: os.Getenv("ENDPOINT"),
		ApiKey:   os.Getenv("APIKEY"),
	}

}

func GetZoneList() {
	utils.Env()
	api := GetAPI()

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
func AddRecords(name, recordType, ttl, content string) (*http.Response, error) {
	utils.Env()

	method := http.MethodPatch
	//追加するレコードを指定
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
	//httpリクエストする処理
	resp, err := SendHTTPRequest(params, method)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// レコード削除処理
func DeleteRecords(name string) (*http.Response, error) {
	utils.Env()

	method := http.MethodPatch
	//削除するレコードの指定
	params := map[string]interface{}{
		"rrsets": []map[string]interface{}{
			{
				"type":       "CNAME",
				"name":       name,
				"changetype": "DELETE",
			},
		},
	}
	//httpリクエストする処理
	resp, err := SendHTTPRequest(params, method)
	if err != nil {
		return nil, err
	}
	return resp, nil

}

func SendHTTPRequest(params map[string]interface{}, method string) (*http.Response, error) {
	api := GetAPI()
	//json形式にするやつ
	payloadBytes, _ := json.Marshal(params)
	//zoneを指定してエンドポイントにつける
	zone := os.Getenv("ZONE")
	endpoint := api.Endpoint + zone
	//リクエスト処理
	req, _ := http.NewRequest(method, endpoint, bytes.NewReader(payloadBytes))
	req.Header.Set("X-API-Key", api.ApiKey)
	req.Header.Set("Content-Type", "application/json")

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
