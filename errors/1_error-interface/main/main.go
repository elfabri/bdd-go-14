package main

import (
	"fmt"
)

/*
We offer a product that allows businesses to send pairs of messages
to couples. It is mostly used by flower shops and movie theaters.

Complete the sendSMSToCouple function. It should send 2 messages,
first to the customer and then to the customer's spouse.

Use sendSMS() to send the msgToCustomer. If an error is encountered,
return 0 and the error.

Do the same for the msgToSpouse

If both messages are sent successfully, return the total cost of the
messages added together.

When you return a non-nil error in Go, it's conventional to return the
"zero" values of all other return values.
*/

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (int, error) {

    customerCost, customerErr := sendSMS(msgToCustomer)
    if customerErr != nil {
        return 0, customerErr
    }

    spouseCost, spouseErr := sendSMS(msgToSpouse)
    if spouseErr != nil {
        return 0, spouseErr
    }

    return customerCost + spouseCost, nil
}

// don't edit below this line

func sendSMS(message string) (int, error) {
	const maxTextLen = 25
	const costPerChar = 2
	if len(message) > maxTextLen {
		return 0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * len(message), nil
}
