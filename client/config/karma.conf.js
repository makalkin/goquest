var path = require('path');

module.exports = function(config) {
  config.set({
    browsers: ['PhantomJS'],
    files: [
      '../src/**/*.spec.js'
    ],
    frameworks: ['jasmine'],
    preprocessors: {
      '../src/**/*.spec.js': ['webpack']
    },
    reporters: ['progress', 'coverage'],
    webpack: {
      devtool: 'inline-source-map',
      entry: {},
      module: {
        preLoaders: [
          {
            test: /\.jsx?$/,
            exclude: [
              /node_modules/,
              /\.spec\.js/
            ],
            loader: 'isparta-instrumenter-loader',
            query: {
              babel: {
                presets: require('./babel.dev').presets
              }
            }
          }
        ],
        loaders: [
          {
            test: /\.jsx?$/,
            loader: 'babel',
            include: path.resolve(__dirname, '../src'),
            query: require('./babel.dev')
          },
          {
            test: /\.jsx?$/,
            loader: 'babel',
            include: path.resolve(__dirname, '../src'),
            query: {
              presets: ['airbnb']
            }
          },
          {
            test: /\.json$/,
            loader: 'json'
          }
        ]
        //postLoaders: [
        //  {
        //    test: /\.js$/,
        //    include: path.resolve('./src/'),
        //    loader: 'istanbul-instrumenter',
        //    query: {
        //      esModules: true
        //    }
        //  }
        //]
      },
      externals: {
        'cheerio': 'window',
        'react/addons': true,
        'react/lib/ExecutionEnvironment': true,
        'react/lib/ReactContext': true
      }
    },
    babelPreprocessor: {
      options: {
        presets: ['airbnb']
      }
    },
    webpackServer: {
      noInfo: true
    },
    singleRun: false,
    autoWatch: true,
    coverageReporter: {
      dir: '../build/coverage/',
      reporters: [
        { type: 'html', subdir: 'html' }
      ]
    }
  });
};
