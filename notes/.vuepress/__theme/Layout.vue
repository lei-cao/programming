<template lang="html">
    <section>
        <component :is="layout"></component>
    </section>
</template>

<script>
import Post from './Post'

export default {
    name: 'layout',
    components: {
        Post
    },
    data() {
        return {
        }
    },
    methods: {
    },
    computed: {
        layout() {
            return this.$page.frontmatter.layout || 'post'
        }
    }
}
</script>