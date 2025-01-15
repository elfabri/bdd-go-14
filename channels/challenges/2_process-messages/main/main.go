/*
                        *** Process Messages ***
Textio needs to speed up message processing
using concurrency.

Assignment:
Implement the processMessages function, it takes
a slice of messages. It should process each message
concurrently within a goroutine. Use the process function
for this to simulate the benefit of using goroutines.
Use a channel to ensure that all messages are processed
and collected correctly, then return the slice
of processed messages.

Example output:

messages := []string{
    "Here are some messages", 
    "Here is another",
    "and another"}

processedMessages := processMessages(messages)

fmt.Println(processedMessages)

// prints [
"Here are some messages-processed",
"Here is another-processed",
"and another-processed"
]

notes:
noooooo sleeeeep timeeeee aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
no sync, just clean
9 (s) with no goroutines
4 (s) with them, ez
*/

package main

import "time"

func processMessages(messages []string) []string {
    var pM []string         // to sort and collect the processed messages
    var ch chan string      // to recive processed messages
    var cch chan []string   // to return all at once

    if l := len(messages); l >= 1 {
        ch = make(chan string, l)
        cch = make(chan []string, 1)
        pM = make([]string, l)
    } else {
        return nil
    }

    i := 0
    for _, m := range messages {
        go func() {
            ch <- process(m)
        }()
    }
    
    go func() {
        for s := range ch {
            pM[i] =  s
            i++

            if i == len(pM) {
                cch <- pM
            }
        }
    }()

    return <- cch
}

// don't touch below this line

func process(message string) string {
	time.Sleep(1 * time.Second)
	return message + "-processed"
}
