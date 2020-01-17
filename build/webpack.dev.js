const merge = require('webpack-merge')
const common = require('./webpack.common.js')

module.exports = merge(common, {
  devtool: 'inline-source-map',
  mode: 'development',
  module: {
    rules: [
      {
        test: /\.(css|less)$/,
        exclude: /node_modules/,
        use: [
          'style-loader',
          {
            loader: 'css-loader',
            options: {
              modules: {
                mode: 'local',
                localIdentName: '[path][name]__[local]--[hash:base64:5]'
              }
            }
          },
          {
            loader: 'less-loader',
            options: {
              javascriptEnabled: true
            }
          }
        ]
      },
      {
        test: /\.(css|less)$/,
        include: /node_modules/,
        use: [
          'style-loader',
          'css-loader',
          {
            loader: 'less-loader',
            options: {
              javascriptEnabled: true
            }
          }
        ]
      }
    ]
  },
  devServer: {
    historyApiFallback: true,
    clientLogLevel: 'warning',
    stats: 'errors-only',
    noInfo: true,
    open: true,
    // 如果你想要代理多个路径特定到同一个 target 下，你可以使用由一个或多个「具有 context 属性的对象」构成的数组
    // proxy: [{
    //   context: ['/auth', '/api'],
    //   target: 'http://localhost:3000',
    // }]
    proxy: {
      '/appFile': {
        target: 'http://qa.asoco.ac.cn',
        changeOrigin: true,
        ws: true,
        pathRewrite: {
          '^/appFile': '/appFile'
        }
      }
    }
  }
})
