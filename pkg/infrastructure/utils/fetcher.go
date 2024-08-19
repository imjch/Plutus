package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type FetcherFunc func(url string, ret any) error

func FetchFromAPI(url string, ret any) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	res := string(body)
	err = json.Unmarshal([]byte(res), &ret)
	if err != nil {
		return fmt.Errorf("errFechFromAPI:%w", err)
	}
	return nil
}
