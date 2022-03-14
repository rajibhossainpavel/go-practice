package main

import "fmt"

func main() {
	var numOfWorkers int
	// How many workers do you want?
	fmt.Printf("Enter the number of workers to spawn : ")
	fmt.Scanf("%d\n", &numOfWorkers)

	c := make(chan int)
	outChanArr := make([]chan map[int]int, numOfWorkers)
	done := make(chan bool)
	collateChan := make(chan map[int]int)

	// generate some values
	go genFunc(c)

	for s := 0; s < numOfWorkers; s++ {
		// calculate the factorials
		outChanArr[s] = fact(c, done, s)
	}

	// collate results from all channels
	collateFunc(done, collateChan, outChanArr...)

	go func() {
		for s := 0; s < numOfWorkers; s++ {
			<-done // syncronise the goroutines
		}
		close(collateChan)
	}()

	for i := range collateChan {
		fmt.Println(i) // Print the result
	}
}

func genFunc(c chan<- int) {
	for i := 0; i < 100; i++ {
		for j := 0; j < 20; j++ {
			c <- j
		}
	}
	close(c)
}

func fact(cin <-chan int, done chan bool, routineName int) chan map[int]int {
	cout := make(chan map[int]int)
	go func() {
		counter := 0
		for i := range cin {
			out := 1
			for j := i; j > 1; j-- {
				out *= j
			}
			cout <- map[int]int{i: out}
			counter++
		}
		fmt.Println("Goroutine fact ", routineName, "processed", counter, "items")
		close(cout)
		fmt.Println("Goroutine fact", routineName, "is finished")
	}()
	return cout
}

func collateFunc(done chan bool, collateChan chan map[int]int, c ...chan map[int]int) {
	for idx, ci := range c {
		go func(ci chan map[int]int, idx int) {
			counter := 0
			for i := range ci {
				collateChan <- i
				counter++
			}
			fmt.Println("Goroutine consume ", idx, "consumed", counter, "items")
			done <- true
		}(ci, idx)
	}
}
