/*
As an Easter egg, we decided to reward our users with a free text message
if they send a prime number of text messages this year.

Complete the printPrimes function. It should print all of the prime numbers
up to and including max. It should skip any numbers that are not prime.

* Here's the pseudocode:

printPrimes(max):
  for n in range(2, max+1):
    if n is 2:
      n is prime, print it
    if n is even:
      n is not prime, skip to next n
    for i in range (3, sqrt(n), 2):
      if i can be multiplied into n:
        n is not prime, skip to next n
    n is prime, print it

* Breakdown
    -   We skip even numbers because they can't be prime

    -   We only check up to the square root of n. A factor higher than 
        the square root of n must multiply with a factor lower than 
        the square root of n, meaning it has no chance 
        of multiplying evenly into n.

    -   In your code, you can set the stop condition as i * i <= n

    -   We start checking at 2 because 1 is not prime

Note: This lesson is graded based on the output of your program,
so don't leave any debugging print statements in your code.

* Example Output
For the first test case, prime number up to 10:

Primes up to 10:
2
3
5
7
===============================================================

We only want you to print the numbers themselves, not the headings and delimiter.
*/

package main

import (
	"fmt"
)

func printPrimes(max int) {
    currIsPosPrime := true
    for i:=2; i<max; i++ {
        if i == 2 {
            fmt.Printf("%v\n",2)
        }
        if i % 2 == 0 {
            continue
        }
        currIsPosPrime = true
        for j:=3; j <= i; j+=2 {
            if i % j == 0 && i != j {
                currIsPosPrime = false
                break
            }
        }
        if currIsPosPrime {
            fmt.Printf("%v\n", i)
        }
    }
}

// don't edit below this line

func test(max int) {
	fmt.Printf("Primes up to %v:\n", max)
	printPrimes(max)
	fmt.Println("===============================================================")
}

func main() {
	test(10)
	test(20)
	test(30)
}
