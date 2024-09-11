package counter

type Counter struct {
	count int
}

func (c *Counter) Increment() {
	c.count += 1
}

func (c *Counter) Decrement() {
	if c.count > 0 {
		c.count -= 1
	}
}

func (c *Counter) Reset() {
	c.count = 0
}

func (c *Counter) GetCount() int {
	return c.count
}
