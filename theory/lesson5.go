package main

func main() {
	/*  19/26

		Помилки
		Програми Go виражають стан помилки за допомогою значень помилок.

		Тип помилки - це вбудований інтерфейс, подібний до fmt.Stringer:

		type error interface {
	    	Error() string
		}

		(Як і у випадку з fmt.Stringer, пакет fmt шукає інтерфейс помилки під час друку значень.)

		Функції часто повертають значення помилки, а код виклику повинен обробляти помилки, перевіряючи, чи дорівнює помилка нулю.

		i, err := strconv.Atoi("42")
		if err != nil {
			fmt.Printf("couldn't convert number: %v\n", err)
			return
		}
		fmt.Println("Converted integer:", i)

		Нульова помилка означає успіх; ненульова помилка означає помилку.
	*/

	//if err := run(); err != nil {
	//	fmt.Println(err)
	//}

	/* 21/26

	Рідери
	Пакет io визначає інтерфейс io.Reader, який надає функціонал для читання потоку даних.

	Стандартна бібліотека Go містить багато реалізацій цього інтерфейсу, включаючи файли,
	мережеві з’єднання, компресори, шифри та інші.

	Інтерфейс io.Reader має метод Read:

	func (T) Read(b []byte) (n int, err error)

	Read заповнює даний фрагмент байтів даними і повертає кількість заповнених байтів і значення помилки.
	Він повертає помилку io.EOF, коли потік закінчується.

	*/

	//r := strings.NewReader("Hello, Reader!")
	//
	//b := make([]byte, 8)
	//for {
	//	n, err := r.Read(b)
	//	fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
	//	fmt.Printf("b[:n] = %q\n", b[:n])
	//	if err == io.EOF {
	//		break
	//	}
	//}

	/* 24/26

	Зображення
	Пакет image визначає інтерфейс Image:

	package image

	type Image interface {
	    ColorModel() color.Model
	    Bounds() Rectangle
	    At(x, y int) color.Color
	}

	Примітка: значення Rectangle, що повертається методом Bounds, насправді є image.Rectangle,
	оскільки оголошення знаходиться всередині зображення пакета.

	(Див. https://pkg.go.dev/image#Image для всіх деталей.)

	Типи color.Color і color.Model також є інтерфейсами, але ми ігноруємо це,
	використовуючи попередньо визначені реалізації color.RGBA і color.RGBAModel.
	Ці інтерфейси та типи визначаються пакетом https://pkg.go.dev/image/color

	*/

	//m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	//fmt.Println(m.Bounds())
	//fmt.Println(m.At(0, 0).RGBA())

	/* 1/11
	goroutine — це легкий потік, керований середовищем виконання Go.

	go f(x, y, z)
	запускає нову goroutine

	f(x, y, z)
	Оцінка f, x, y і z відбувається в поточній goroutine, а виконання f відбувається в новій goroutine.
	*/

	//go say("world")
	//say("hello")

	/*  2/11
	Канали — це типізований провідник, через який ви можете надсилати та отримувати значення за допомогою оператора каналу <-.

	ch <- v // Надіслати v в канал ch.
	v := <-ch // Отримати від ch і присвоїти значення v.
	(Дані надходять у напрямку стрілки.)

	Як і мапи та слайси, канали необхідно створити перед використанням:

	ch := make(chan int)
	*/

	//s := []int{7, 2, 8, -9, 4, 0}
	//
	//c := make(chan int)
	//go sum(s[:len(s)/2], c)
	//go sum(s[len(s)/2:], c)
	//x, y := <-c, <-c // receive from c
	//
	//fmt.Println(x, y, x+y)

	/*

		Буферизовані канали.
		Канали можна буферизувати. Надайте довжину буфера як другий аргумент,
		який потрібно зробити для ініціалізації буферизованого каналу:

		ch := make(chan int, 100)

	*/

	//ch := make(chan int, 2)
	//ch <- 1
	//ch <- 2
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)

	/* 4/11

	Range і Close
	Відправник може закрити канал, щоб вказати, що більше не надсилатимуться значення.
	Приймачі можуть перевірити, чи був закритий канал

	v, ok := <-ch

	ok є хибним, якщо більше немає значень для отримання і канал закритий.

	Цикл < for i := range c > отримує значення з каналу кілька разів, поки він не буде закритий.

	Примітка: тільки відправник повинен закривати канал, але ніколи одержувач. Передача на закритий канал викличе паніку.

	Інша примітка: канали не схожі на файли; зазвичай не потрібно їх закривати.
	Закриття потрібне лише тоді, коли одержувачу потрібно повідомити, що більше не будуть надходити значення,
	наприклад, щоб завершити range цикл.

	*/

	//c := make(chan int, 10)
	//go fibonacci(cap(c), c)
	//for i := range c {
	//	fmt.Println(i)
	//}

	/* 5/11

	select оператор дозволяє goroutine чекати на кілька комунікаційних операцій
	select блокує до тих пір, поки один із його випадків не запуститься, потім він виконує цей випадок.
	Він випадково вибирає одне, якщо готові кілька.

	*/

	//output1 := make(chan string)
	//output2 := make(chan string)
	//go server1(output1)
	//go server2(output2)
	//
	//select {
	//case s1 := <-output1:
	//	fmt.Println(s1)
	//case s2 := <-output2:
	//	fmt.Println(s2)
	//}

	// 6/11
	// дефолтний кейс виконується якщо ніякий інший ще не готовий

	//tick := time.Tick(100 * time.Millisecond)
	//boom := time.After(500 * time.Millisecond)
	//for {
	//	select {
	//	case <-tick:
	//		fmt.Println("tick.")
	//	case <-boom:
	//		fmt.Println("BOOM!")
	//		return
	//	default:
	//		fmt.Println("    .")
	//		time.Sleep(50 * time.Millisecond)
	//	}
	//}

	//ch := make(chan bool) // канал
	//for i := 1; i < 5; i++ {
	//	go work(i, ch)
	//}
	//// очікуємо завершення всіх горутин
	//for i := 1; i < 5; i++ {
	//	<-ch
	//}
	//fmt.Println("The End")

	//ch := make(chan bool) // канал
	//var mutex sync.Mutex  // визначаємо м'ютекс
	//for i := 1; i < 5; i++ {
	//	go work(i, ch, &mutex)
	//}
	//
	//for i := 1; i < 5; i++ {
	//	<-ch
	//}
	//
	//fmt.Println("The End")
}

