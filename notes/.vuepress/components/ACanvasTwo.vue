<template>
    <div>

        <div :id="id" class="canvas">
        </div>

        <button @click="stop">Stop</button>
        <button @click="resume">Resume</button>
        <input v-model="size"/>
        <button @click="restart">Restart</button>
        <ClientOnly>
            <vue-slider
                ref="slider"
                v-model="speed"
                v-bind="options"
                @callback="updateSpeed"
            ></vue-slider>
        </ClientOnly>


    </div>
</template>

<script>
export default {
    name: "ACanvasTwo",
    props: {
        id: String
    },
    components: {
        'vueSlider': () => import('../../../node_modules/vue-slider-component')
    },
    data () {
        return {
            speed: 200,
            size: 30,
            options: {
                min: 100,
                max: 2000,
                width: "50%",
                height: 10,
                interval: 100,
                formatter: "Speed {value}",
                tooltipDir: "bottom",
                bgStyle: {
                    "backgroundColor": "#fff",
                    "boxShadow": "inset 0.5px 0.5px 3px 1px rgba(0,0,0,.36)"
                },
                tooltipStyle: {
                    "backgroundColor": "#666",
                    "borderColor": "#666"
                },
                processStyle: {
                    "backgroundColor": "#999"
                }
            },
            config: {speed: 200, size: 10},
            controller: null
        }
    },
    methods: {
        updateSpeed: function () {
            this.config.SetSpeed(this.speed)
            this.controller.UpdateConfig(this.config)
        },
        init() {
            var that = this
            that.config = window.algorithm.ControllerConfig()
            that.config.SetSpeed(that.speed)
            that.config.SetSize(that.size)
            this.controller = window.algorithm.Controller()
            this.controller.Init(that.id, that.config)
            this.controller.Run()
        },
        stop() {
            this.controller.Stop()
        },
        resume() {
            this.controller.Run()
        },
        restart() {
            this.controller.Stop()
            this.init()
        }
    },
    mounted() {
        var that = this
        var avEl = document.getElementById('av')
        if (avEl !== null) {
            if (avEl.hasAttribute('finished')) {
                that.init()
            } else {
                avEl.addEventListener('load', function () {
                    avEl.setAttribute(
                        'finished',
                        'true'
                    )
                    that.init()
                })
            }
            return
        }
        const plugin = document.createElement('script')

        plugin.setAttribute(
            'src',
            this.$withBase('/scripts/main.js?' + Date.now())
        )
        plugin.setAttribute(
            'id',
            'av'
        )
        plugin.addEventListener('load', function () {
            plugin.setAttribute(
                'finished',
                'true'
            )
            that.init()
        })
        // plugin.async = true;
        document.head.appendChild(plugin)
    }
}
</script>

<style lang="stylus">
    canvas
        border 1px solid #DAD7CD
</style>