package Bridge

import "fmt"

type Renderer interface {
	RenderCircle(radius float32)
}

type VectorRenderer struct {}

func (r *VectorRenderer) RenderCircle(radius float32) {
	fmt.Println("Drawing a circle of radius:", radius)
}

type RasterRenderer struct {
	Dpi int
}

func (r *RasterRenderer) RenderCircle(radius float32) {
	fmt.Println("Drawing pixes for a circle of radisu:", radius)
}

type Circle struct {
	renderer Renderer
	radius float32
}

func NewCircle(renderer Renderer, radius float32) *Circle {
	return &Circle{renderer: renderer, radius: radius}
}

func (c *Circle) Draw() {
	c.renderer.RenderCircle(c.radius)
}

func (c *Circle) Resize(factor float32) {
	c.radius *= factor
}

func main() {
	raster := RasterRenderer{}
	//vector := VectorRenderer{}
	circle := NewCircle(&raster, 5)
	circle.Draw()
	circle.Resize(2)
	circle.Draw()
}
