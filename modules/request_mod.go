package modules

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"
)

func XPoster(url string, PostData url.Values) string {
	body := strings.NewReader(PostData.Encode())
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		Logger("error", err.Error())
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		Logger("error", err.Error())
	}

	defer resp.Body.Close()

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		Logger("error", err.Error())
	}

	return string(byteArray)
}

func XGet(url string, header http.Header) string {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		Logger("error", err.Error())
	}

	req.Header = header

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		Logger("error", err.Error())
	}

	defer resp.Body.Close()

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		Logger("error", err.Error())
	}

	return string(byteArray)

}

func Filter(mapData interface{}, CutData string, target int64) bool {
	leng := reflect.ValueOf(mapData).Len()
	for i := 0; i < leng; i++ {
		if mapData.([]interface{})[i].(map[string]interface{})[CutData].(string) == fmt.Sprintf("%d", target) {
			return true
		}
	}

	return false
}

func Decoder(Data string) interface{} {
	var maps map[string]interface{}
	err := json.Unmarshal([]byte(Data), &maps)
	if err != nil {
		Logger("error", err.Error())
	}
	return maps

}

func Decoder_in(Data string) interface{} {
	var maps interface{}
	err := json.Unmarshal([]byte(Data), &maps)
	if err != nil {
		Logger("error", err.Error())
	}
	return maps

}
