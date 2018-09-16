var fs    = require("fs")
const path = require('path')
const HtmlWebpackPlugin = require('html-webpack-plugin')
const CleanWebpackPlugin = require('clean-webpack-plugin')
const CopyWebpackPlugin = require('copy-webpack-plugin')

module.exports = {
    entry: {
         app: './src/index.js',
         print: './src/print.js'
    },
    plugins: [
        new CleanWebpackPlugin(['dist']),
        new HtmlWebpackPlugin({
            title: 'Output Management'
        }),
        new CopyWebpackPlugin([
            { from: 'src/.well-known', to: '.well-known' }
        ])
    ],
    output: {
        filename: '[name].bundle.js',
        path: path.resolve(__dirname, 'dist'),
        publicPath: '/'
    },
    devServer: {
        contentBase: './dist',
        before: function(app) {
            app.get('/key.cert', function(req, res) {
                res.sendFile('/Users/l.cao/programming/cert/localhost+4-key.pem');
            });
        },
        host: 'lei.io',
        port: '443',
        https: {
            key: fs.readFileSync('/Users/l.cao/programming/cert/localhost+4-key.pem'),
            cert: fs.readFileSync('/Users/l.cao/programming/cert/localhost+4.pem')
        }
    }
};
