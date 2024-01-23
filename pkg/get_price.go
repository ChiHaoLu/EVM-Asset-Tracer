package pkg

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

type CMCQuoteResponse struct {
	Status struct {
		ErrorCode    int    `json:"error_code"`
		ErrorMessage string `json:"error_message"`
	} `json:"status"`
	Data map[string][]struct {
		Name  string `json:"name"`
		Quote map[string]struct {
			Price float64 `json:"price"`
		} `json:"quote"`
	} `json:"data"`
}

func Quote(From string, To string) (float64, error) {
	req, err := http.NewRequest("GET", os.Getenv("CMC_URL"), nil)
	if err != nil {
		return 0, err
	}

	que := url.Values{}
	que.Add("convert", To)
	que.Add("symbol", From)

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", os.Getenv("CMC_API_KEY"))
	req.URL.RawQuery = que.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("error sending request to server")
	}

	var response CMCQuoteResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return 0, err
	}
	if response.Status.ErrorCode != 0 {
		return 0, errors.New(response.Status.ErrorMessage)
	}
	rate := response.Data[string(From)][0].Quote[string(To)].Price
	if rate <= 0.0 {
		return 0, fmt.Errorf("invalid price feed")
	}
	return rate, nil
}
