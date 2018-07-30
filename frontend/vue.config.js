module.exports = {
    devServer: {
        proxy: {
          '/api':{
            target:'http://dev.irisplorer.io',
            //target:'http://192.168.150.235:8080',
            //target:'http:localhost:8080',
            secure:false,
          }
        }
    }
}
