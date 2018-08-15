# Algorithm Visualization

I made a very simple canvas app for visualising some basic sorting algorithms by [GopherJS](https://github.com/gopherjs/gopherjs)

<<< @/code/v1/sort/bubble.go

Then I just loaded it in a Vue Component. Hacky.. But works now.

```html
<ACanvasV1 id="bubble"/>
<ACanvasV1 id="bubble_swapped"/>
```

## Demo

### Bubble sort

<ACanvasV1 id="bubble"/>


### Bubble sort swapped
<ACanvasV1 id="bubble_swapped"/>


### Selection sort

<<< @/code/v1/sort/selection.go

<ACanvasV1 id="selection"/>

## What's Next?

Improve the visualizing part to make the algorithm code clean and the animation better. Considering using [svg.js](https://github.com/svgdotjs/svg.js)