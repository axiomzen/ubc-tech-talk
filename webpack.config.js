const webpack = require('webpack');
const path = require('path');

const SRC = path.resolve(__dirname, 'src');
const PUBLIC = path.resolve(__dirname, 'public');

module.exports = {
  entry: ['whatwg-fetch', SRC + '/index.jsx'],
  output: {
    path: PUBLIC,
    filename: 'index.js'
  },
  module: {
    loaders: [{
      include: SRC,
      loader: 'babel',
      query: {
        presets: [
          'es2015'
        ]
      }
    }

    ]
  }
};
