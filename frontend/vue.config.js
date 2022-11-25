//@ts-check
console.log('env:',process.env);
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
