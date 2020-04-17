module.exports = {
    devServer: {
        proxy: {
          '/api':{
              // target:'http://localhost:8080',
              target:'http://192.168.150.31:30201',
              // target:'http://10.1.4.50:8080',
              // target:'https://www.irisplorer.io',
              // target:'https://stage.irisplorer.io',
              // target:'https://nyancat.irisplorer.io',
              // target:'https://qa.irisplorer.io',
            secure:false,
          },
        }
    },
};
