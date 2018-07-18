<template>
    <div id="body">
    </div>
</template>

<script>

export default {
    name: "ACanvas",
    props: {
        id: String
    },
    mounted() {
        var avEl = document.getElementById('av')
        if (avEl !== null) {
            var a = window.av.New()
            a.Show()
        }
        const plugin = document.createElement('script')
        plugin.onload = function () {
            var a = window.av.New()
            a.Show()
        };
        plugin.setAttribute(
            'src',
            this.$withBase('/scripts/main.js')
        );
        plugin.setAttribute(
            'id',
            'av'
        );
        plugin.async = true;
        document.head.appendChild(plugin);
    }
}
</script>

<style lang="stylus">

</style>