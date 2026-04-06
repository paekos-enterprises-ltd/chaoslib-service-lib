package chaoslib_service_api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type remoteWeightProvider struct {
	client http.Client
	url    string
}

func (r *remoteWeightProvider) Weight(count int) ([]int64, error) {
	url := r.url + "/api/weight"
	if count > 0 {
		url = fmt.Sprintf("%s?count=%d", url, count)
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
		return result.ResultInt64Array, nil
	}
	return []int64{result.ResultInt64}, nil
}

func NewWeightProvider(url string) WeightProvider {
	res := &remoteWeightProvider{
		client: http.Client{
			Timeout: time.Second * 5,
		},
		url: url,
	}
	return res
}
