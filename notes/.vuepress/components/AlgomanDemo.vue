<template>
    <div>
        <ClientOnly>
            <Algoman/>
        </ClientOnly>
    </div>
</template>

<script>
    import Algoman from './algoman/Algoman'

    export default {
        name: "AlgomanDemo",
        props: {},
        components: {
            Algoman
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