package shared

import (
	"errors"
	"reflect"
)

type BatchReq struct {
	size    int
	request interface{}
}

func NewBatch(size int, request interface{}) BatchReq {
	return BatchReq{
		size:    size,
		request: request,
	}
}

func (b BatchReq) MustBatchingReq() []interface{} {
	return b.mustBatchingReq(reflect.ValueOf(b.request))
}

func (b BatchReq) BatchingReq() ([]interface{}, error) {
	val := reflect.ValueOf(b.request)
	if val.Kind() != reflect.Slice {
		return nil, errors.New("request must be in slice form")
	}

	return b.mustBatchingReq(val), nil
}

func (b BatchReq) MustBatchingReqInts() [][]int {
	return b.mustBatchingReqInts(b.request.([]int))
}

func (b BatchReq) BatchingSize(batchNum int) []int {
	var (
		batchSize []int
		i         int
	)

	if batchNum > b.size {
		return []int{b.size}
	}

	for i = batchNum; i <= b.size; i += batchNum {
		batchSize = append(batchSize, batchNum)
	}

	if b.size%batchNum > 0 {
		sub := b.size - (batchNum * len(batchSize))
		batchSize = append(batchSize, sub)
	}

	return batchSize
}

func (b BatchReq) mustBatchingReqInts(val []int) [][]int {
	var (
		result [][]int
		page   int
	)

	n := len(b.request.([]int))
	for low := 0; low < n; low = page * b.size {
		high := (page + 1) * b.size
		if high > n {
			high = n
		}
		result = append(result, val[low:high])
		page++
	}

	return result
}

func (b BatchReq) mustBatchingReq(val reflect.Value) []interface{} {
	var (
		result []interface{}
		page   int
	)

	n := val.Len()
	for low := 0; low < n; low = page * b.size {
		high := (page + 1) * b.size
		if high > n {
			high = n
		}
		result = append(result, val.Slice(low, high).Interface())
		page++
	}

	return result
}
