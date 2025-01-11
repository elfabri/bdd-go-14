package impl

// If an interface exists and a type has the proper methods defined,
// then the type automatically fulfills that interface.

// A type implements an interface by implementing its methods.

// Implicit interfaces decouple the definition of an interface from its
// implementation. You may add methods to a type and in the process
// be unknowingly implementing various interfaces, and that's okay.

// A type can implement any number of interfaces in Go. For example,
// the empty interface, interface{}, is always implemented by every type
// because it has no requirements.

type employee interface {
	getName() string
	getSalary() int
}

type contractor struct {
	name         string
	hourlyPay    int
	hoursPerYear int
}

func (c contractor) getName() string {
	return c.name
}

func (c contractor) getSalary() int {
	return c.hourlyPay * c.hoursPerYear
}

type fullTime struct {
	name   string
	salary int
}

func (ft fullTime) getSalary() int {
	return ft.salary
}

func (ft fullTime) getName() string {
	return ft.name
}
