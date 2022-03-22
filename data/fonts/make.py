import pathlib

import fontforge
import svgwrite
from PIL import Image

SPRITE_PATH = pathlib.Path("./raw")
SVG_PATH = pathlib.Path("./svg")
SVG_PATH.mkdir(exist_ok=True)

SVG_PIXEL_SIZE = 128
SVG_PIXEL_SIZE_TUP = tuple([f"{SVG_PIXEL_SIZE}px"] * 2)

FONT_FAMILY = "Teoran"
FONT_NAME = f"{FONT_FAMILY} Standard"


def svg_size(sprite):
    return (str(sprite.width * SVG_PIXEL_SIZE), str(sprite.height * SVG_PIXEL_SIZE))


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
    # stage 1: convert spritesheet to individual SVGs
    for file in SPRITE_PATH.iterdir():
        if file.name.endswith(".png"):
            sprite = Image.open(file)
            svg = svgwrite.Drawing(
                filename=str(SVG_PATH / f"{file.stem}.svg"), size=svg_size(sprite)
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

    for file in SVG_PATH.iterdir():
        glyph = font.createChar(int(file.stem, 16))
        glyph.importOutlines(str(file))

    # Add glyph for spacebar
    space = font.createMappedChar(" ")
    space.width = (5 * SVG_PIXEL_SIZE) // 2

    font.generate("teoran.ttf")
    font.save("teoran.sfd")


if __name__ == "__main__":
    main()