// 19/26

//type MyError struct {
//	When time.Time
//	What string
//}
//
//func (e *MyError) Error() string {
//	return fmt.Sprintf("at %v, %s", e.When, e.What)
//}
//
//func run() error {
//	return &MyError{
//		time.Now(),
//		"it didn't work",
//	}
//}

// 1/11

//func say(s string) {
//	for i := 0; i < 5; i++ {
//		time.Sleep(100 * time.Millisecond)
//		fmt.Println(s)
//	}
//}

// 2/11

//func sum(s []int, c chan int) {
//	sum := 0
//	for _, v := range s {
//		sum += v
//	}
//	c <- sum // send sum to c
//}

// 4/11

//func fibonacci(n int, c chan int) {
//	x, y := 0, 1
//	for i := 0; i < n; i++ {
//		c <- x
//		x, y = y, x+y
//	}
//	close(c)
//}

// 5/11

//func server1(ch chan string) {
//	time.Sleep(6 * time.Second)
//	ch <- "from server1"
//}
//
//func server2(ch chan string) {
//	time.Sleep(3 * time.Second)
//	ch <- "from server2"
//}

// 9/11

//var counter int = 0

//func work(number int, ch chan bool) {
//	counter = 0
//	for k := 1; k <= 5; k++ {
//		counter++
//		fmt.Println("Goroutine", number, "-", counter)
//	}
//	ch <- true
//}

//func work(number int, ch chan bool, mutex *sync.Mutex) {
//	mutex.Lock() // блокуємо доступ до змінної counter
//	counter = 0
//	for k := 1; k <= 5; k++ {
//		counter++
//		fmt.Println("Goroutine", number, "-", counter)
//	}
//	mutex.Unlock() // деблокуємо доступ
//	ch <- true
//}
