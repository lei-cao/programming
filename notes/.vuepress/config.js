module.exports = {
    port : '8083',
    title: 'Learning CS Again',
    description: 'Learning Computer Science Once Again',
    lastUpdated: 'Last Updated', // string | boolean
    repo: 'lei-cao/learning-cs-again',
    base: '/learning-cs-again/',
    ga: 'UA-81346198-2',
    markdown: {
    },
    themeConfig: {
        nav: [
            {text: 'Home', link: '/'},
            {text: 'Algorithms', link: '/algorithms/'},
            {text: 'Blog', link: '/blog/'},
        ],
        sidebar: {
            '/blog/': [
                {
                    title: 'Blog',
                    collapsable: false,
                    children: [
                        '',
                        'setting-sail',
                        'algorithm-visualization',
                        'algorithm-visualization-refactoring'
                    ]
                }
            ]
        }
    }
}