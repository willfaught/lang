package goo

var _ ChanReceive = ChanReceiveBool(nil)

var _ Pointer = (*ChanReceiveBool)(nil)

// ChanReceiveBool is a receive channel of bool.
type ChanReceiveBool chan bool

// Cap implements ChanReceive.
func (c ChanReceiveBool) Cap() int {
	return cap(c)
}

// Dereference implements ChanReceive.
func (c *ChanReceiveBool) Dereference() Value {
	return *c
}

// Len implements ChanReceive.
func (c ChanReceiveBool) Len() int {
	return len(c)
}

// Make implements ChanReceive.
func (c ChanReceiveBool) Make(cap int) Chan {
	return make(ChanBool, cap)
}

// Receive implements ChanReceive.
func (c ChanReceiveBool) Receive() interface{} {
	return <-c
}

// ReceiveCheck implements ChanReceive.
func (c ChanReceiveBool) ReceiveCheck() (interface{}, bool) {
	var v, ok = <-c

	return v, ok
}

// Reference implements ChanReceive.
func (c ChanReceiveBool) Reference() Pointer {
	return &c
}
