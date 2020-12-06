package factory

import "fmt"

// Since go isn't an oop language, the complete implementtaion of the factory pattern
// can't be implemented (inheritance and classes aren't used in go).

// The problem we try to fix here is : consider a company which business is logistics
// in the beginning, they work just with trucks. A year later, maritime companies want
// to be part of the business. Now developpers should change all the code to adapt to the
// fresher. Finally, developpers decided to implement factory pattern to solve the problem.

// Logistics is the type interface of all transports methods
type Logistics interface {
	Deliver()
	SetName(name string)
	SetEngine(engine int)
	Name() string
	Engine() int
}

// Transport holds data for transport methods
type Transport struct {
	name   string
	engine int
}

// Deliver print the method of delivering
func (t *Transport) Deliver() {
	fmt.Println("We make delivering with: ", t.name)
}

// Name return the name of the transport
func (t *Transport) Name() string {
	return t.name
}

// SetName set the name of the transport method
func (t *Transport) SetName(name string) {
	t.name = name
}

// Engine return the power of the engine
func (t *Transport) Engine() int {
	return t.engine
}

// SetEngine set new engine to transport
func (t *Transport) SetEngine(engine int) {
	t.engine = engine
}
