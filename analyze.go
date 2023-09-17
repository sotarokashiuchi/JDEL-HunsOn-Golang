package main

import "fmt"

// 入力を受付、入力された整数の個数に合わせて各分析用関数を呼び出す関数。
func main() {
	var inputNumber []int
	var inputNumberTotal int

	// input
	fmt.Printf("解析する整数の個数を入力してください>>>")
	fmt.Scan(&inputNumberTotal)
	for i := 0; i < inputNumberTotal; i++ {
		var input int
		fmt.Printf("%vつ目の整数を入力してください>>>", i+1)
		fmt.Scan(&input)
		inputNumber = append(inputNumber, input)
	}

	fmt.Println("")
	fmt.Println("Analyze....")
	switch inputNumberTotal {
	case 1:
		isOddNumber(inputNumber[0])
		isEvenNumber(inputNumber[0])
		isPrimeNumber(inputNumber[0])
		isPerfectNumber(inputNumber[0])
	case 2:
		isAmicableNumbers(inputNumber[0], inputNumber[1])

	}
}


// 奇数判定関数。与えられた引数が奇数の場合、その趣旨を表示する
// 奇数とは、「2で割り切れない整数」
// ex) 1,3,5,7,...
func isOddNumber(number int) {
	return
}

// 偶数判定関数。与えられた引数が偶数の場合、その趣旨を表示する
// 偶数とは、「2で割り切れる整数」
// ex) 2,4,6,8,...
func isEvenNumber(number int) {
	return
}

// 素数判定関数。与えられた引数が素数の場合、その趣旨を表示する
// 素数とは「1を除く約数が1とその数自身だけである自然数」
// ex) 2,3,5,7,...
func isPrimeNumber(number int) {
	return
}

// 完全数判定関数。与えられた引数が完全数の場合、その趣旨を表示する
// 完全数とは「自身を除く約数の和が自身に等しい数の自然数」
// ex) 6,28,496,...
func isPerfectNumber(number int) {
	return
}

// 親和数判定関数。与えられた引数が親和数の場合、その趣旨を表示する
// 親和数とは「自身を除く約数の和が、他方に等しい自然数の組」
// ex) 220と284, 1184と1210, 2620と2924,...
func isAmicableNumbers(x, y int) {
	return
}

// 平方数。与えられた引数が平方数の場合、その趣旨を表示する
func isSquareNumber(number int) {
	return
}

// メルセンヌ数判定関数。与えられた引数がメルセンス数の場合、その趣旨を表示する
// メルセンヌ数とは「2^n-1で表される数」
// ex) 1,3,7,31|
func isMersenneNumber(number int){
	return
}

