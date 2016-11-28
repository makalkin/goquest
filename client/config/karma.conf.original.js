var path = require('path');

module.exports = function(config) {
  config.set({
    basePath: '',
    frameworks: ['jasmine'],

    plugins: [
      'karma-webpack',
      'karma-jasmine',
      'karma-coverage',
      'karma-sourcemap-loader',
      'karma-chrome-launcher',
      'karma-phantomjs-launcher'
    ],

    files: [
      '../src/**/*.spec.js'
    ],

    preprocessors: {
      // add webpack as preprocessor
      '../src/**/*.spec.js': ['webpack']
      //'src/**/*.js': ['webpack'],
    },

    webpack: { //kind of a copy of your webpack config
      devtool: 'inline-source-map', //just do inline source maps instead of the default
      module: {
        loaders: [
          {
            test: /\.js$/,
            loader: 'babel',
            exclude: path.resolve(__dirname, 'node_modules'),
            query: {
              presets: ['airbnb']
            }
          },
          {
            test: /\.json$/,
            loader: 'json'
          }
        ]
      },
      externals: {
        'cheerio': 'window',
        'react/addons': true,
        'react/lib/ExecutionEnvironment': true,
        'react/lib/ReactContext': true
      }
    },

    webpackServer: {
      noInfo: true //please don't spam the console when running in karma!
    },

    babelPreprocessor: {
      options: {
        presets: ['airbnb']
      }
    },
    reporters: ['progress', 'coverage'],
    port: 9876,
    colors: true,
    logLevel: config.LOG_INFO,
    autoWatch: true,
    browsers: ['PhantomJS'],
    singleRun: false,

    //coverageReporter: {
    //  type : 'html',
    //  dir : 'coverage/'
    //}

    coverageReporter: {
      type: 'html'
    }
  })
};
