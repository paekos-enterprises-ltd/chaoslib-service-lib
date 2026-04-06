package chaoslib_service_api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type remoteSlotsProvider struct {
	client http.Client
	url    string
}

func (r *remoteSlotsProvider) Slots(reels []int, count int) ([][]int, error) {
	reelsStr := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(reels)), ","), "[]")
	url := fmt.Sprintf("%s/api/slots?reels=%s", r.url, reelsStr)
	if count > 0 {
		url = fmt.Sprintf("%s&count=%d", url, count)
	}

	resp, err := r.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result jsonResult
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	if result.Error {
		return nil, errors.New(result.ErrorDescription)
	}
	if count > 0 {
		return result.ResultInt2Array, nil
	}
	return [][]int{result.ResultIntArray}, nil
}

func NewSlotsProvider(url string) SlotsProvider {
	res := &remoteSlotsProvider{
		client: http.Client{
			Timeout: time.Second * 5,
		},
		url: url,
	}
	return res
}
