//@ts-check
const TerserPlugin = require('terser-webpack-plugin');
console.log('env:',process.env);
module.exports = {
  devServer: {
    proxy: {
      '/api': {
        target: 'http://localhost:8080',
        secure: false,
      },
    }
  },
  configureWebpack: {
    optimization: {
      minimize: process.env.NODE_ENV === 'production',
      minimizer: [
        new TerserPlugin({
          terserOptions: {
            compress: {
              drop_console: true,
              drop_debugger: true,
            },
          },
        }),
      ],
    },
  },
}
