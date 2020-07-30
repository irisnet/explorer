module.exports = {
    devServer : {
        proxy : {
            '/api' : {
                target : 'http://10.1.4.185:8081',
                secure : false,
            }
        }
    },
};
