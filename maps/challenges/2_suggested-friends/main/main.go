/*
                    *** Suggested Friends ***
Textio is enhancing its social networking features to help users
identify their mutual friends.

Complete the findSuggestedFriends function. It takes a username string,
and a friendships map as inputs. The map keys are usernames,
and the values are slices of strings representing the direct friends
of that user.

Example friendship map:

friendships := map[string][]string{
    "Alice":   {"Bob", "Charlie"},
    "Bob":     {"Alice", "Charlie", "David"},
    "Charlie": {"Alice", "Bob", "David", "Eve"},
    "David":   {"Bob", "Charlie"},
    "Eve":     {"Charlie"},
}

The function should return a slice of strings containing
the user's suggested friends. A suggested friend is someone who
is not a direct friend of the user but is a direct friend of one
or more of the user's direct friends. Each suggested friend should
appear only once in the slice, even if they are found
through multiple direct friends.

If there are no suggested friends, return a nil slice.

For example using the friendships map above:

suggestedFriends := findSuggestedFriends("Alice", friendships)
// suggestedFriends = [David, Eve]
*/
package main

func findSuggestedFriends(username string, friendships map[string][]string) []string {

    // using map to filter
    // already friends (true) or already suggested friends (false)
    aF := map[string]bool{}
    // userFriends
    if uF, ok := friendships[username]; ok {
        for _, f := range uF {
            aF[f] = true
        }
    }

    // users with no friends should remain that way
    if len(aF) == 0 {
        return nil
    }

    // suggest friends
    sF := []string{}
    // username, friends of that un
    for un, f := range friendships {
        // skip current user
        if username == un {
            continue
        }

        // take undirect friends
        if isDirect, ok := aF[un]; ok && isDirect {
            // filter through friend's friends
            for _, ff := range f {
                // skip current user from friends list
                if ff == username {
                    continue
                }

                // filter through already suggested
                if _, ok := aF[ff]; !ok {
                    // friend's friend

                    // to be suggested
                    sF = append(sF, ff)

                    // and to be filtered
                    aF[ff] = false
                }
            }
        }
    }
    return sF
}
