module.exports = {
    port : '8083',
    title: 'Learning CS Again',
    description: 'Learning Computer Science Once Again',
    lastUpdated: 'Last Updated', // string | boolean
    repo: 'lei-cao/lei-cao.github.io',
    base: '/',
    ga: 'UA-81346198-2',
    serviceWorker: true,
    markdown: {
    },
    head: [
        ['link', { rel: 'icon', href: '/favicon.ico' }],
        ['script', { type: 'application/javascript', src: '/scripts/main.js?' + Date.now() }]
    ],
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
    },
    configureWebpack: (config, isServer) => {
        if (!isServer) {
            // mutate the config for client
            externals: {
                algorithm: 'algorithm'
            }
        }
    }
}