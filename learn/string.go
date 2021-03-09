package learn

import "fmt"

func String() {
	str := "中国"
	fmt.Println(len(str))
	for i := 0; i < len(str); i++ {
		fmt.Println(i, str[i])
	}
	for i, c := range str {
		fmt.Println(i, c)
	}
}

func cs() {
	const (
		zz = "zz"
		zzz
	)
	fmt.Println(zz)
	fmt.Println(zzz)

	var i byte = 65
	// var i int32 = 65
	fmt.Println(string(i))

}

func f(n int) (r int) {
	defer func() {
		r += n
		recover()
	}()

	var f func()

	defer f()

	f = func() {
		r += 2
	}
	return n + 1
}

func Evaluation() {
	fmt.Println(f(3))

	var a int8 = 5
	fmt.Printf("a = %b, ^a = %b, ^^a = %b\n", a, ^a, ^^a)
}
