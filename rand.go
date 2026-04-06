package chaoslib_service_api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type remoteChaosLibRand struct {
	client http.Client
	url    string
}

func (r *remoteChaosLibRand) Int63() (int64, error) {
	resp, err := r.client.Get(r.url + "/api/rand/int63")
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

func (r *remoteChaosLibRand) Uint64() (uint64, error) {
	resp, err := r.client.Get(r.url + "/api/rand/uint64")
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
	return result.ResultUint64, nil
}

func (r *remoteChaosLibRand) Uint32() (uint32, error) {
	resp, err := r.client.Get(r.url + "/api/rand/uint32")
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
	return result.ResultUint32, nil
}

func (r *remoteChaosLibRand) Int31() (int32, error) {
	resp, err := r.client.Get(r.url + "/api/rand/int31")
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

func (r *remoteChaosLibRand) Int() (int, error) {
	resp, err := r.client.Get(r.url + "/api/rand/int")
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
	return result.ResultInt, nil
}

func (r *remoteChaosLibRand) Int63n(n int64) (int64, error) {
	url := fmt.Sprintf("%s/api/rand/int63n?n=%d", r.url, n)
	resp, err := r.client.Get(url)
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

func (r *remoteChaosLibRand) Int31n(n int32) (int32, error) {
	url := fmt.Sprintf("%s/api/rand/int31n?n=%d", r.url, n)
	resp, err := r.client.Get(url)
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

func (r *remoteChaosLibRand) Intn(n int) (int, error) {
	url := fmt.Sprintf("%s/api/rand/intn?n=%d", r.url, n)
	resp, err := r.client.Get(url)
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
	return result.ResultInt, nil
}

func (r *remoteChaosLibRand) Float64() (float64, error) {
	resp, err := r.client.Get(r.url + "/api/rand/float64")
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
	return result.ResultFloat64, nil
}

func (r *remoteChaosLibRand) Float32() (float32, error) {
	resp, err := r.client.Get(r.url + "/api/rand/float32")
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
	return result.ResultFloat32, nil
}

func (r *remoteChaosLibRand) Perm(n int) ([]int, error) {
	url := fmt.Sprintf("%s/api/rand/perm?n=%d", r.url, n)
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
	return result.ResultIntArray, nil
}

func NewChaosLibRand(url string) Rand {
	res := &remoteChaosLibRand{
		client: http.Client{
			Timeout: time.Second * 5,
		},
		url: url,
	}
	return res
}
