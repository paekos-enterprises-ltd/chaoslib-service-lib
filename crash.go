package chaoslib_service_api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

type remoteCrashProvider struct {
	client http.Client
	url    string
}

func (r *remoteCrashProvider) Crash() (int64, error) {
	resp, err := r.client.Get(r.url + "/api/crash")
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
	return result.ResultInt64, nil
}

func NewCrashProvider(url string) CrashProvider {
	res := &remoteCrashProvider{
		client: http.Client{
			Timeout: time.Second * 5,
		},
		url: url,
	}
	return res
}
