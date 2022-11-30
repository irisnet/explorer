const prodPlugins = [];
if(process.env.NODE_ENV === 'production') {
    prodPlugins.push("transform-remove-console");
}
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
        ],
        ...prodPlugins
    ]
}
