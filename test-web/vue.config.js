module.exports = {
    outputDir: '../dist',
    devServer: {
        proxy: {
            '/api': {
                target: 'http://localhost:1323/api',
                changeOrigin: true,
                pathRewrite: {
                    '^/api': ''
                }
            }
        },
        port: 8080
    }
}