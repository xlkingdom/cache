package redis

import (
	"context"
	"encoding/json"
	"github.com/mizuki1412/go-core-kit/class/exception"
	"github.com/mizuki1412/go-core-kit/service/rediskit"
)

func ExistJudge(key string) int64 {
	client := rediskit.Instance()
	var (
		err  error
		base int64
	)
	base, err = client.Exists(context.Background(), key).Result()
	if err != nil {
		panic(exception.New("出现错误:" + err.Error()))
	}
	return base
}

func Get(key string) []map[string]interface{} {
	client := rediskit.Instance()
	var (
		dat []map[string]interface{}
		err error
		val []byte
	)
	val, err = client.Get(context.Background(), key).Bytes()
	_ = json.Unmarshal(val, &dat)
	if err != nil {
		panic(exception.New("出现错误:" + err.Error()))
	}
	return dat
}

func Set(key string, val interface{}) {
	rediskit.Set(context.Background(), key, val, 15000000000)
}
