/*
We need to be able to quickly detect bad words in the messages our system sends.

Complete the indexOfFirstBadWord function. If it finds any bad words in
the message it should return the index of the first bad word in the msg slice.
This will help us filter out naughty words from our messaging system.
If no bad words are found, return -1 instead.

Use the range keyword.
*/

package main

func indexOfFirstBadWord(msg []string, badWords []string) int {
    for i, w := range msg {

        for _, bw := range badWords {

            if w == bw {
                return i
            }
        }
    }
	return -1
}
