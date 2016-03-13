package chans

type ChanBool chan bool

func (c ChanBool) Cap() int {
	return cap(c)
}

func (c ChanBool) ChanReceive() ChanReceive {
	return c
}

func (c ChanBool) ChanSend() ChanSend {
	return c
}

func (c ChanBool) Close() {
	close(c)
}

func (c ChanBool) Len() int {
	return len(c)
}

func (c ChanBool) Make(cap int) Chan {
	return make(ChanBool, cap)
}

func (c ChanBool) Receive() interface{} {
	return <-c
}

func (c ChanBool) ReceiveCheck() (interface{}, bool) {
	var v, ok = <-c

	return v, ok
}

func (c ChanBool) Send(v interface{}) {
	c <- v.(bool)
}

type ChanInt chan int

func (c ChanInt) Cap() int {
	return cap(c)
}

func (c ChanInt) ChanReceive() ChanReceive {
	return c
}

func (c ChanInt) ChanSend() ChanSend {
	return c
}

func (c ChanInt) Close() {
	close(c)
}

func (c ChanInt) Len() int {
	return len(c)
}

func (c ChanInt) Make(cap int) Chan {
	return make(ChanInt, cap)
}

func (c ChanInt) Receive() interface{} {
	return <-c
}

func (c ChanInt) ReceiveCheck() (interface{}, bool) {
	var v, ok = <-c

	return v, ok
}

func (c ChanInt) Send(v interface{}) {
	c <- v.(int)
}

type ChanInterface chan interface{}

func (c ChanInterface) Cap() int {
	return cap(c)
}

func (c ChanInterface) ChanReceive() ChanReceive {
	return c
}

func (c ChanInterface) ChanSend() ChanSend {
	return c
}

func (c ChanInterface) Close() {
	close(c)
}

func (c ChanInterface) Len() int {
	return len(c)
}

func (c ChanInterface) Make(cap int) Chan {
	return make(ChanInterface, cap)
}

func (c ChanInterface) Receive() interface{} {
	return <-c
}

func (c ChanInterface) ReceiveCheck() (interface{}, bool) {
	var v, ok = <-c

	return v, ok
}

func (c ChanInterface) Send(v interface{}) {
	c <- v
}

type ChanRune chan rune

func (c ChanRune) Cap() int {
	return cap(c)
}

func (c ChanRune) ChanReceive() ChanReceive {
	return c
}

func (c ChanRune) ChanSend() ChanSend {
	return c
}

func (c ChanRune) Close() {
	close(c)
}

func (c ChanRune) Len() int {
	return len(c)
}

func (c ChanRune) Make(cap int) Chan {
	return make(ChanRune, cap)
}

func (c ChanRune) Receive() interface{} {
	return <-c
}

func (c ChanRune) ReceiveCheck() (interface{}, bool) {
	var v, ok = <-c

	return v, ok
}

func (c ChanRune) Send(v interface{}) {
	c <- v.(rune)
}

type ChanString chan string

func (c ChanString) Cap() int {
	return cap(c)
}

func (c ChanString) ChanReceive() ChanReceive {
	return c
}

func (c ChanString) ChanSend() ChanSend {
	return c
}

func (c ChanString) Close() {
	close(c)
}

func (c ChanString) Len() int {
	return len(c)
}

func (c ChanString) Make(cap int) Chan {
	return make(ChanString, cap)
}

func (c ChanString) Receive() interface{} {
	return <-c
}

func (c ChanString) ReceiveCheck() (interface{}, bool) {
	var v, ok = <-c

	return v, ok
}

func (c ChanString) Send(v interface{}) {
	c <- v.(string)
}
