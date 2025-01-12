/*
It's important to keep up with privacy regulations and to respect
our user's data. We need a function that will delete user records.

Complete the deleteIfNecessary function.

Check the scheduledForDeletion bool to determine if they
are scheduled for deletion or not.

If the user doesn't exist in the map, return the error not found.
If they exist but aren't scheduled for deletion, return deleted
as false with no errors.
If they exist and are scheduled for deletion, return deleted as true
with no errors and delete their record from the map.

Note on Passing Maps
Like slices, maps are also passed by reference into functions.
This means that when a map is passed into a function we write,
we can make changes to the original, we don't have a copy.
*/

package main

import "errors"

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
    user, ok := users[name]
    if (!ok) {
        return false, errors.New("not found")
    }
    if (user.scheduledForDeletion) {
        delete(users, name)
        return true, nil
    }
    return false, nil
}

type user struct {
	name                 string
	number               int
	scheduledForDeletion bool
}
