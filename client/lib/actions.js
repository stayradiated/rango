'use strict';

var Constants = require('./constants');

module.exports = {

  openDirectory: function (directoryName) {
    this.dispatch(Constants.OPEN_DIRECTORY, directoryName);
  },

  openParentDirectory: function () {
    this.dispatch(Constants.OPEN_PARENT_DIRECTORY);
  },

  openFile: function (fileName) {
    this.dispatch(Constants.OPEN_FILE, fileName);
  },

};
