<template>
    <div>
    </div>
</template>


<script>
    import SorterFactory from "./algorithms/sorting";
    import {SortingAlgorithms} from "./algorithms/sorting";

    export default {
        name: "Algoman",
        props: {},
        components: {},
        data() {
            return {
                state: null,
                app: null,
                gameScene: null,
                sortingScene: null,
                size: 15,
                array: [],
                sorters: new SorterFactory(),
                sorter: null,
                currentStep: 0
            }
        },
        methods: {
            randomArray() {
                for (var i = 0; i < this.size; i++) {
                    this.array.push(i+1)
                }
                this.shuffleArray(this.array)
            },
            shuffleArray(array) {
                for (let i = array.length - 1; i > 0; i--) {
                    const j = Math.floor(Math.random() * (i + 1));
                    [array[i], array[j]] = [array[j], array[i]]; // eslint-disable-line no-param-reassign
                }
            },
            setupGameScene() {
                this.setupSortingScene()
            },
            setupSortingScene() {
                let vm = this
                let widthUnit = 5
                let heightUnit = 5
                let gap = 2

                let sortingScene = new PIXI.Container()
                this.sortingScene = sortingScene
                this.app.stage.addChild(sortingScene)
                this.randomArray()
                let array = this.array
                this.sorter = this.sorters.newSorter(SortingAlgorithms.bubble)

                for (var i = 0; i < array.length; i++) {
                    let rectangle = new PIXI.Graphics()
                    rectangle.beginFill(0x66CCFF)
                    rectangle.drawRect(0, 0, widthUnit, array[i] * heightUnit);
                    rectangle.endFill()
                    rectangle.x = (gap + widthUnit) * i
                    rectangle.y = heightUnit * (array.length - array[i])
                    rectangle.destination = {
                        x: rectangle.x,
                        y: rectangle.y
                    }
                    rectangle.moveTo = function (x, y) {
                        this.destination.x = x
                        this.destination.y = y
                    }
                    sortingScene.addChild(rectangle);
                }

                sortingScene.children.swap = function (a, b) {
                    let temp = this[a]
                    this[a] = this[b]
                    this[b] = temp

                    let cmdA = vm.makeMoveRectangleCommand(this[a], this[b].x, this[b].y)
                    cmdA.execute()

                    let cmdB = vm.makeMoveRectangleCommand(this[b], this[a].x, this[a].y)
                    cmdB.execute()
                }

                this.sorter.sort(array)
            },
            play(delta) {
                let rectangles = this.sortingScene.children
                let done = []
                for (var i = 0; i < rectangles.length; i++) {
                    if (rectangles[i].destination.x > rectangles[i].x) {
                        rectangles[i].x += 1
                    } else if (rectangles[i].destination.x < rectangles[i].x) {
                        rectangles[i].x -= 1
                    } else {
                        done.push(true)
                    }
                }

                if (done.length === rectangles.length && this.currentStep < this.sorter.commands.length) {
                    let step = this.sorter.commands[this.currentStep]

                    let cmdA = this.makeSwapCommand(rectangles, step.a, step.b)
                    cmdA.execute()
                    this.currentStep++
                }
            },
            load() {
                PIXI.loader.load(this.setup)
            },
            // This code will run when the loader has finished loading the image
            setup() {
                // this.setupAtlas()
                this.setupGameScene()

                this.state = this.play;
                this.app.ticker.add(delta => this.gameLoop(delta));
            },
            gameLoop(delta) {
                this.state(delta)
            },
            makeMoveRectangleCommand(rectangle, x, y) {
                var xBefore, yBefore
                return {
                    execute: function () {
                        xBefore = rectangle.x
                        yBefore = rectangle.y
                        rectangle.moveTo(x, y)
                    },
                    undo: function () {
                        rectangle.moveTo(xBefore, yBefore)
                    }
                }
            },
            makeSwapCommand(rectangles, a, b) {
                var aBefore, bBefore
                return {
                    execute: function () {
                        aBefore = a
                        bBefore = b
                        rectangles.swap(a, b)
                    },
                    undo: function () {
                        rectangles.swap(aBefore, bBefore)
                    }
                }
            }
        },
        mounted() {
            // Creating the Pixi Application
            this.app = new PIXI.Application({
                width: 256,         // default: 800
                height: 256,        // default: 600
                antialias: true,    // default: false
                transparent: false, // default: false
                resolution: 1       // default: 1
            });

            // Add the canvas that Pixi automatically created for you to the HTML document
            this.$el.appendChild(this.app.view)

            // Change the background
            this.app.renderer.backgroundColor = 0x012A36;

            // To make sure the canvas is resized to match the resolution, set autoResize to true
            this.app.renderer.autoResize = true;
            this.app.renderer.resize(512, 512);

            this.gameScene = new PIXI.Container();
            this.app.stage.addChild(this.gameScene);

            this.load()
        }
    }
</script>

<style lang="stylus">
</style>