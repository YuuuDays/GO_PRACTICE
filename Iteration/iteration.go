package iteration

// func main() {
// 	ans := Repeat("x")
// 	fmt.Println(ans)
// }
const RepeatCount = 5

func Repeat(character string) string {
	var repeated string
	for i := 0; i < RepeatCount; i++ {
		repeated += character
	}
	return repeated
}
