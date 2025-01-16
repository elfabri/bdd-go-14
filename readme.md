# Boot dot Dev - Course 14

Resolving the Go lang course from boot.dev

up to chapter 4 was done on the platform's web

running test locally with:

```
go test ./*_test.go ./main.go -v

```

## progress:

* up to chapter 4 - Structs:        completed and tested
* chapter 5 -       Interfaces:     completed and tested
* chapter 6 -       Errors:         completed and tested
* chapter 7 -       Loops:          completed and tested
* chapter 8 -       Slices:         completed and tested
* chapter 9 -       Map:            completed and tested
* chapter 10 -      Pointers:       completed and tested
* chapter 11 -      Packages
                    and Modules:    completed and tested
* chapter 12 -      Channels:       completed and tested
* chapter 13 -      Mutexes:        completed and tested
* chapter 14 -      Generics:       completed and tested
* chapter 15 -      Enums:          completed and tested
* chapter 16 -      Quiz:           a

## [Go proverbs](https://go-proverbs.github.io/) from Rob Pike

* Don't communicate by sharing memory, share memory by communicating.

* Concurrency is not parallelism.

* Channels orchestrate; mutexes serialize.

* The bigger the interface, the weaker the abstraction.

* Make the zero value useful.

* interface{} says nothing.

* Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.

* A little copying is better than a little dependency.

* Syscall must always be guarded with build tags.

* Cgo must always be guarded with build tags.

* Cgo is not Go.

* With the unsafe package there are no guarantees.

* Clear is better than clever.

* Reflection is never clear.

* Errors are values.

* Don't just check errors, handle them gracefully.

* Design the architecture, name the components, document the details.

* Documentation is for users.

* Don't panic.
