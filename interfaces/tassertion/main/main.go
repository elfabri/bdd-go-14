package main

// When working with interfaces in Go, every once-in-awhile
// you'll need access to the underlying type of an interface value.
// You can cast an interface to its underlying type using a
// type assertion.

func getExpenseReport(e expense) (string, float64) {
    c, ok := e.(email)
    if ok {
        return c.toAddress, c.cost()
    }
    d, ok := e.(sms)
    if ok {
        return d.toPhoneNumber, d.cost()
    }
    return "", 0.0
}

// don't touch below this line

type expense interface {
	cost() float64
}

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type invalid struct{}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

func (i invalid) cost() float64 {
	return 0.0
}
