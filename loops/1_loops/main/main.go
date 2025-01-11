/*
At Textio we have a dynamic formula for determining how
much a batch of bulk messages costs to send. Complete 
the bulkSend() function.

It should return the total cost (as a float64) to send a 
batch of numMessages messages. Each message costs 1.0, 
plus an additional fee. The fee structure is:

    1st message: 1.0 + 0.00
    2nd message: 1.0 + 0.01
    3rd message: 1.0 + 0.02
    4th message: 1.0 + 0.03
    ...

Use a loop to calculate the total cost and return it.
*/

package main

func bulkSend(numMessages int) float64 {
    totalCost := 0.0
    for i:=0; i<numMessages; i++ {
        totalCost += 1.0 + float64(i)/100
    }
    return totalCost
}
