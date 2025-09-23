package main

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)

	// Первые два чтения OK
	println(<-ch) // 1
	println(<-ch) // 2

	// Дальше - нулевые значения, т.к. канал закрыт и пуст
	println(<-ch) // 0
	println(<-ch) // 0
	println(<-ch) // 0

	// И range тоже будет бесконечно выдавать нули
	for n := range ch { // ВНИМАНИЕ: БЕСКОНЕЧНЫЙ ЦИКЛ!
		println(n) // Бесконечно печатает 0
	}
}
