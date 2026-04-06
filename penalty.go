package chaoslib_service_api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

type remotePenaltyProvider struct {
	client http.Client
	url    string
}

func (r *remotePenaltyProvider) Penalty() (int32, error) {
	resp, err := r.client.Get(r.url + "/api/penalty")
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	var result jsonResult
	err = json.Unmarshal(body, &result)
	if err != nil {
		return 0, err
	}
	if result.Error {
		return 0, errors.New(result.ErrorDescription)
	}
	return result.ResultInt32, nil
}

func NewPenaltyProvider(url string) PenaltyProvider {
	res := &remotePenaltyProvider{
		client: http.Client{
			Timeout: time.Second * 5,
		},
		url: url,
	}
	return res
}
