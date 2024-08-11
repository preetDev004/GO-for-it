package main 

import "fmt"

// WITHOUT GENRICS
// Below funcs are basically doing the same logic but because of the types we need two different functions.
// func splitIntSlices(s []int) (s1,s2 []int){
// 	mid := len(s)/2
// 	return s[:mid], s[mid:]
// }
// func splitStringSlices(s []string) (s1, s2  []string){
// 	mid := len(s)/2
// 	return s[:mid], s[mid:]
// }

func splitAnySlice[T any](s []T) (t1, t2 []T){
	mid := len(s)/2
	return s[:mid], s[mid:]
}

func main() {
	// WITHOUT GENRICS
	// iSlice := []int{1,2,3,4,5,6,7,8,9,10}
	// sSlice := []string{"a","b","c","d","e","f","g","h","i","j"}

	// is1 , is2 := splitIntSlices(iSlice)
	// fmt.Println(is1,is2)
	// ss1 , ss2 := splitStringSlices(sSlice)
	// fmt.Println(ss1,ss2)


	iSlice := []int{1,2,3,4,5,6,7,8,9,10}
	sSlice := []string{"a","b","c","d","e","f","g","h","i","j"}

	is1 , is2 := splitAnySlice(iSlice)
	fmt.Println(is1,is2)
	ss1 , ss2 := splitAnySlice(sSlice)
	fmt.Println(ss1,ss2)

}