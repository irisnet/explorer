module.exports = {
    devServer: {
        proxy: {
          '/api':{
            // target:'http://localhost:8080',
            // target:'http://10.1.4.202:8080',
            target:'https://dev.irisplorer.io',
            secure:false,
          }
        }
    }
};
