package selectPRA

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	aDuration := measureResponseTime(a)
	bDuration := measureResponseTime(b)
	if aDuration < bDuration {
		return a
	}
	return b
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
