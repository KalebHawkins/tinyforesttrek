package game

type Camera struct {
	X, Y           float64
	ViewPortWidth  int
	ViewPortHeight int
	WorldWidth     int
	WorldHeight    int
}

func (c *Camera) Follow(targetX, targetY float64) {
	halfW := c.ViewPortWidth / 2
	halfH := c.ViewPortHeight / 2

	c.X = targetX - float64(halfW)
	c.Y = targetY - float64(halfH)

	if c.X < 0 {
		c.X = 0
	}
	if c.Y < 0 {
		c.Y = 0
	}

	maxX := float64(c.WorldWidth - c.ViewPortWidth)
	maxY := float64(c.WorldHeight - c.ViewPortHeight)
	if c.X > maxX {
		c.X = maxX
	}
	if c.Y > maxY {
		c.Y = maxY
	}
}

func NewCamera(viewportWidth, viewportHeight, worldWidth, worldHeight int) *Camera {
	c := &Camera{
		X:              0,
		Y:              0,
		ViewPortWidth:  viewportWidth,
		ViewPortHeight: viewportHeight,
		WorldWidth:     worldWidth,
		WorldHeight:    worldHeight,
	}

	return c
}
