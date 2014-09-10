package animal

type Cat struct {
	content string
}

func (c *Cat) Eat(food string) {
	c.content = food
}

func (c *Cat) Shit() string {
	food := c.content
	c.content = ""
	return food + "shit"
}
