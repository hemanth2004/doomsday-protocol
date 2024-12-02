package util

type Counter struct{ value int }

func (c *Counter) ValueAdd() int {
	v := c.value
	c.value += 1
	return v
}

func (c *Counter) Value() int {
	return c.value
}
