module.exports = {
	entry: './public/index.js',
	output: {
		path: './public',
		filename: 'bundle.js'
	},
	resolve: {
		alias: {
			'vue$': 'vue/dist/vue.js'
		}
	},
	module: {
		loaders: [{
			test: /\.vue$/,
			loader: 'vue'
		}]
	}
};
