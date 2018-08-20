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
        ['script', { type: 'application/javascript', src: 'https://pkg.jsgo.io/github.com/lei-cao/programming/code.d6f2627b0d2c3a0ffc07a8574f7705d841bb7f7f.js' }]
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
                        'algorithm-visualization-refactoring',
                        'algoman'
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