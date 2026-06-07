package main

func counter() (func(), int) {
	count := 0

	return func() {
		count += 1
	}, count
}

func main() {
	increment, _ := counter()
	increment()
	increment()
	increment()

}
