//@ts-check
module.exports = {
    devServer: {
        proxy: {
          '/api':{
              
              target:'http://localhost:8080',
            secure:false,
          },
        }
    },
};
