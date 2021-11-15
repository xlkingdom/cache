package user_limit_token

import (
	"sync"
	"time"
)

type TokenBucket struct {
	Rate         int64 //固定的token放入速率, r/s
	Capacity     int64 //桶的容量
	Tokens       int64 //桶中当前token数量
	LastTokenSec int64 //桶上次放token的时间戳 s

	lock sync.Mutex
}

func (l *TokenBucket) Allow() bool {
	l.lock.Lock()
	defer l.lock.Unlock()
	now := time.Now().Unix()
	l.Tokens = l.Tokens + (now-l.LastTokenSec)*l.Rate // 先添加令牌
	if l.Tokens > l.Capacity {
		l.Tokens = l.Capacity
	}
	if (now - l.LastTokenSec) != 0 {
		l.LastTokenSec = now
	}
	if l.Tokens > 0 {
		// 还有令牌，领取令牌
		l.Tokens--
		return true
	} else {
		// 没有令牌,则拒绝
		return false
	}
}

func (l *TokenBucket) Set(r, c int64) {
	l.Rate = r
	l.Capacity = c
	l.Tokens = 0
	l.LastTokenSec = time.Now().Unix()
}
