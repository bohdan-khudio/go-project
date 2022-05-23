package main

import (
	"fmt"
)

/*

https://go.dev/doc/install

*/

//var c, python, java bool
//var i, j int = 1, 2

//var (
//	ToBe   bool       = false
//	MaxInt uint64     = 1<<64 - 1
//	z      complex128 = cmplx.Sqrt(-5 + 12i)
//)

//const Pi = 3.14

func main() {
	fmt.Println("Welcome to the Grass Business Labs!")

	//fmt.Println("The time is", time.Now())

	//now := time.Now()
	//fmt.Println(now.Format("Mon Jan 15:04:05 2006"))
	//fmt.Println(now.Format(time.Kitchen))

	/*

		Пакети

		Кожна програма Go складається з пакетів.

		Програми запускаються в пакеті main.

		Ця програма використовує пакунки з шляхами імпорту "fmt" і "math/rand".

		За домовленістю ім’я пакета збігається з останнім елементом шляху імпорту.
		Наприклад, пакет "math/rand" містить файли, які починаються з пакета операторів rand.
	*/

	//fmt.Println("My favorite number is", rand.Intn(10))

	/*

		Імпорт

		Цей код групує імпорт у круглі дужки, «факторні» імпортні оператори.

		Ви також можете написати кілька операторів імпорту, наприклад:

			import "fmt"
			import "math"

		Але добре використовувати оператор факторного імпорту.

		https://pkg.go.dev/fmt

	*/

	//fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))

	//fmt.Println(math.Pi)

	/*
		Функції

		Функція може приймати нуль або більше аргументів.

		У цьому прикладі add приймає два параметри типу int.

		Зверніть увагу, що тип йде після імені змінної.

		https://go.dev/blog/declaration-syntax

		Якщо два або більше послідовних іменованих параметрів функції мають
		спільний тип, ви можете опустити тип з усіх, крім останнього.

		func add(x, y int) int
	*/

	//fmt.Println(add(42, 13))

	/*
		Функція може повертати будь-яку кількість результатів.

		Функція swap() повертає два рядки.
	*/

	//a, b := swap("hello", "world")
	//fmt.Println(a, b)

	/*

		Іменовані значення функій

		Значення в Go можна іменувати. Вони розглядаються як змінні, визначені у верхній частині функції.

		Оператор return без аргументів повертає іменовані значення.

	*/

	//fmt.Println(split(17))

	/*
		Змінні

		Оператор var оголошує список змінних; як і в списках аргументів функції, тип є останнім.

		Оператор var може бути на рівні пакета або функції. У цьому прикладі ми бачимо і те, і інше.
	*/

	//var i int
	//fmt.Println(i, c, python, java)

	//var c, python, java = true, false, "no!"
	//fmt.Println(i, j, c, python, java)

	/*
		Усередині функції замість оголошення var з неявним типом можна використовувати короткий оператор присвоєння :=.

		За межами функції кожен оператор починається з ключового слова (var, func тощо), тому конструкція := недоступна.
	*/

	//var i, j int = 1, 2
	//k := 3
	//c, python, java := true, false, "no!"
	//
	//fmt.Println(i, j, k, c, python, java)

	/*
		Базовими типами в Go є наступні

		bool

		string

		int  int8  int16  int32  int64
		uint uint8 uint16 uint32 uint64 uintptr

		byte // alias for uint8

		rune // alias for int32
			 // represents a Unicode code point

		float32 float64

		complex64 complex128

		Типи int, uint і uintptr зазвичай мають ширину 32 біти в 32-розрядних системах і 64 біти в 64-розрядних системах.
		Якщо вам потрібне ціле значення, ви повинні використовувати int якщо немає конкретної причини використовувати інше

	*/

	//fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	//fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	//fmt.Printf("Type: %T Value: %v\n", z, z)

	/*

		Дефолтні значення

		Змінні, оголошені без явного початкового значення, отримують дефолтне значення.

		0 для числових типів,
		false для логічного типу і
		"" (порожній рядок) для рядків.

	*/

	//var i int
	//var f float64
	//var b bool
	//var s string
	//fmt.Printf("%v %v %v %q\n", i, f, b, s)

	/*

		Вираз T(v) перетворює значення v на тип T.

		Деякі числові перетворення:

		var i int = 42
		var f float64 = float64(i)
		var u uint = uint(f)

		Або, простіше:

		i := 42
		f := float64(i)
		u := uint(f)

		На відміну від C, у Go призначення між елементами різного типу вимагає явного перетворення.
	*/

	//var x, y int = 3, 4
	//var f float64 = math.Sqrt(float64(x*x + y*y))
	//var z uint = uint(f)
	//fmt.Println(x, y, z)

	/*

		Константи

		Константи оголошуються як змінні, але з ключовим словом const.

		Константи можуть бути символьними, рядковими, булевими або числовими значеннями.

		Константи не можна оголошувати за допомогою синтаксису :=
	*/

	//const World = "世界"
	//fmt.Println("Hello", World)
	//fmt.Println("Happy", Pi, "Day")
	//
	//const Truth = true
	//fmt.Println("Go rules?", Truth)

}

//func add(x int, y int) int {
//	return x + y
//}

//func swap(x, y string) (string, string) {
//	return y, x
//}

//func split(sum int) (x, y int) {
//	x = sum * 4 / 9
//	y = sum - x
//	return
//}
