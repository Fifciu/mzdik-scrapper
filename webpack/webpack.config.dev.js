const Path = require('path');
const Webpack = require('webpack');
const merge = require('webpack-merge');
const common = require('./webpack.common.js');
const {InjectManifest} = require('workbox-webpack-plugin');
const ManifestPlugin = require('webpack-manifest-plugin');

module.exports = merge(common, {
  mode: 'development',
  devtool: 'cheap-eval-source-map',
  output: {
    chunkFilename: 'js/[name].chunk.js'
  },
  devServer: {
    inline: true//,
   /* historyApiFallback: {
      index: '/'
    }*/
  },
  plugins: [
    new Webpack.DefinePlugin({
      'process.env.NODE_ENV': JSON.stringify('development')
    }),
    new InjectManifest({
      swSrc: Path.resolve(__dirname, '../src/service-worker.js')
    }),
    new ManifestPlugin({
      filter: fileDesc => {
        if(fileDesc.path.match(/.*\.js/) || fileDesc.path.match(/public/) || fileDesc.path.match(/.*\.html/)){
          return false;
        }

        return true;
      },
      seed: {
        "name": "MzdikPWA",
        "icons": [
          {
            "src": "\/images\/favicon\/android-icon-36x36.png",
            "sizes": "36x36",
            "type": "image\/png",
            "density": "0.75"
          },
          {
            "src": "\/images/favicon\/android-icon-48x48.png",
            "sizes": "48x48",
            "type": "image\/png",
            "density": "1.0"
          },
          {
            "src": "\/images\/favicon\/android-icon-72x72.png",
            "sizes": "72x72",
            "type": "image\/png",
            "density": "1.5"
          },
          {
            "src": "\/images\/favicon\/android-icon-96x96.png",
            "sizes": "96x96",
            "type": "image\/png",
            "density": "2.0"
          },
          {
            "src": "\/images\/favicon\/android-icon-144x144.png",
            "sizes": "144x144",
            "type": "image\/png",
            "density": "3.0"
          },
          {
            "src": "\/images\/favicon\/android-icon-192x192.png",
            "sizes": "192x192",
            "type": "image\/png",
            "density": "4.0"
          }
        ]
      }
    })
  ],
  module: {
    rules: [
      {
        test: /\.(js)$/,
        include: Path.resolve(__dirname, '../src'),
        enforce: 'pre',
        loader: 'eslint-loader',
        options: {
          emitWarning: true,
        }
      },
      {
        test: /\.(js)$/,
        include: Path.resolve(__dirname, '../src'),
        loader: 'babel-loader'
      },
      {
        test: /\.s?css$/i,
        use: ['style-loader', 'css-loader?sourceMap=true', 'sass-loader']
      }
    ]
  }
});
