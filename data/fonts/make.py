import pathlib
import string

import fontforge
import svgwrite
from PIL import Image

SPRITE_PATH = pathlib.Path("raw")
SPRITE_PATH.mkdir(exist_ok=True)

SPRITE_SIZE = 5
SPRITE_OFFSET = SPRITE_SIZE + 1

SPRITE_NAMES = string.ascii_letters + string.digits + "!?.,;/<>"

SPRITESHEET = "teoran.png"

SVG_PIXEL_SIZE = 128
SVG_PIXEL_SIZE_TUP = tuple([f"{SVG_PIXEL_SIZE}px"] * 2)

SVG_SIZE_TUP = tuple([f"{SPRITE_SIZE * SVG_PIXEL_SIZE}px"] * 2)

FONT_FAMILY = "Teoran"
FONT_NAME = f"{FONT_FAMILY} Standard"


def parse_sprites(spritesheet):
    for y in range(0, spritesheet.height, SPRITE_OFFSET):
        for x in range(0, spritesheet.width, SPRITE_OFFSET):
            size = (x, y, x + SPRITE_SIZE, y + SPRITE_SIZE)
            sprite = spritesheet.crop(size)

            if sprite.getextrema()[3] != (0, 0):
                yield sprite


def sprite2svg(sprite, svg):
    for y in range(sprite.width):
        for x in range(sprite.height):
            pixel = sprite.getpixel((x, y))
            alpha = pixel[3]

            if alpha > 0:

                args = {
                    "insert": (f"{x * SVG_PIXEL_SIZE}px", f"{y * SVG_PIXEL_SIZE}px"),
                    "size": SVG_PIXEL_SIZE_TUP,
                    "fill": svgwrite.rgb(*pixel[:3]),
                }

                if alpha != 255:
                    args["opacity"] = alpha / 255

                svg.add(svg.rect(**args))


def main():
    spritesheet = Image.open(SPRITESHEET)

    # stage 1: convert spritesheet to individual SVGs
    for count, sprite in enumerate(parse_sprites(spritesheet)):
        svg = svgwrite.Drawing(
            filename=str(SPRITE_PATH / f"{count}.svg"), size=SVG_SIZE_TUP
        )
        sprite2svg(sprite, svg)
        svg.save()

    # stage 2: create a new font using fontforge and map the svgs
    font = fontforge.font()

    font.familyname = FONT_FAMILY
    font.fullname = FONT_NAME
    font.fontname = FONT_NAME

    font.encoding = "UnicodeFull"
    font.em = 1000

    for count, char in enumerate(SPRITE_NAMES):
        glyph = font.createMappedChar(char)
        glyph.importOutlines(str(SPRITE_PATH / f"{count}.svg"))

    # Add glyph for spacebar
    space = font.createMappedChar(" ")
    space.width = (SPRITE_SIZE * SVG_PIXEL_SIZE) // 2

    font.generate("teoran.ttf")
    font.save("teoran.sfd")


if __name__ == "__main__":
    main()
