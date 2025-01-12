/*
                *** Password Strength ***
As part of improving security, Textio wants to enforce a new password policy.
A valid password must meet the following criteria:

    * At least 5 characters long but no more than 12 characters.
    * Contains at least one uppercase letter.
    * Contains at least one digit.

A string is really just a read-only slice of bytes. This means that you
can use the same techniques you learned in the previous lesson to iterate
over the characters of a string.

Implement the isValidPassword function by looping through each character
in the password string. Make sure the password is long enough
and includes at least one uppercase letter and one digit.

*/

package main

import "unicode"

func isValidPassword(password string) bool {
    if len(password) < 5 || len(password) > 12 {
        return false
    }
    hasUpper := false
    hasDigit := false

    for _, c := range password {
        if unicode.IsUpper(c) {
            hasUpper = true
        }
        if unicode.IsDigit(c) {
            hasDigit = true
        }
    }
    return hasUpper && hasDigit
}
