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
                v-model="duration"
                v-bind="options"
                @callback="updateDuration"
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
            duration: 300,
            size: 30,
            options: {
                value: 30,
                min: 10,
                max: 2000,
                data: [
                    10,
                    50,
                    100,
                    300,
                    500,
                    1000,
                    2000
                ],
                width: '50%',
                height: 10,
                interval: 10,
                formatter: 'Duration {value}',
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
            config: {},
            controller: null
        }
    },
    methods: {
        updateDuration: function () {
            this.config.SetDuration(this.duration)
            this.controller.UpdateConfig(this.config)
        },
        init() {
            this.config = window.algorithm.ControllerConfig()
            this.config.SetDuration(this.duration)
            this.config.SetSize(this.size)
            this.config.SetId(this.id)
            this.controller = window.algorithm.Controller()
            this.controller.Init(this.config)
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
        this.init()
    }
}
</script>

<style lang="stylus">
    canvas
        border 1px solid #DAD7CD
</style>