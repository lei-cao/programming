<template>
    <div>
        <ClientOnly>
            <pixi-renderer
                    @tick="update"
                    :backgroundColor="0x012A36"
                    :height="320"
                    v-bind:width="width"
            >
                <pixi-container
                        :x="240" :y="160"
                        :rotation="-t / 40"
                >
                    <pixi-sprite v-for="(sprite, key) in sprites" :key="key"
                                 src="/favicon.ico"
                                 :x="sprite.x" :y="sprite.y"
                                 :scaleX="sprite.scale" :scaleY="sprite.scale"
                                 :anchorX="0.5" :anchorY="0.5"
                                 :rotation="sprite.angle + t / 60"/>
                </pixi-container>
            </pixi-renderer>
        </ClientOnly>
    </div>
</template>

<script>
    import PixiRenderer from './pixi/Renderer'
    import PixiContainer from './pixi/Container'
    import PixiSprite from './pixi/Sprite'

    export default {
        name: "VuePixiDemo",
        props: {
            id: String
        },
        components: {
            PixiRenderer,
            PixiContainer,
            PixiSprite
        },
        data() {
            return {
                width: 300,
                t: 0,
                sprites: []
            }
        },
        methods: {
            addSprite() {
                this.sprites.push({
                    x: 480 * (0.5 - Math.random()),
                    y: 320 * (0.5 - Math.random()),
                    angle: 2 * Math.PI * Math.random(),
                    scale: 0.25 + 0.5 * Math.random(),
                })
            },
            update(dt) {
                this.t += dt
            }
        },
        mounted() {
            this.addSprite()
            this.addSprite()
            this.addSprite()
        }
    }
</script>
