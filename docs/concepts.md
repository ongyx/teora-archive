# Concepts

## Component

The foundation of bento's model is the component.

A component encapsulates some state and logic that is updated every tick by calling it's `Update` method, similar to ebiten's `Game.Update`.

If any subcomponents are being used, their `Update` method **must** be called before using their state.

## Sprite

A sprite is a component that can render to an image.

## Entity

An entity contains a sprite to draw to the screen, as well as it's rendering state such as the options used to draw the sprite to the screen.

## Animation

An animation is a component that draws on top of a sprite/scene for a finite number of ticks.

## Transition

A transition is an animation that draws over a sprite/scene when it's entering or exiting the stage.

Transitions also control the visibility of a sprite.

## Scene

A scene is a special kind of entity: it renders other entities to the screen, and optionally draws other non-entity images over them.
This can be considered analogous to a game level.

bento renders the slice of entities returned from the `Entities` method of a scene in order; from first to last.
This allows you to control how entities overlap in a scene.

## Stage

A stage draws and changes scenes on screen.
It implements the `ebiten.Game` interface, and therefore an instance can be passed directly to `ebiten.RunGame`.
