/*
At Mailio we store all the emails for a campaign in memory as a slice.
We store payments for a single user in the same way.

Complete the getLast() function. It should be a generic function
that returns the last element from a slice,
no matter the types stored in the slice. If the slice is empty,
it should return the zero value of the type.
*/

package main

func getLast[T any](s []T) T {
    var last T
    if l := len(s); l >=1 {
        last = s[l-1]
    }
    return last
}
