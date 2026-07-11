type ZeroEvenOdd struct {
	n        int
    zeroChan chan bool
    evenChan chan bool
    oddChan chan bool
}

func NewZeroEvenOdd(n int) *ZeroEvenOdd {
	zeo := &ZeroEvenOdd{
		n:        n,
        zeroChan: make(chan bool),
        evenChan: make(chan bool),
        oddChan: make(chan bool),
	}

    go func() {
        zeo.zeroChan <- true
    }()

	return zeo
}

func (z *ZeroEvenOdd) Zero(printNumber func(int)) {
    for i := 1; i <= z.n; i++{
        <-z.zeroChan
        printNumber(0)

        if i%2 != 0 {
            z.oddChan <- true
        } else {
            z.evenChan <- true
        }
    }
    <-z.zeroChan
}

func (z *ZeroEvenOdd) Even(printNumber func(int)) {
    for i := 2 ; i <= z.n ; i+=2{
        <-z.evenChan
        printNumber(i)
        z.zeroChan <- true
    }
}

func (z *ZeroEvenOdd) Odd(printNumber func(int)) {
    for i := 1 ; i <= z.n ; i+=2{
        <-z.oddChan
        printNumber(i)
        z.zeroChan <- true
    }
}