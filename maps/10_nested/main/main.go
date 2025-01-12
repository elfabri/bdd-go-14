/*
Maps can contain maps, creating a nested structure. For example:

map[string]map[string]int
map[rune]map[string]int
map[int]map[string]map[string]int

Because Textio is a glorified customer database, we have a lot
of internal logic for sorting and dealing with customer names.

Complete the getNameCounts function. It takes a slice of strings names
and returns a nested map. The parent map's keys are all the unique
first characters (see runes) of the names, the nested maps keys
are all the names themselves, and the value is the count of each name.

For example:

billy
billy
bob
joe
Copy icon
Creates the following nested map:

b: {
    billy: 2,
    bob: 1
},
j: {
    joe: 1
}
*/

package main

func getNameCounts(names []string) map[rune]map[string]int {
	// Your code here
    m := make(map[rune]map[string]int)
    for _, n := range names {
        // turn name (string) to slice of runes and
        // take first rune, I hate emojis
        fC := []rune(n)[0]
        if mm, ok := m[fC]; ok {
            // first char exists
            if _, ok := m[fC][n]; ok {
                // name exists, plus plus yk
                m[fC][n]++
                continue
            }
            // name does not exists, initiate it 
            mm[n] = 1
            continue
        }
        // first char does not exist, initiate it 
        m[fC] = map[string]int {
            n: 1,
        } 
    }
    return m
}
