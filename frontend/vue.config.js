module.exports = {
    devServer: {
        proxy: {
          '/api':{
            // target:'http://localhost:8080',
            // target:'http://localhost:8080',
            // target:'http://localhost:8080',
            // target:'http://192.168.150.31:30201',
            // qa
            // target:'http://35.236.146.181:8080',
            // target:'http://34.80.141.14:8080',
            // target:'http://10.1.4.166:8080',
            target:'http://35.220.142.249:8080',
            // target:'http://localhost:8080',
            // target:'http://10.1.4.115:8080',
            secure:false,
          }
        }
    }
};
