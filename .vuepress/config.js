module.exports = {
    port : '8083',
    title: 'Learning CS Again',
    description: 'Learning Computer Science Once Again',
    lastUpdated: 'Last Updated', // string | boolean
    repo: 'lei-cao/learning-cs-again',
    base: '/learning-cs-again/',
    markdown: {
    },
    themeConfig: {
        nav: [
            {text: 'Home', link: '/'},
            {text: 'Blog', link: '/blog/'},
        ],
        sidebar: {
            '/blog/': [
                {
                    title: 'Blog',
                    collapsable: false,
                    children: [
                        '',
                        'setting-sail'
                    ]
                }
            ]
        }
    }
}