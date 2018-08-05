<template>
    <div>
        <div :id="id" class="canvas">
        </div>

        <button @click="stop">Stop</button>
        <button @click="resume">Resume</button>
        <button @click="nextStep"> > </button>
        <input v-model="size"/>
        <button @click="restart">Restart</button>
        <ClientOnly>
            <vue-slider
                ref="slider"
                v-model="velocity"
                v-bind="options"
                @callback="updateVelocity"
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
            velocity: 30,
            size: 30,
            options: {
                value: 30,
                min: 10,
                max: 10000,
                data: [
                    10,
                    20,
                    30,
                    50,
                    80,
                    150,
                    300,
                    500,
                    1000,
                    7000,
                    10000
                ],
                width: '50%',
                height: 10,
                interval: 10,
                formatter: 'Velocity {value}',
                tooltipDir: 'bottom',
                bgStyle: {
                    'backgroundColor': '#fff',
                    'boxShadow': 'inset 0.5px 0.5px 3px 1px rgba(0,0,0,.36)'
                },
                tooltipStyle: {
                    'backgroundColor': '#666',
                    'borderColor': '#666'
                },
                processStyle: {
                    'backgroundColor': '#999'
                }
            },
            config: {velocity: 30, size: 30},
            controller: null
        }
    },
    methods: {
        updateVelocity: function () {
            this.config.SetVelocity(this.velocity)
            this.controller.UpdateConfig(this.config)
        },
        init() {
            var that = this
            that.config = window.algorithm.ControllerConfig()
            that.config.SetVelocity(that.velocity)
            that.config.SetSize(that.size)
            that.config.SetId(that.id)
            this.controller = window.algorithm.Controller()
            this.controller.Init(that.config)
        },
        stop() {
            this.controller.Stop()
        },
        resume() {
            this.controller.Resume()
        },
        nextStep() {
            this.controller.NextStep()
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