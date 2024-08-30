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

	ans := CaluculateTotal(num, discount, inputs2)

	// 答え
	fmt.Println(ans)

}

// リファクタリング前↓
func CaluculateTotal(num int, discount int, inputs2 []string) int {
	var ps []int

	var maxValue int
	var sub int

	for n, input := range inputs2 {
		p, _ := strconv.Atoi(input)

		if maxValue < p {
			maxValue = p
			sub = n
		}
		ps = append(ps, p)
	}

	if discount < maxValue {
		ps[sub] = maxValue / 2
	}

	var ans int
	//　合算
	for i := 0; i < num; i++ {
		ans += ps[i]
	}
	return ans
}
