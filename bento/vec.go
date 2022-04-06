package bento

import (
	"image"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	// Deg90 is a 90 degree turn in radians.
	Deg90 = math.Pi / 2

	// Deg180 is a 180 degree turn in radians.
	Deg180 = math.Pi

	// Deg270 is a 270 degree turn in radians.
	Deg270 = 1.5 * math.Pi
)

var (
	buffer, empty *ebiten.Image
)

func init() {
	// strangely, vector grahpics require a buffer of 3x3 to render correctly.
	buffer = ebiten.NewImage(3, 3)
	buffer.Fill(color.White)

	empty = buffer.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)
}

// Vec is a wrapper around ebiten's vector path for drawing operations.
type Vec struct {
	Path vector.Path
}

// Move moves the vec's currnet position to a new position.
// This doesn't draw anything.
func (vec *Vec) Move(to image.Point) {
	vec.Path.MoveTo(float32(to.X), float32(to.Y))
}

// Line draws a line from the current position to another position.
func (vec *Vec) Line(to image.Point) {
	vec.Path.LineTo(float32(to.X), float32(to.Y))
}

// Rect draws a rectangle with bounds.
func (vec *Vec) Rect(bounds image.Rectangle) {
	d := image.Pt(0, bounds.Dy())

	p1 := bounds.Min
	p2 := bounds.Min.Add(d)
	p3 := bounds.Max
	p4 := bounds.Max.Sub(d)

	vec.Move(p1)
	vec.Line(p2)
	vec.Line(p3)
	vec.Line(p4)
}

// Arc draws a circular arc with a center and radius.
// from and to are angles in radians.
func (vec *Vec) Arc(center image.Point, radius int, from, to float32) {
	cx := float32(center.X)
	cy := float32(center.Y)

	vec.Path.Arc(cx, cy, float32(radius), from, to, vector.CounterClockwise)
}

// Circle draws a circle with a center and radius.
func (vec *Vec) Circle(center image.Point, radius int) {
	vec.Arc(center, radius, 0, 2*math.Pi)
}

func (vec *Vec) draw(clr color.Color) ([]ebiten.Vertex, []uint16) {
	r, g, b, a := clr.RGBA()
	nr := float32(r) / 0xFF
	ng := float32(g) / 0xFF
	nb := float32(b) / 0xFF
	na := float32(a) / 0xFF

	vs, is := vec.Path.AppendVerticesAndIndicesForFilling(nil, nil)
	for i := range vs {
		v := &vs[i]
		v.SrcX = 1
		v.SrcY = 1
		v.ColorR = nr
		v.ColorG = ng
		v.ColorB = nb
		v.ColorA = na
	}

	return vs, is
}

// Draw renders this vector path with color to an image.
func (vec *Vec) Draw(
	clr color.Color,
	img *ebiten.Image,
	o *ebiten.DrawTrianglesOptions,
) {
	vs, is := vec.draw(clr)

	img.DrawTriangles(vs, is, empty, o)
}

// DrawShader renders this vector path with a shader to an image.
func (vec *Vec) DrawShader(
	clr color.Color,
	img *ebiten.Image,
	shader *ebiten.Shader,
	o *ebiten.DrawTrianglesShaderOptions,
) {
	vs, is := vec.draw(clr)

	img.DrawTrianglesShader(vs, is, shader, o)
}
