type FooBar struct {
	n int
    fooCh chan struct{}
    barCh chan struct{}
}

func NewFooBar(n int) *FooBar {
	return &FooBar{
        n: n,
        fooCh: make(chan struct{}),
        barCh: make(chan struct{}),
        }
}

func (fb *FooBar) Foo(printFoo func()) {
	for i := 0; i < fb.n; i++ {

        if i>0 {
            <-fb.barCh
        }

		// printFoo() outputs "foo". Do not change or remove this line.
        printFoo()

        fb.fooCh <- struct{}{}
	}
}

func (fb *FooBar) Bar(printBar func()) {
	for i := 0; i < fb.n; i++ {

        <-fb.fooCh

		// printBar() outputs "bar". Do not change or remove this line.
        printBar()

        if i < fb.n-1 {
            fb.barCh <- struct{}{}
        }
	}
}