module.exports = {
    presets: [
        '@vue/app'
    ],
    "presets": [
        ["@babel/preset-env", { "modules": false }]
    ],
    "plugins": [
        [
            "component",
            {
                "libraryName": "element-ui",
                "styleLibraryName": "theme-chalk"
            }
        ],
        [
            "@babel/plugin-transform-runtime",
        ]
    ]
}
