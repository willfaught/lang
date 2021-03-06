package goo

var _ Chan = ChanUint16(nil)

var _ Pointer = (*ChanUint16)(nil)

// ChanUint16 is a channel of uint16.
type ChanUint16 chan uint16

// Cap implements Chan.
func (c ChanUint16) Cap() int {
	return cap(c)
}

// ChanReceive implements Chan.
func (c ChanUint16) ChanReceive() ChanReceive {
	return c
}

// ChanSend implements Chan.
func (c ChanUint16) ChanSend() ChanSend {
	return c
}

// Close implements Chan.
func (c ChanUint16) Close() {
	close(c)
}

// Dereference implements Chan.
func (c *ChanUint16) Dereference() Value {
	return *c
}

// Len implements Chan.
func (c ChanUint16) Len() int {
	return len(c)
}

// Make implements Chan.
func (c ChanUint16) Make(cap int) Chan {
	return make(ChanUint16, cap)
}

// Receive implements Chan.
func (c ChanUint16) Receive() interface{} {
	return <-c
}

// ReceiveCheck implements Chan.
func (c ChanUint16) ReceiveCheck() (interface{}, bool) {
	var v, ok = <-c

	return v, ok
}

// Reference implements Chan.
func (c ChanUint16) Reference() Pointer {
	return &c
}

// Send implements Chan.
func (c ChanUint16) Send(v interface{}) {
	c <- v.(uint16)
}
