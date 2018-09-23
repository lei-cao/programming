<template>
    <div>
        <ClientOnly>
            <LPixi/>
        </ClientOnly>
    </div>
</template>

<script>
    import LPixi from './learningpixi/LearningPixi'

    export default {
        name: "LearningPixi",
        props: {},
        components: {
            LPixi
        },
        data() {
            return {}
        },
        methods: {},
        mounted() {
        }
    }
</script>

<style lang="stylus">
</style>