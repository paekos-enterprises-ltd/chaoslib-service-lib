package chaoslib_service_api

import (
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"time"
)

type remoteChaosLibSource struct {
	client http.Client
	url    string
}

func (r *remoteChaosLibSource) Seed(seed int64) {}

func NewChaosLibRemoteSource(url string) rand.Source {
	res := &remoteChaosLibSource{
		client: http.Client{
			Timeout: time.Second * 5,
		},
		url: url,
	}
	return res
}

func (r *remoteChaosLibSource) Int63() int64 {
	resp, err := r.client.Get(r.url + "/api/rand/int63")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var result jsonResult
	err = json.Unmarshal(body, &result)
	if err != nil {
		panic(err)
	}
	if result.Error {
		panic(result.ErrorDescription)
	}
	return result.ResultInt64
}
