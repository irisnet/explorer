module.exports = {
    devServer: {
        proxy: {
          '/api':{
            // target:'http://dev.irisplorer.io',
             target:'http://47.104.155.125:30201/',
            //target:'http://192.168.150.251:8888',
            secure:false,
          }
        }
    }
};
