package chaoslib_service_api

type jsonResult struct {
	Error            bool    `json:"error"`
	ErrorDescription string  `json:"error_description,omitempty"`
	ResultInt64      int64   `json:"result_int64,omitempty"`
	ResultUint64     uint64  `json:"result_uint64,omitempty"`
	ResultInt32      int32   `json:"result_int32,omitempty"`
	ResultUint32     uint32  `json:"result_uint32,omitempty"`
	ResultInt        int     `json:"result_int,omitempty"`
	ResultFloat64    float64 `json:"result_float64,omitempty"`
	ResultFloat32    float32 `json:"result_float32,omitempty"`
	ResultInt64Array []int64 `json:"result_int64_array,omitempty"`
	ResultIntArray   []int   `json:"result_int_array,omitempty"`
	ResultInt2Array  [][]int `json:"result_int2_array,omitempty"`
}

type Rand interface {
	Uint64() (uint64, error)
	Int63() (int64, error)
	Uint32() (uint32, error)
	Int31() (int32, error)
	Int() (int, error)
	Int63n(n int64) (int64, error)
	Int31n(n int32) (int32, error)
	Intn(n int) (int, error)
	Float64() (float64, error)
	Float32() (float32, error)
	Perm(n int) ([]int, error)
}

type CrashProvider interface {
	Crash() (int64, error)
}

type WeightProvider interface {
	Weight(count int) ([]int64, error)
}

type SlotsProvider interface {
	Slots(reels []int, count int) ([][]int, error)
}

type PenaltyProvider interface {
	Penalty() (int32, error)
}
