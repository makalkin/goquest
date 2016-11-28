var path = require('path');
var autoprefixer = require('autoprefixer');
var webpack = require('webpack');
var HtmlWebpackPlugin = require('html-webpack-plugin');
var ExtractTextPlugin = require('extract-text-webpack-plugin');
var url = require('url');
var paths = require('./paths');

var homepagePath = require(paths.appPackageJson).homepage;
var publicPath = homepagePath ? url.parse(homepagePath).pathname : '/';
if (!publicPath.endsWith('/')) {
  // Prevents incorrect paths in file-loader
  publicPath += '/';
}

module.exports = {
  bail: true,
  devtool: 'source-map',
  entry: [
    require.resolve('./polyfills'),
    path.join(paths.appSrc, 'index')
  ],
  output: {
    path: paths.appBuild,
    filename: 'static/js/[name].js',
    chunkFilename: 'static/js/[name].chunk.js',
    publicPath: publicPath,
    libraryTarget: 'umd'
  },
  //externals: {'react': 'React'},
  externals1: [
    {
      react: {
        root: 'React',
        commonjs2: 'react',
        commonjs: 'react',
        amd: 'react'
      }
    },
    {
      'react-dom': {
        root: 'ReactDOM',
        commonjs2: 'react-dom',
        commonjs: 'react-dom',
        amd: 'react-dom'
      }
    },
    {
      'redux-thunk': {
        root: 'thunkMiddleware',
        commonjs2: 'redux-thunk',
        commonjs: 'redux-thunk',
        amd: 'redux-thunk'
      }
    },
    'react-router',
    'react-redux',
    'redux'
  ],
  resolve: {
    extensions: ['', '.js', '.json'],
    alias: {
      'babel-runtime/regenerator': require.resolve('babel-runtime/regenerator')
    }
  },
  //resolveLoader: {
  //  root: paths.ownNodeModules,
  //  moduleTemplates: ['*-loader']
  //},
  module: {
    preLoaders: [
      {
        test: /\.js$/,
        loader: 'eslint',
        include: paths.appSrc
      }
    ],
    loaders: [
      {
        test: /\.js$/,
        include: paths.appSrc,
        loader: 'babel',
        query: require('./babel.prod')
      },
      {
        test: /\.css$/,
        include: [paths.appSrc, paths.appNodeModules],
        loader: ExtractTextPlugin.extract('style', 'css?-autoprefixer!postcss')
      },
      {
        test: /\.scss$/,
        include: [paths.appSrc, paths.appNodeModules],
        loaders: ['style', 'css', 'sass']
      },
      {
        test: /\.json$/,
        include: [paths.appSrc, paths.appNodeModules],
        loader: 'json'
      },
      {
        test: /\.(jpg|png|gif|eot|svg|ttf|woff|woff2)(\?.*)?$/,
        include: [paths.appSrc, paths.appNodeModules],
        loader: 'file',
        query: {
          name: 'static/media/[name].[hash:8].[ext]'
        }
      },
      {
        test: /\.(mp4|webm)(\?.*)?$/,
        include: [paths.appSrc, paths.appNodeModules],
        loader: 'url',
        query: {
          limit: 10000,
          name: 'static/media/[name].[hash:8].[ext]'
        }
      }
    ]
  },
  eslint: {
    configFile: path.join(__dirname, 'eslint.js'),
    useEslintrc: false
  },
  postcss: function () {
    return [autoprefixer];
  },
  plugins: [
    new HtmlWebpackPlugin({
      inject: true,
      template: paths.appHtml,
      favicon: paths.appFavicon,
      minify: {
        removeComments: true,
        collapseWhitespace: true,
        removeRedundantAttributes: true,
        useShortDoctype: true,
        removeEmptyAttributes: true,
        removeStyleLinkTypeAttributes: true,
        keepClosingSlash: true,
        //minifyJS: true,
        //minifyCSS: true,
        //minifyURLs: true
      }
    }),
    new webpack.DefinePlugin({'process.env.NODE_ENV': '"development"'}),
    new webpack.optimize.OccurrenceOrderPlugin(),
    new webpack.optimize.DedupePlugin(),
    //new webpack.optimize.UglifyJsPlugin({
    //  compress: {
    //    screw_ie8: true,
    //    warnings: false
    //  },
    //  mangle: {
    //    screw_ie8: true
    //  },
    //  output: {
    //    comments: false,
    //    screw_ie8: true
    //  }
    //}),
    new ExtractTextPlugin('static/css/[name].css')
  ]
};
