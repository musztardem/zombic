package components

type Collider struct {
	Position *Position
	Width    float64
	Height   float64
}

func (c *Collider) CollidesFromRightWith(c2 *Collider) bool {
	cLeft := c.Position.X
	cRight := c.Position.X + c.Width
	c2Left := c2.Position.X

	return cLeft < c2Left && cRight >= c2Left
}

func (c *Collider) CollidesFromLeftWith(c2 *Collider) bool {
	return c2.CollidesFromRightWith(c)
}

func (c *Collider) CollidesFromDownWith(c2 *Collider) bool {
	cTop := c.Position.Y
	cBottom := c.Position.Y + c.Height
	c2Bottom := c2.Position.Y + c2.Height

	return cBottom > c2Bottom && cTop <= c2Bottom
}

func (c *Collider) CollidesFromTopWith(c2 *Collider) bool {
	return c2.CollidesFromDownWith(c)
}

func (c *Collider) CollidesWith(c2 *Collider) bool {
	return c.overlapsX(c2) && c.overlapsY(c2)
}

func (c *Collider) overlapsX(c2 *Collider) bool {
	cLeft := c.Position.X
	cRight := c.Position.X + c.Width
	c2Left := c2.Position.X
	c2Right := c2.Position.X + c2.Width
	return cLeft < c2Right && cRight > c2Left
}

func (c *Collider) overlapsY(c2 *Collider) bool {
	cTop := c.Position.Y
	cBottom := c.Position.Y + c.Height
	c2Top := c2.Position.Y
	c2Bottom := c2.Position.Y + c2.Height
	return cTop < c2Bottom && cBottom > c2Top
}
