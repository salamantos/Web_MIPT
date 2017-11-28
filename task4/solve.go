package main


func RemoveEven(array []int) (result []int) {
	for _, elem := range array {
		if elem%2 == 0 {
			result = append(result, elem)
		}
	}
	return
}

func PowerGenerator(n int) (func() int) {
	var previousRes = 1
	return func() (res int) {
		res = previousRes * n
		previousRes = res
		return
	}
}

func DifferentWordsCount(str string) (res int) {
	var word = ""
	set := make(map[string]bool)
	res = 0

	for _, symb := range str + " " {
		if unicode.IsLetter(symb) {
			word += string(unicode.ToLower(symb))
		} else if word != "" {
			if !set[word] {
				res += 1
			}
			set[word] = true
			word = ""
		}
	}
	return
}

//func main() {
//	//input := []int{0, 3, 2, 5}
//	//result := RemoveEven(input)
//	//fmt.Println(input, result) // Должно напечататься [3 5]
//
//	//gen := PowerGenerator(2)
//	//fmt.Println(gen()) // Должно напечатать 2
//	//fmt.Println(gen()) // Должно напечатать 4
//	//fmt.Println(gen()) // Должно напечатать 8
//
//	//fmt.Println(DifferentWordsCount("Hello, world!HELLO  wOrlD...12")) // Должно напечатать 2
//}
