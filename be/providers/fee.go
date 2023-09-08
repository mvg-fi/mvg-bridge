package providers

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/tidwall/gjson"
)

func AssetFee(assetID string) (string, string) {
	path := fmt.Sprintf("https://api.mixin.one/network/assets/%s", assetID)
	resp, err := http.Get(path)
	if err != nil {
		log.Println(fmt.Sprintf("http.Get(%s)", path), err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	if gjson.Get(string(body), "data.fee").Exists() {
		return gjson.Get(string(body), "data.chain_id").String(), gjson.Get(string(body), "data.fee").String()
	}
	return "", ""
}
