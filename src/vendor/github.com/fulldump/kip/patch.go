package kip

type Patch struct {
	Operation string      `json:"operation"`
	Key       string      `json:"key"`
	Value     interface{} `json:"value"`
}
