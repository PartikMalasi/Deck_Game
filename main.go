package main

import "fmt"

func main() {
	var a int = 10
	b := 20
	fmt.Println(a, b)
	fmt.Println(getNumber(22))

	cards := deck{"Ace", "Ten"}

	fmt.Println(cards)
	cards = append(cards, "new card")
	fmt.Println(cards)

	// for i, c := range cards {
	// 	fmt.Println(i, c)
	// }

	cards.print()
}

func getNumber(temp int) int {
	return temp * 2
}
