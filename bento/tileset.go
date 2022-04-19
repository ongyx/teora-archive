package bento

import (
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

// Tileset is a image with tiles that can be tiled into a single sprite.
type Tileset struct {
	image    *ebiten.Image
	tilesize int

	size image.Point
}

// NewTileset creates a new tileset with a tile size and image.
func NewTileset(img *ebiten.Image, tilesize int) *Tileset {
	w, h := img.Size()

	if tilesize <= 0 || tilesize > w || tilesize > h {
		panic("tileset: tilesize must be within bounds of tileset")
	}

	return &Tileset{img, tilesize, image.Pt(w/tilesize, h/tilesize)}
}

// Size returns the tileset size as (columns, rows).
func (t *Tileset) Size() image.Point {
	return t.size
}

// Tilesize returns the size of a tile in the tileset.
func (t *Tileset) Tilesize() int {
	return t.tilesize
}

// Tile returns the tile at index as a subimage of the tileset,
// where index is the tile column multiplied by the tile row.
// This panics if the index is out of bounds.
func (t *Tileset) Tile(index int) *ebiten.Image {
	if index >= (t.size.X * t.size.Y) {
		panic(fmt.Sprintf("tileset: index %d out of bounds", index))
	}

	x := (index % t.size.X) * t.tilesize
	y := (index / t.size.X) * t.tilesize

	return t.image.SubImage(image.Rect(x, y, x+t.tilesize, y+t.tilesize)).(*ebiten.Image)
}

// Render renders the tileset to an image, given a tilemap.
// The tilemap must have at least one row, and all rows must have equal length.
func (t *Tileset) Render(tilemap [][]int) *ebiten.Image {
	w := len(tilemap[0]) * t.tilesize
	h := len(tilemap) * t.tilesize
	img := ebiten.NewImage(w, h)

	for y, row := range tilemap {
		for x, index := range row {
			dx := x * t.tilesize
			dy := y * t.tilesize

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(dx), float64(dy))

			img.DrawImage(t.Tile(index), op)
		}
	}

	return img
}
