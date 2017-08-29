import webpack from 'webpack'
import { resolve } from 'path'
import HtmlWebpackPlugin from 'html-webpack-plugin'

const HtmlWebpackPluginConfig = new HtmlWebpackPlugin({
	template: resolve(__dirname, 'src', 'index.html'),
	filename: 'index.html',
	inject: 'body',
})

const PATHS = {
	entry: resolve(__dirname, "src"),
	app: resolve(__dirname, "src", "app"),
	build: resolve(__dirname, "dist"),
}

const productionPlugin = new webpack.DefinePlugin({
	'process.env': {
		NODE_ENV: JSON.stringify('production')
	}
})

// Set up environment things
const LaunchCommand = process.env.npm_lifecycle_event
const isProduction = LaunchCommand === 'build'
process.env.BABEL_ENV = LaunchCommand

const baseConfig = {
	context: resolve(__dirname, "src"),
	entry: PATHS.entry,
	output: {
		path: PATHS.build,
		filename: "bundle.js"
	},
	module: {
		rules: [
			{test: /\.js$/, use: 'babel-loader' },
			{test: /\.css$/, use: ['style-loader', 'css-loader?modules&importLoaders=1&localIdentName=[name]__[local]___[hash:base64:5]']},
			{test: /\.sass$/, use: ['style-loader', 'css-loader?modules&importLoaders=1&localIdentName=[name]__[local]___[hash:base64:5]', 'sass-loader?indentedSyntax']},
			{
				test: /\.(ico|eot|otf|webp|ttf|woff|woff2)$/i,
				use: `file-loader?limit=100000&name=assets/[name].[hash].[ext]`,
			},
			{
				test: /\.(jpe?g|png|gif|svg)$/i,
				use: [
				`file-loader?limit=100000&name=assets/[name].[hash].[ext]`,
				],
			},
		],
	},
	resolve: {
		alias: {
			state: resolve(PATHS.app, "state"),
			components: resolve(PATHS.app, "components"),
			utils: resolve(PATHS.app, "utils"),
		},
	},
}


const devConfig = {
	devtool: 'cheap-module-inline-source-map',
	devServer: {
		contentBase: PATHS.build,
		hot: true,
		inline: true,
	},
	plugins: [HtmlWebpackPluginConfig, new webpack.HotModuleReplacementPlugin()],
}

const prodConfig = {
	devtool: 'cheap-module-source-map',
	plugins: [HtmlWebpackPluginConfig, productionPlugin],
}

const configToUse = isProduction ? prodConfig : devConfig

export default {...baseConfig, ...configToUse}
