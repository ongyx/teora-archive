package bento

import (
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type tileCache map[int]*ebiten.Image

// Tileset is a image with tiles that can be tiled into a single sprite.
type Tileset struct {
	image *ebiten.Image
	tsize int

	cache tileCache
	size  image.Point
}

// NewTileset creates a new tileset with a image and tile size.
// If cache is true, calls to the Tile method will lazily cache tile images for faster loading.
func NewTileset(img *ebiten.Image, tsize int, cache bool) *Tileset {
	w, h := img.Size()

	if tsize <= 0 || tsize > w || tsize > h {
		panic("tileset: tilesize must be within bounds of tileset")
	}

	var tc tileCache
	if cache {
		tc = make(tileCache)
	}

	return &Tileset{img, tsize, tc, image.Pt(w/tsize, h/tsize)}
}

// Size returns the tileset size as (columns, rows).
func (t *Tileset) Size() image.Point {
	return t.size
}

// Tilesize returns the size of a tile in the tileset.
func (t *Tileset) Tilesize() int {
	return t.tsize
}

// Tile returns the tile at index as a subimage of the tileset,
// where index is the tile column multiplied by the tile row.
// This panics if the index is out of bounds.
func (t *Tileset) Tile(index int) *ebiten.Image {
	if index >= (t.size.X * t.size.Y) {
		panic(fmt.Sprintf("tileset: index %d out of bounds", index))
	}

	var tile *ebiten.Image

	if t.cache != nil {
		if ct, ok := t.cache[index]; ok {
			tile = ct
		} else {
			tile = t.tile(index)
			t.cache[index] = tile
		}
	} else {
		tile = t.tile(index)
	}

	return tile
}

// Render renders the tileset to an image, given a tilemap.
// The tilemap must have at least one row, and all rows must have equal length.
func (t *Tileset) Render(tilemap [][]int) *ebiten.Image {
	w := len(tilemap[0]) * t.tsize
	h := len(tilemap) * t.tsize
	img := ebiten.NewImage(w, h)

	for y, row := range tilemap {
		for x, index := range row {
			dx := x * t.tsize
			dy := y * t.tsize

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(dx), float64(dy))

			img.DrawImage(t.Tile(index), op)
		}
	}

	return img
}

func (t *Tileset) tile(index int) *ebiten.Image {
	x := (index % t.size.X) * t.tsize
	y := (index / t.size.X) * t.tsize

	return t.image.SubImage(image.Rect(x, y, x+t.tsize, y+t.tsize)).(*ebiten.Image)
}
