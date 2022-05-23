package main

func main() {

	/*
		Методи
		У Go немає скласів. Однак ви можете визначити методи на певних типах.

		Метод — це функція зі спеціальним аргументом отримувачем.

		Отримувач з’являється у власному списку аргументів між ключовим словом func та іменем методу.

		У цьому прикладі метод Abs має отримувач типу Vertex з ім’ям v.

		Метод - це функція з отримувачем
	*/

	//v := Vertex{3, 4}
	//fmt.Println(v.Abs())

	/*

		Ви також можете оголосити метод для неструктурованих типів.

		У цьому прикладі ми бачимо числовий тип MyFloat з методом Abs.

		Ви можете оголосити метод лише з отримувачем, тип якого визначено в тому ж пакеті, що й метод.
		Ви не можете оголосити метод за допомогою отримувача, тип якого визначено в іншому пакеті
		(який включає вбудовані типи, такі як int).

	*/

	//f := MyFloat(-math.Sqrt2)
	//fmt.Println(f.Abs())

	/*

	 */
	//
	//v := Vertex{3, 4}
	//v.Scale(10)
	//fmt.Println(v.Abs())

	/*

		В функціях якщо ми вказуємо поінтер як параметер ми маємо передати поінтер
		Для методів це не обов'язково

	*/

	//var v = Vertex{3, 4}
	//v.Scale(5) // OK
	//fmt.Println(v.Abs())
	//p := &v
	//p.Scale(10) // OK
	//fmt.Println(v.Abs())

	/*

		Є дві причини використовувати поінтер на отримувача.

		Перший полягає в тому, щоб метод міг змінити значення, на яке вказує його отримувача.

		По-друге, уникати копіювання значення під час кожного виклику методу.
		Це може бути більш ефективним, якщо отримувач є великою структурою,

	*/

	/*
		Інтерфейси
		Тип інтерфейсу визначається як набір сигнатур методів.

		Примітка. У прикладі коду в рядку 22 є помилка. Vertex (тип значення)
		не реалізує Abser, оскільки метод Abs визначено лише для *Vertex (тип поінтер).
	*/

	//var a Abser
	//f := MyFloat(-math.Sqrt2)
	//v := Vertex{3, 4}
	//
	//a = f  // a MyFloat implements Abser
	//a = &v // a *Vertex implements Abser
	//
	//// In the following line, v is a Vertex (not *Vertex)
	//// and does NOT implement Abser.
	//a = v
	//
	//fmt.Println(a.Abs())

	//var i I = T{"hello"}
	//i.M()

	/*

		Метод можна викликати на nil отримувачі

	*/

	//var i I
	//
	//var t *T
	//i = t
	//describe(i)
	//i.M()
	//
	//i = &T{"hello"}
	//describe(i)
	//i.M()

	/*
			Стрингери
			Одним з найбільш поширених інтерфейсів є Stringer, визначений пакетом fmt.

			type Stringer interface {
		    String() string
			}

			Stringer — це тип, який може описати себе як рядок.
			Пакет fmt (і багато інших) використовують цей інтерфейс для друку значень.
	*/

	//a := Person{"Arthur Dent", 42}
	//z := Person{"Zaphod Beeblebrox", 9001}
	//fmt.Println(a, z)

}

//type Vertex struct {
//	X, Y float64
//}

//
//func (v Vertex) Abs() float64 {
//	return math.Sqrt(v.X*v.X + v.Y*v.Y)
//}

//type MyFloat float64

//func (f MyFloat) Abs() float64 {
//	if f < 0 {
//		return float64(-f)
//	}
//	return float64(f)
//}

//func (v *Vertex) Scale(f float64) {
//	v.X = v.X * f
//	v.Y = v.Y * f
//}

//type Abser interface {
//	Abs() float64
//}
//
//func (f MyFloat) Abs() float64 {
//	if f < 0 {
//		return float64(-f)
//	}
//	return float64(f)
//}
//
//func (v *Vertex) Abs() float64 {
//	return math.Sqrt(v.X*v.X + v.Y*v.Y)
//}

//type I interface {
//	M()
//}
//
//type T struct {
//	S string
//}

//
//// This method means type T implements the interface I,
//// but we don't need to explicitly declare that it does so.
//func (t T) M() {
//	fmt.Println(t.S)
//}

//func (t *T) M() {
//	if t == nil {
//		fmt.Println("<nil>")
//		return
//	}
//	fmt.Println(t.S)
//}
//
//func describe(i I) {
//	fmt.Printf("(%v, %T)\n", i, i)
//}

//type Person struct {
//	Name string
//	Age  int
//}
//
//func (p Person) String() string {
//	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
//}
