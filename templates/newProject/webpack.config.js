const path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin');

module.exports = {
	entry: "./src/index.js",
	output: {
		path: path.resolve(__dirname, 'dist'),
		filename: "bundle.js"
	},
	module: {
		rules: [
			{test: /\.js$/, use: 'babel-loader' },
			{test: /\.css$/, use: ['style-loader', 'css-loader?modules&importLoaders=1&localIdentName=[name]__[local]___[hash:base64:5]']},
			{test: /\.sass$/, use: ['style-loader', 'css-loader?modules&importLoaders=1&localIdentName=[name]__[local]___[hash:base64:5]', 'sass-loader?indentedSyntax']},
			{
				test: /\.(ico|eot|otf|webp|ttf|woff|woff2)$/i,
				use: `file-loader?limit=100000&name=assets/[name].[hash].[ext]`
			},
			{
				test: /\.(jpe?g|png|gif|svg)$/i,
				use: [
				`file-loader?limit=100000&name=assets/[name].[hash].[ext]`
				// NOTE: it looks like there is an issue using img-loader in some environments
				// {
				// 	loader: 'img-loader',
				// 	options: {
				// 		enabled: true,
				// 		optipng: true
				// 	}
				// }
				]
			}
		]
	},
	devServer: {
		historyApiFallback: true,
		proxy: {
			"/users": {
				target: "http://localhost:1337",
			},
			"/authenticate": {
				target: "http://localhost:1337",
			},
			"/api": {
				target: "http://localhost:1337",
			}
		}
	},
	plugins: [
		new HtmlWebpackPlugin({
			template: 'src/index.html'
		})
	]
	
}
