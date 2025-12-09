package selectPRA

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(a, b string) (winner string, error error) {
	// <- チャネルから受け取るという意味
	// <-は終了通知を送るまで待つという意味
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(10 * time.Second):
		return "", fmt.Errorf("timed out waiting ro %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	//チャネル　ゴルーチン同士の値を渡すために使う
	ch := make(chan struct{})
	//ゴルーチン
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)

}

// 	// URLを取得する直前の時間を記録する
// 	startA := time.Now()
// 	// URLをコンテンツを取得する
// 	http.Get(a)
// 	aDuration := time.Since(startA)

// 	startB := time.Now()
// 	http.Get(b)
// 	bDuration := time.Since(startB)

// 	if aDuration < bDuration {
// 		return a
// 	}

// 	return b
// }
