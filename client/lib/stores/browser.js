'use strict';

var jQuery    = require('jquery');
var Fluxxor   = require('fluxxor');
var Constants = require('../constants');

var PAGES_URL = './api/files/';

var BrowserStore = Fluxxor.createStore({

  initialize: function () {
    // a stack storing the current path stored in an array
    // /usr/local/bin => ['usr', 'local', 'bin']
    this.path = [];

    // the directories and pages in the current path
    this.contents = {
      directories: [],
      pages: [],
    };

    // fetch contents of root directory
    this.fetchContents();

    // listen to actions
    this.bindActions(
      Constants.OPEN_PATH, this.handleOpenPath,
      Constants.OPEN_DIRECTORY, this.handleOpenDirectory,
      Constants.OPEN_PARENT_DIRECTORY, this.handleOpenParentDirectory
    );
  },

  getState: function () {
    return {
      path:      this.path,
      contents:  this.contents,
      isRoot:    this.path.length === 0,
    };
  },

  fetchContents: function () {
    var self = this;

    var path = this.path.join('/');
    var url = PAGES_URL + path;

    jQuery.get(url).then(function (response) {
      self.contents = response;
      self.emit('change');
    });
  },

  handleOpenPath: function (path) {
    this.path = path;
    this.fetchContents();
  },

  handleOpenDirectory: function (dirName) {
    this.path.push(dirName);
    this.fetchContents();
  },

  handleOpenParentDirectory: function () {
    this.path.pop();
    this.fetchContents();
  },

});

module.exports = BrowserStore;
