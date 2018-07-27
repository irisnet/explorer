module.exports = {
    devServer: {
        proxy: {
          '/api':{
            target:'http://dev.irisplorer.io',
            secure:false,
          }
        }
    }
}
