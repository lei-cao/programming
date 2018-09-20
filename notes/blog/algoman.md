# Algoman

[Ebiten](https://github.com/hajimehoshi/ebiten) is a great tool for writing cross-platform 2D games in Golang.
It's very easy to converted my old canvas [algorithm visualizer](algorithm-visualization-refactoring.md) into it. Now I can 
build it for all major platforms.

[jsgo](https://jsgo.io/) is another awesome tool for using golang with GopherJS. It compiles the go (GopherJS) lib into js, serve and cache it. 
Then you can easily see the [result](https://jsgo.io/lei-cao/programming/code/algoman)

Loading it as a simple `iframe` in Vuepress:

<Ebiten id="heap"/>
(Quick Sort)

I gave it a final name too.
 
 `Algoman` - An interactive game for learning algorithms. 
 
 I will visualise all the major algorithms and made it as a game for learning algorithms.
 
> After writing in Vuepress for couple of weeks, I really loved it. So much fun to blogging and coding at the same time! Just powerful! 
 Waiting for the official blog theme support.
 
## Algoman with PIXI

After explore `Ebiten` for a while and made the first version Algoman. I kind feel that it's harder to 
achieve the goal of Algoman - visualising the algorithms in Go. So I need an easier and faster framework to work with to avoid rebuilding the wheels. 
Some Javascript Animation or Gaming frameworks are the obvious choices. That's why I am now exploring `pixi.js`, `vuepress` and `Go`.
I can use go to write the game logic, it's like the model layer. Then pass this data to `pixi.js` to enabling the animation, aka the view layer. 
`vue` then could be the controller. 

First step, loading `pixi.js` into `vuepress`

Found this binding [https://github.com/Granipouss/vue-pixi](https://github.com/Granipouss/vue-pixi), quite easy to set up.
<PIXI/>
