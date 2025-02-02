/*
It's that time again, Mailio is hiring and we've been
assigned to do the interview. The Fibonacci sequence
is Mailio's interview problem of choice. We've been
tasked with building a small toy program we can use
in the interview.

Complete the concurrentFib function. It should:

    * Create a new channel of ints
    * Call fibonacci concurrently
    * Use a range loop to read from the channel and append the values to a slice
    * Return the slice
*/

package main

func concurrentFib(n int) []int {
    if n <= 0 {
        return []int{}
    }

    ch := make(chan int, n)
    result := []int{}
    fibonacci(n, ch)

    for v := range ch {
        result = append(result, v)
    }

    return result
}

// don't touch below this line

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}
