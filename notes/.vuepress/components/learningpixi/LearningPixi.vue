<template>
    <div class="learning-pixi">

    </div>
</template>

<script>
    //Aliases
    let Application = PIXI.Application,
        Container = PIXI.Container,
        loader = PIXI.loader,
        resources = PIXI.loader.resources,
        TextureCache = PIXI.utils.TextureCache,
        Sprite = PIXI.Sprite,
        Rectangle = PIXI.Rectangle,
        utils = PIXI.utils,
        TextStyle = PIXI.TextStyle,
        Text = PIXI.Text

    export default {
        name: "LearningPixi",
        props: {},
        components: {},
        data() {
            return {
                state: null,
                app: null,
                gameScene: null,
                gameOverScene: null,
                explorer: null,
                blobs: [],
                treasure: null,
                healthBar: null,
                door: null,
                dungeon: null,
                message: null
            }
        },
        methods: {
            initGameScene() {
                let dungeon, door, id
                //Create an alias for the texture atlas frame ids
                id = resources["/images/learningpixi/treasureHunter.json"].textures;

                //Dungeon
                this.dungeon = new Sprite(id["dungeon.png"]);
                this.gameScene.addChild(this.dungeon);

                //Door
                this.door = new Sprite(id["door.png"]);
                this.door.position.set(32, 0);
                this.gameScene.addChild(this.door);

                //Explorer
                this.explorer = new Sprite(id["explorer.png"]);
                this.explorer.x = 68;
                this.explorer.y = this.gameScene.height / 2 - this.explorer.height / 2;
                this.explorer.vx = 0;
                this.explorer.vy = 0;
                this.gameScene.addChild(this.explorer);
                this.setupExplorer()

                //Treasure
                this.treasure = new Sprite(id["treasure.png"]);
                this.treasure.x = this.gameScene.width - this.treasure.width - 48;
                this.treasure.y = this.gameScene.height / 2 - this.treasure.height / 2;
                this.gameScene.addChild(this.treasure);


                let numberOfBlobs = 6,
                    spacing = 48,
                    xOffset = 150,
                    speed = 2,
                    direction = 1;

                //An array to store all the blob monsters

                //Make as many blobs as there are `numberOfBlobs`
                for (let i = 0; i < numberOfBlobs; i++) {

                    //Make a blob
                    let blob = new Sprite(id["blob.png"]);

                    //Space each blob horizontally according to the `spacing` value.
                    //`xOffset` determines the point from the left of the screen
                    //at which the first blob should be added
                    let x = spacing * i + xOffset;

                    //Give the blob a random `y` position
                    let y = this.randomInt(0, this.app.stage.height - blob.height);

                    //Set the blob's position
                    blob.x = x;
                    blob.y = y;

                    //Set the blob's vertical velocity. `direction` will be either `1` or
                    //`-1`. `1` means the enemy will move down and `-1` means the blob will
                    //move up. Multiplying `direction` by `speed` determines the blob's
                    //vertical direction
                    blob.vy = speed * direction;

                    //Reverse the direction for the next blob
                    direction *= -1;

                    //Push the blob into the `blobs` array
                    this.blobs.push(blob);

                    //Add the blob to the `gameScene`
                    this.gameScene.addChild(blob);
                }

                //Create the health bar
                this.healthBar = new PIXI.DisplayObjectContainer();
                this.healthBar.position.set(this.app.stage.width - 170, 4)
                this.gameScene.addChild(this.healthBar);

                //Create the black background rectangle
                let innerBar = new PIXI.Graphics();
                innerBar.beginFill(0x000000);
                innerBar.drawRect(0, 0, 128, 8);
                innerBar.endFill();
                this.healthBar.addChild(innerBar);

                //Create the front red rectangle
                let outerBar = new PIXI.Graphics();
                outerBar.beginFill(0xFF3300);
                outerBar.drawRect(0, 0, 128, 8);
                outerBar.endFill();
                this.healthBar.addChild(outerBar);

                this.healthBar.outer = outerBar;

            },
            initGameOverScene() {
                let style = new TextStyle({
                    fontFamily: "Futura",
                    fontSize: 64,
                    fill: "white"
                });
                this.message = new Text("The End!", style);
                this.message.x = 120;
                this.message.y = this.app.stage.height / 2 - 32;
                this.gameOverScene.addChild(this.message);
            },
            play(delta) {
                let vm = this
                this.explorer.x += this.explorer.vx;
                this.explorer.y += this.explorer.vy;
                this.contain(this.explorer, {x: 28, y: 10, width: 488, height: 480});


                let explorerHit = false
                this.blobs.forEach(function (blob) {

                    //Move the blob
                    blob.y += blob.vy;

                    //Check the blob's screen boundaries
                    let blobHitsWall = vm.contain(blob, {x: 28, y: 10, width: 488, height: 480});

                    //If the blob hits the top or bottom of the stage, reverse
                    //its direction
                    if (blobHitsWall === "top" || blobHitsWall === "bottom") {
                        blob.vy *= -1;
                    }

                    //Test for a collision. If any of the enemies are touching
                    //the explorer, set `explorerHit` to `true`
                    if (vm.hitTestRectangle(vm.explorer, blob)) {
                        explorerHit = true;
                    }
                });

                if (explorerHit) {

                    //Make the explorer semi-transparent
                    this.explorer.alpha = 0.5;

                    //Reduce the width of the health bar's inner rectangle by 1 pixel
                    this.healthBar.outer.width -= 1;

                } else {

                    //Make the explorer fully opaque (non-transparent) if it hasn't been hit
                    this.explorer.alpha = 1;
                }

                if (this.hitTestRectangle(this.explorer, this.treasure)) {
                    this.treasure.x = this.explorer.x + 8;
                    this.treasure.y = this.explorer.y + 8;
                }

                if (this.hitTestRectangle(this.treasure, this.door)) {
                    this.state = this.end
                    this.message.text = "You won!"
                }

                if (this.healthBar.outer.width < 0) {
                    this.state = this.end
                    this.message.text = "You lost!"
                }

            },
            end(delta) {
                this.gameScene.visible = false;
                this.gameOverScene.visible = true;
            },
            contain(sprite, container) {

                let collision = undefined;

                //Left
                if (sprite.x < container.x) {
                    sprite.x = container.x;
                    collision = "left";
                }

                //Top
                if (sprite.y < container.y) {
                    sprite.y = container.y;
                    collision = "top";
                }

                //Right
                if (sprite.x + sprite.width > container.width) {
                    sprite.x = container.width - sprite.width;
                    collision = "right";
                }

                //Bottom
                if (sprite.y + sprite.height > container.height) {
                    sprite.y = container.height - sprite.height;
                    collision = "bottom";
                }

                //Return the `collision` value
                return collision;
            },
            hitTestRectangle(r1, r2) {

                //Define the variables we'll need to calculate
                let hit, combinedHalfWidths, combinedHalfHeights, vx, vy;

                //hit will determine whether there's a collision
                hit = false;

                //Find the center points of each sprite
                r1.centerX = r1.x + r1.width / 2;
                r1.centerY = r1.y + r1.height / 2;
                r2.centerX = r2.x + r2.width / 2;
                r2.centerY = r2.y + r2.height / 2;

                //Find the half-widths and half-heights of each sprite
                r1.halfWidth = r1.width / 2;
                r1.halfHeight = r1.height / 2;
                r2.halfWidth = r2.width / 2;
                r2.halfHeight = r2.height / 2;

                //Calculate the distance vector between the sprites
                vx = r1.centerX - r2.centerX;
                vy = r1.centerY - r2.centerY;

                //Figure out the combined half-widths and half-heights
                combinedHalfWidths = r1.halfWidth + r2.halfWidth;
                combinedHalfHeights = r1.halfHeight + r2.halfHeight;

                //Check for a collision on the x axis
                if (Math.abs(vx) < combinedHalfWidths) {

                    //A collision might be occurring. Check for a collision on the y axis
                    if (Math.abs(vy) < combinedHalfHeights) {

                        //There's definitely a collision happening
                        hit = true;
                    } else {

                        //There's no collision on the y axis
                        hit = false;
                    }
                } else {

                    //There's no collision on the x axis
                    hit = false;
                }

                //`hit` will be either `true` or `false`
                return hit;
            },
            // Loading images into the texture cache
            loader() {
                loader
                    .add([
                        "/images/learningpixi/cat.png",
                        "/images/learningpixi/blob.png"
                    ])
                    .add('tileset', '/images/learningpixi/tileset.png')
                    .add('/images/learningpixi/treasureHunter.json')
                    .load(this.setup)
            },
            // This code will run when the loader has finished loading the image
            setup() {
                // this.setupAtlas()
                this.initGameScene()
                this.initGameOverScene()

                this.state = this.play;
                this.app.ticker.add(delta => this.gameLoop(delta));
            },
            setupExplorer() {
                //Capture the keyboard arrow keys
                let left = this.keyboard(37),
                    up = this.keyboard(38),
                    right = this.keyboard(39),
                    down = this.keyboard(40);

                //Left arrow key `press` method
                left.press = () => {
                    //Change the cat's velocity when the key is pressed
                    this.explorer.vx = -5;
                    this.explorer.vy = 0;
                };

                //Left arrow key `release` method
                left.release = () => {
                    //If the left arrow has been released, and the right arrow isn't down,
                    //and the cat isn't moving vertically:
                    //Stop the cat
                    if (!right.isDown && this.explorer.vy === 0) {
                        this.explorer.vx = 0;
                    }
                };

                //Up
                up.press = () => {
                    this.explorer.vy = -5;
                    this.explorer.vx = 0;
                };
                up.release = () => {
                    if (!down.isDown && this.explorer.vx === 0) {
                        this.explorer.vy = 0;
                    }
                };

                //Right
                right.press = () => {
                    this.explorer.vx = 5;
                    this.explorer.vy = 0;
                };
                right.release = () => {
                    if (!left.isDown && this.explorer.vy === 0) {
                        this.explorer.vx = 0;
                    }
                };

                //Down
                down.press = () => {
                    this.explorer.vy = 5;
                    this.explorer.vx = 0;
                };
                down.release = () => {
                    if (!up.isDown && this.explorer.vx === 0) {
                        this.explorer.vy = 0;
                    }
                };
            },
            randomInt(min, max) {
                return Math.floor(Math.random() * (max - min + 1)) + min;
            },
            gameLoop(delta) {
                this.state(delta)
            },
            keyboard(keyCode) {
                let key = {};
                key.code = keyCode;
                key.isDown = false;
                key.isUp = true;
                key.press = undefined;
                key.release = undefined;
                //The `downHandler`
                key.downHandler = event => {
                    if (event.keyCode === key.code) {
                        if (key.isUp && key.press) key.press();
                        key.isDown = true;
                        key.isUp = false;
                    }
                    event.preventDefault();
                };

                //The `upHandler`
                key.upHandler = event => {
                    if (event.keyCode === key.code) {
                        if (key.isDown && key.release) key.release();
                        key.isDown = false;
                        key.isUp = true;
                    }
                    event.preventDefault();
                };

                //Attach event listeners
                window.addEventListener(
                    "keydown", key.downHandler.bind(key), false
                );
                window.addEventListener(
                    "keyup", key.upHandler.bind(key), false
                );
                return key;
            }
        },
        mounted() {
            // Creating the Pixi Application
            this.app = new Application({
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

            this.gameScene = new Container();
            this.app.stage.addChild(this.gameScene);

            this.gameOverScene = new Container();
            this.app.stage.addChild(this.gameOverScene);
            this.gameOverScene.visible = false;

            this.loader()
        }
    }
</script>

<style lang="stylus">
</style>