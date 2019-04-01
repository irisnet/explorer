module.exports = {
    devServer: {
        proxy: {
          '/api':{
            // target:'http://dev.irisplorer.io',
            target:'http://192.168.150.31:30201/',
            // target:'http://35.236.146.181:8080/',
            // target:'http://34.80.228.1:30201/',
            // target:'http://35.220.138.85:8080/',
            // target:'http://localhost:8080',
            secure:false,
          }
        }
    }
};
