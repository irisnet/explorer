module.exports = {
    devServer: {
        proxy: {
          '/api':{
            // target:'http://localhost:8080',
            target:'http://nyancat.irisplorer.io',
            secure:false,
          }
        }
    }
};
