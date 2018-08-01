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
        var that = this
        var avEl = document.getElementById('av')
        if (avEl !== null) {
            if (avEl.hasAttribute('finished')) {
                var a = window.algorithm.Algorithm()
                a.Display(that.id)
            } else {
                avEl.addEventListener('load', function () {
                    avEl.setAttribute(
                        'finished',
                        'true'
                    )
                    var a = window.algorithm.Algorithm()
                    a.Display(that.id)
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
            var a = window.algorithm.Algorithm()
            a.Display(that.id)
        })
        // plugin.async = true;
        document.head.appendChild(plugin)
    }
}
</script>

<style lang="stylus">

</style>