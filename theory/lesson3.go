package main

func main() {

	/*

		Масиви
		Тип [n]T — це масив із n значень типу T.

		Вираз

		var a [10]int

		оголошує змінну a як масив із десяти цілих чисел.

		Розміри масивів не можна змінювати, але не хвилюйтеся;
		Go забезпечує зручний спосіб роботи з масивами.

	*/

	//var a [2]string
	//a[0] = "Hello"
	//a[1] = "World"
	//fmt.Println(a[0], a[1])
	//fmt.Println(a)
	//
	//primes := [6]int{2, 3, 5, 7, 11, 13}
	//fmt.Println(primes)

	/*
		Слайс
		Масив має фіксований розмір. З іншого боку, слайс — це гнучкий перегляд
		елементів масиву. На практиці слайси зустрічаються набагато частіше, ніж масиви.

		Тип []T є слайсом з елементами типу T.

		Слайс формується шляхом вказуванням двох індексів, нижньої та високої меж, розділених двокрапкою:

		a[low : high]
		Це визначає напіввідкритий діапазон, який включає перший елемент, але виключає останній.

		Наступний вираз створює фрагмент, який включає елементи з 1 по 3:

		a[1:4]

	*/

	//primes := [6]int{2, 3, 5, 7, 11, 13}
	//var s []int = primes[1:4]
	//fmt.Println(s)

	//names := [4]string{
	//	"John",
	//	"Paul",
	//	"George",
	//	"Ringo",
	//}
	//fmt.Println(names)
	//
	//a := names[0:2]
	//b := names[1:3]
	//fmt.Println(a, b)
	//
	//b[0] = "XXX"
	//fmt.Println(a, b)
	//fmt.Println(names)

	/*

		Літерал слайсу подібний літералу масиву без вказання довжини.

		Це літерал масиву:

		[3]bool{true, true, false}

		А цей створює той самий масив, що й вище, а потім створює слайс,
		який посилається на нього:

		[]bool{true, true, false}

	*/

	//q := []int{2, 3, 5, 7, 11, 13}
	//fmt.Println(q)
	//
	//r := []bool{true, false, true, true, false, true}
	//fmt.Println(r)
	//
	//s := []struct {
	//	i int
	//	b bool
	//}{
	//	{2, true},
	//	{3, false},
	//	{5, true},
	//	{7, true},
	//	{11, false},
	//	{13, true},
	//}
	//fmt.Println(s)

	//s := []int{2, 3, 5, 7, 11, 13}
	//
	//fmt.Println(s[1:4])
	//fmt.Println(s[:2])
	//fmt.Println(s[1:])

	/*
		Слайс має довжину і місткість.

		Довжина слайсу — це кількість елементів, які він містить.

		Місткість слайсу — це кількість елементів у базовому масиві, рахуючи від першого
		елемента в слайсі.

		Довжину та місткість слайсу s можна отримати за допомогою виразів len(s) і cap(s).
	*/

	//s := []int{2, 3, 5, 7, 11, 13}
	//printSlice(s)
	//
	//// Slice the slice to give it zero length.
	//s = s[:0]
	//printSlice(s)
	//
	//// Extend its length.
	//s = s[:4]
	//printSlice(s)
	//
	//// Drop its first two values.
	//s = s[2:]
	//printSlice(s)

	//Нульовий слайс
	//var s []int
	//fmt.Println(s, len(s), cap(s))
	//if s == nil {
	//	fmt.Println("nil!")
	//}

	// Слайси можуть включати будь-який інший тип, в тому числі інші слайси

	// Create a tic-tac-toe board.
	//board := [][]string{
	//	[]string{"_", "_", "_"},
	//	[]string{"_", "_", "_"},
	//	[]string{"_", "_", "_"},
	//}
	//
	//// The players take turns.
	//board[0][0] = "X"
	//board[2][2] = "O"
	//board[1][2] = "X"
	//board[1][0] = "O"
	//board[0][2] = "X"
	//
	//for i := 0; i < len(board); i++ {
	//	fmt.Printf("%s\n", strings.Join(board[i], " "))
	//}

	/*
		https://go.dev/blog/slices-intro
	*/

	//var s []int
	//printSlice(s)
	//
	//// append works on nil slices.
	//s = append(s, 0)
	//printSlice(s)
	//
	//// The slice grows as needed.
	//s = append(s, 1)
	//printSlice(s)
	//
	//// We can add more than one element at a time.
	//s = append(s, 2, 3, 4)
	//printSlice(s)

	// Створення слайсу з make

	//a := make([]int, 5)
	//printSlice("a", a)
	//
	//b := make([]int, 0, 5)
	//printSlice("b", b)
	//
	//c := b[:2]
	//printSlice("c", c)
	//
	//d := c[2:5]
	//printSlice("d", d)

	// Range
	//var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	//
	//for i, v := range pow {
	//	fmt.Printf("2**%d = %d\n", i, v)
	//}

	//pow := make([]int, 10)
	//for i := range pow {
	//	pow[i] = 1 << uint(i) // == 2**i
	//}
	//for _, value := range pow {
	//	fmt.Printf("%d\n", value)
	//}

	// Maps

	//type Vertex struct {
	//	Lat, Long float64
	//}

	//var m map[string]Vertex
	//
	//m = make(map[string]Vertex)
	//m["Bell Labs"] = Vertex{
	//	40.68433, -74.39967,
	//}
	//fmt.Println(m["Bell Labs"])

	//var m = map[string]Vertex{
	//	"Bell Labs": Vertex{
	//		40.68433, -74.39967,
	//	},
	//	"Google": Vertex{
	//		37.42202, -122.08408,
	//	},
	//}
	//
	//fmt.Println(m)

	/*

		Як додати або оновити елемент в мапі:
		m[key] = elem

		Отримати елемент:
		elem = m[key]

		Видалити елемент:
		delete(m, key))

		Перевірте наявність ключа з двозначним призначенням:
		elem, ok = m[key]

		Якщо ключ знаходиться в m, це правда. Якщо ні, ок — неправда.

		Якщо ключа немає в мапі, то elem – це нульове значення для типу елемента мапа.
	*/

	//m := make(map[string]int)
	//
	//m["Answer"] = 42
	//fmt.Println("The value:", m["Answer"])
	//
	//m["Answer"] = 48
	//fmt.Println("The value:", m["Answer"])
	//
	//delete(m, "Answer")
	//fmt.Println("The value:", m["Answer"])
	//
	//v, ok := m["Answer"]
	//fmt.Println("The value:", v, "Present?", ok)

	/*
		Функції як значення

		Функції також є значеннями. Їх можна передавати так само як і інші значення.

		Функції як значення можуть використовуватися як аргументи функції та повертаються як значення.
	*/

	//hypot := func(x, y float64) float64 {
	//	return math.Sqrt(x*x + y*y)
	//}
	//fmt.Println(hypot(5, 12))
	//
	//fmt.Println(compute(hypot))
	//fmt.Println(compute(math.Pow))

	//pos, neg := adder(), adder()
	//for i := 0; i < 10; i++ {
	//	fmt.Println(
	//		pos(i),
	//		neg(-2*i),
	//	)
	//}

}

//func printSlice(s []int) {
//	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
//}

//func compute(fn func(float64, float64) float64) float64 {
//	return fn(3, 4)
//}

//func adder() func(int) int {
//	sum := 0
//	return func(x int) int {
//		sum += x
//		return sum
//	}
//}
