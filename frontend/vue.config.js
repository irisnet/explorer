module.exports = {
    devServer: {
        proxy: {
          '/api':{
            // target:'http://localhost:8080',
            // target:'http://192.168.150.31:30201',
              //qa
           // target:'http://35.236.146.181:8080/',
           // target:'http://35.236.151.88:8080/',
           // target:'http://35.236.151.88:8080/',
            // target:'http://10.2.10.140:30201/',
            // target:'http://34.80.141.14:8080',
            // target:'http://qa.irisplorer.io',
            // target:'http://nyancat.irisplorer.io',
            // target:'http://testnet.irisplorer.io',
            // target:'http://34.80.141.14:8080',
            //  target:'https://qa2.irisplorer.io',
            // target:'https://www.irisplorer.io',
            target:'https://stage.irisplorer.io',
            secure:false,
          }
        }
    },
    css:{
        loaderOptions:{
            sass:{
                data:"@import'~@/style/mixin.scss';"
            }
        }
    },
  /*  loaders: [
        { test: /iview.src.*?js$/, loader: 'babel' },
        { test: /\.js$/, loader: 'babel', exclude: /node_modules/ }
    ]*/
};
