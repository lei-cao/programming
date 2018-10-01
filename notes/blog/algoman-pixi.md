# Algoman with PIXI

After explore `Ebiten` for a while and made the first version Algoman. I kind feel that it's harder to 
achieve the goal of Algoman - visualising the algorithms in Go. So I need an easier and faster framework to work with to avoid rebuilding the wheels. 
Some Javascript Animation or Gaming frameworks are the obvious choices. That's why I am now exploring `pixi.js`, `vuepress` and `Go`.
I can use go to write the game logic, it's like the model layer. Then pass this data to `pixi.js` to enabling the animation and UI.

First step, loading `pixi.js` into `vuepress`

Found this binding [https://github.com/Granipouss/vue-pixi](https://github.com/Granipouss/vue-pixi), quite easy to set up.
<VuePixiDemo/>

## Learning PIXI

Still new to `pixi`, learnt the basics by following [Learning Pixi](https://github.com/kittykatattack/learningPixi) 
<LearningPixi/>

## Algoman first try with PIXI

<AlgomanDemo/>

## Resources

- [https://opengameart.org](https://opengameart.org)

