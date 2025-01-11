/*
Complete the maxMessages function. Given a cost threshold,
it should calculate the maximum number of messages that can be sent.

Each message costs 100 pennies, plus an additional fee.
The fee structure is:

    1st message: 100 + 0
    2nd message: 100 + 1
    3rd message: 100 + 2
    4th message: 100 + 3

Browser Freeze

If you lock up your browser by creating an infinite loop that isn't breaking,
just click the cancel button.

*/

package main

func maxMessages(thresh int) int {
    msgCost := 100
    currCost := 0
    if thresh == msgCost {
        return 1
    }

    for i:=0;;i++ {
        currCost += msgCost + i
        if thresh < currCost {
            return i
        }
    }
}
