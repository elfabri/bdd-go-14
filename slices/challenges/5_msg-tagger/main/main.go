/*
                *** Message Tagger ***
Textio needs a way to tag messages based on specific criteria.

Complete the tagMessages function. It should take a slice
of sms messages, and a function (that takes a sms as input
and returns a slice of strings) as inputs. And it should return
a slice of sms messages.

It should loop through each message and set the tags to the result
of the passed in function.

Be sure to modify the messages of the original slice.

Ensure the tags field contains an initialized slice. No nil slices.

Complete the tagger function. It should take an sms message
and return a slice of strings.

For any message that contains "urgent" (regardless of casing)
in the content, the Urgent tag should be applied first.

For any message that contains "sale" (regardless of casing),
the Promo tag should be applied second.

Note: regardless of casing just means that even "uRGent" or "SALE"
should trigger the tag.

Example usage:

messages := []sms{
	{id: "001", content: "Urgent! Last chance to see!"},
	{id: "002", content: "Big sale on all items!"},
	// Additional messages...
}

taggedMessages := tagMessages(messages, tagger)
// `taggedMessages` will now have tags based on the content.
// 001 = [Urgent]
// 002 = [Promo]

*/

package main

import "strings"

type sms struct {
	id      string
	content string
	tags    []string
}

func tagMessages(messages []sms, tagger func(sms) []string) []sms {
    for i := range messages {
        if messages[i].content != "" {
            messages[i].tags = tagger(messages[i])
        }
    }
    return messages
}

func tagger(msg sms) []string {
	tags := []string{}
    words := strings.Split(msg.content, " ")
    isPromo := false
    isUrgent := false
    for _, w := range words {
        w = strings.ToLower(w)
        // avoid short words
        if (len(w) < 4) {
            continue
        }

        // we cut the word to avoid comas, etc
        if (len(w) > 6) {
            w = w[0:6]
        }
        if w == "urgent" {
            isUrgent = true
            continue
        }

        if (len(w) > 4) {
            w = w[0:4]
        }
        if w == "sale" {
            isPromo = true
        }
    }
    // Urgent tag has to be first
    if isUrgent {
        tags = append(tags, "Urgent")
    }
    if isPromo {
        tags = append(tags, "Promo")
    }
    return tags
}
