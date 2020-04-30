module.exports = {
    devServer: {
        proxy: {
          '/api':{
              // target:'http://localhost:8080',
              target:'http://192.168.150.31:30201',
            secure:false,
          },
        }
    },
};
