const { createProxyMiddleware } = require('http-proxy-middleware')

module.exports = function(app) {
  // /api 表示代理路径
  // target 表示目标服务器的地址
  app.use(
    createProxyMiddleware('/api', {
      'target': 'http://192.168.50.109:8080',
      'changeOrigin':true,
      'pathRewrite': {
        '^/api': '/'
      }
    })
  )
}