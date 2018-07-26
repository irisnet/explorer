module.exports = {
    devServer: {
        proxy: {
          '/api':{
            target:'http://192.168.150.122:8080',
            secure:false,
          }
        }
    }
}
