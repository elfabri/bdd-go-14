/*

& references to something 
* dereferences that reference

& will be a direction in memory of something
* 

Complete the removeProfanity function.

It should use the strings.ReplaceAll function to replace all instances of the following words in the input message with asterisks.

    Word	Replacement
    fubb	****
    shiz	****
    witch	*****

It should update the value in the pointer and return nothing.

Do not alter the function signature.
*/

package main

import (
	"strings"
)

func removeProfanity(message *string) {
    badW := [3]string{"fubb", "shiz", "witch"}
    for _, w := range badW {
        censor := strings.Repeat("*", len(w))
        *message = strings.ReplaceAll(*message, w, censor)
    }
}
