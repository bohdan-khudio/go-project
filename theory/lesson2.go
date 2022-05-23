package main

func main() {

	/*

		Go має лише одну циклічну конструкцію — цикл for.

		Основний цикл for складається з трьох компонентів, розділених крапкою з комою:

		оператор init: виконується перед першою ітерацією
		вираз умови: оцінюється перед кожною ітерацією
		оператор post: виконується в кінці кожної ітерації
		Інструкція init часто є короткою декларацією змінної, а змінні, оголошені там, видимі лише в області дії оператора for.

		Цикл припинить ітерацію, коли логічна умова стане хибною.

		Примітка: на відміну від інших мов, таких як C, Java або JavaScript,
		три компоненти оператора for не містять круглих дужок, а фігурні дужки { } завжди потрібні.

		Перший і останній оператори є необов'язковими.
		У цей момент ви можете відкинути крапку з комою: C while пишеться як for у Go.

	*/

	//sum := 0
	//for i := 0; i < 10; i++ {
	//	sum += i
	//}
	//fmt.Println(sum)

	//sum := 1
	//for sum < 1000 {
	//	sum += sum
	//}
	//fmt.Println(sum)

	//for {
	//}

	/*
		if оператори в Go подібні до циклів for; вираз необов'язково оточувати круглими дужками ( ), але дужки { } є обов'язковими.
	*/

	//fmt.Println(sqrt(2), sqrt(-4))

	/*

		if з коротким оператором
		Як і у випадку for, оператор if може починатися з короткого оператора, який слід виконати перед умовою.

		Змінні, оголошені оператором, знаходяться лише в області дії до кінця if.
	*/

	//fmt.Println(
	//	pow(3, 2, 10),
	//	pow(3, 3, 20),
	//)

	/*

		Оператор switch — це коротший спосіб написати послідовність операторів if - else.
		Він запускає перший випадок, значення якого дорівнює виразу умови.

		Перемикач Go подібний до перемикача в C, C++, Java, JavaScript і PHP, за винятком того,
		що Go запускає лише вибраний регістр, а не всі наступні.
		Фактично, оператор break, який необхідний у кінці кожного випадку цими мовами, в Go додається автоматично.

	*/

	//fmt.Print("Go runs on ")
	//os := runtime.GOOS;
	//switch os {
	//case "darwin":
	//	fmt.Println("OS X.")
	//case "linux":
	//	fmt.Println("Linux.")
	//default:
	//	// freebsd, openbsd,
	//	// plan9, windows...
	//	fmt.Printf("%s.\n", os)
	//}

	//fmt.Println("When's Saturday?")
	//today := time.Now().Weekday()
	//
	//switch time.Wednesday {
	//case today + 0:
	//	fmt.Println("Today.")
	//case today + 1:
	//	fmt.Println("Tomorrow.")
	//case today + 2:
	//	fmt.Println("In two days.")
	//default:
	//	fmt.Println("Too far away.")
	//}

	//t := time.Now()
	//switch {
	//case t.Hour() < 12:
	//	fmt.Println("Good morning!")
	//case t.Hour() < 17:
	//	fmt.Println("Good afternoon.")
	//default:
	//	fmt.Println("Good evening.")
	//}

	/*
		Оператор defer відкладає виконання функції, до виходу з функції через return.

		Аргументи defer виклику оцінюються негайно, але виклик функції не виконується до return.

		https://gobyexample.com/defer#:~:text=Defer%20is%20used%20to%20ensure,be%20used%20in%20other%20languages.&text=Suppose%20we%20wanted%20to%20create,close%20when%20we're%20done.

		https://go.dev/blog/defer-panic-and-recover

		Відкладені виклики функцій поміщаються в стек. Коли функція повертається,
		її відкладені виклики виконуються в порядку last-in-first-out.
	*/

	//defer fmt.Println("world")
	//
	//fmt.Println("hello")

	//fmt.Println("counting")
	//
	//for i := 0; i < 10; i++ {
	//	defer fmt.Println(i)
	//}
	//
	//fmt.Println("done")

	/*

		Go має вказівники. Вказівник містить адресу пам'яті значення.

		Тип *T є вказівником на значення T. Його нульове значення дорівнює nil.

		var p *int

		Оператор & створює вказівник на свій операнд.

		i := 42
		p = &i

		Оператор * позначає значення вказівника.

		fmt.Println(*p) // читаємо i через вказівник p
		*p = 21 // встановлюємо i через вказівник p
	*/

	//i, j := 42, 2701
	//
	//p := &i         // point to i
	//fmt.Println(*p) // read i through the pointer
	//*p = 21         // set i through the pointer
	//fmt.Println(i)  // see the new value of i
	//
	//p = &j         // point to j
	//*p = *p / 37   // divide j through the pointer
	//fmt.Println(j) // see the new value of j

	//fmt.Println(Vertex{1, 2})

	//v := Vertex{1, 2}
	//v.X = 4
	//fmt.Println(v)

	/*

		Доступ до полів структури можна отримати через вказівник структури.

		Щоб отримати доступ до поля X структури, коли у нас є вказівник на структуру p,
		ми можемо написати (*p).X. Однак це позначення є громіздким,
		тому мова дозволяє нам замість цього писати лише p.X, без явного розіменування.

	*/

	//v := Vertex{1, 2}
	//p := &v
	//p.X = 1e9
	//fmt.Println(v)

	//var (
	//	v1 = Vertex{1, 2}  // has type Vertex
	//	v2 = Vertex{X: 1}  // Y:0 is implicit
	//	v3 = Vertex{}      // X:0 and Y:0
	//	p  = &Vertex{1, 2} // has type *Vertex
	//)
	//
	//fmt.Println(v1, p, v2, v3)

}

//func sqrt(x float64) string {
//	if x < 0 {
//		return sqrt(-x) + "i"
//	}
//	return fmt.Sprint(math.Sqrt(x))
//}

//func pow(x, n, lim float64) float64 {
//
//	if v := math.Pow(x, n); v < lim {
//		return v
//	} else if v > lim {
//
//	}
//	return lim
//}

//type Vertex struct {
//	X int
//	Y int
//}
