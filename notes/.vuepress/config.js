module.exports = {
    port : '8083',
    title: 'Programming Lei',
    description: 'All about Programming',
    lastUpdated: 'Last Updated', // string | boolean
    ga: 'UA-81346198-2',
    serviceWorker: true,
    markdown: {
    },
    head: [
        ['link', { rel: 'icon', href: '/favicon.ico' }],
        ['script', { type: 'application/javascript', src: '/scripts/main.js?' + Date.now() }]
    ],
    themeConfig: {
        repo: 'lei-cao/programming',
        nav: [
            {text: 'Home', link: '/'},
            {text: 'Blog', link: '/blog/'},
            {text: 'System Design', link: '/system-design/'}
        ],
        sidebar: {
            '/blog/': [
                {
                    title: 'Blog',
                    collapsable: false,
                    children: [
                        '',
                        'algorithm-visualization',
                        'algorithm-visualization-refactoring'
                    ]
                }
            ],
            '/system-design/': [
                {
                    title: 'System Design',
                    collapsable: false,
                    children: [
                        '',
                        'hadoop/'
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