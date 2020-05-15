/*
ここに、パッケージのドキュメントを書く。句点までが概要として表示される。

アルファベットの大文字で始まり、句読点を含まない1行だけの段落があれば、ヘッダとして装飾される。

Headerです
*/
package fizzbuzz

import "fmt"

// 関数に対するdoc
func FizzBuzz(i int) {

	if i%15 == 0 {
		fmt.Println("FizzBuzz")
	} else if i%3 == 0 {
		fmt.Println("Fizz")
	} else if i%5 == 0 {
		fmt.Println("Buzz")
	} else {
		fmt.Println(i)
	}
}

// func main() {

// 	// forに対するdoc
// 	for i := 1; i <= 30; i++ {
// 		FizzBuzz(i)
// 	}
// }
