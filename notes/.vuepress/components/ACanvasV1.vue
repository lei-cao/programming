<template>
    <div :id="id">
    </div>
</template>

<script>
export default {
    name: "ACanvas",
    props: {
        id: String
    },
    mounted() {
        var a = window.algorithm.Algorithm()
        a.Display(this.id)
    }
}
</script>

<style lang="stylus">

</style>