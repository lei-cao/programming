<template>
    <div>
        <iframe
                id="ebitenIframe"
                src="about:blank"
                frameborder="0"
                width="320"
                height="580"
                allowtransparency
        ></iframe>
    </div>
</template>

<script>
export default {
    name: "Ebiten",
    props: {
        id: String
    },
    mounted() {
        window.loadEbiten()
    }
}
</script>

<style lang="stylus">
</style>