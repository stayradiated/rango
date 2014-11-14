'use strict';

var Constants = require('./constants');

module.exports = {

  openDirectory: function (directoryName) {
    this.dispatch(Constants.OPEN_DIRECTORY, directoryName);
  },

  openParentDirectory: function () {
    this.dispatch(Constants.OPEN_PARENT_DIRECTORY);
  },

  openPage: function (fileName) {
    this.dispatch(Constants.OPEN_PAGE, fileName);
  },

  openPath: function (path) {
    if (typeof path !== 'string') {
      throw new Error('openPath: path must be a string');
    }
    this.dispatch(Constants.OPEN_PATH, path);
  },

  createPage: function (name) {
    this.dispatch(Constants.CREATE_PAGE, name);
  },

  createDirectory: function (name) {
    this.dispatch(Constants.CREATE_DIRECTORY, name);
  },

  selectFile: function (fileName) {
    this.dispatch(Constants.SELECT_FILE, fileName);
  },

  deselectAll: function () {
    this.dispatch(Constants.DESELECT_ALL);
  },

  removeSelectedFiles: function () {
    this.dispatch(Constants.REMOVE_SELECTED_FILES);
  },

};
