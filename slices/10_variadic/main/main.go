/*
We need to sum up the costs of all individual messages so
we can send an end-of-month bill to our customers.

Complete the sum function to return the sum of all inputs.

Take note of how the variadic inputs and
the spread operator are used in the test suite.
*/

package main

func sum(nums ...int) int {
    totalCost := 0
    for i:=0; i<len(nums); i++ {
        totalCost += nums[i]
    }
    return totalCost
}
