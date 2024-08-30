package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan() //1行分スキャン
	inputs := strings.Split(sc.Text(), " ")

	num, _ := strconv.Atoi(inputs[0])
	discount, _ := strconv.Atoi(inputs[1])

	sc.Scan()
	inputs2 := strings.Split(sc.Text(), " ")

	ans := CalculateTotal(num, discount, inputs2)

	// 答え
	fmt.Println(ans)

}

/*
リファクタリング前↓
自分で書いたコード
benchMark結果:

7310925               181.3 ns/op           248 B/op          5 allocs/op
PASS
ok      github.com/YuuHikida/GO_PRACTICE/paiza_practice 1.665s
*/
// func CalculateTotal(num int, discount int, inputs2 []string) int {
// 	var ps []int

// 	var maxValue int
// 	var sub int

// 	for n, input := range inputs2 {
// 		p, _ := strconv.Atoi(input)

// 		if maxValue < p {
// 			maxValue = p
// 			sub = n
// 		}
// 		ps = append(ps, p)
// 	}

// 	if discount < maxValue {
// 		ps[sub] = maxValue / 2
// 	}

// 	var ans int
// 	//　合算
// 	for i := 0; i < num; i++ {
// 		ans += ps[i]
// 	}
// 	return ans
// }

/*
	リファクタリング後(chatGPT)
	BenchmarkRepeat-16
	7322055               162.6 ns/op           248 B/op          5 allocs/op
	PASS
	ok      github.com/YuuHikida/GO_PRACTICE/paiza_practice 1.528s
*/

func CalculateTotal(num int, discount int, inputs2 []string) int {
	var maxValue int
	var ans int

	for n := 0; n < num; n++ {
		p, _ := strconv.Atoi(inputs2[n])

		if p > maxValue {
			maxValue = p
		}
		ans += p
	}

	if discount < maxValue {
		ans -= maxValue     // 最大値を引く
		ans += maxValue / 2 // 割引後の値を足す
	}

	return ans
}
