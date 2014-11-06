'use strict';

var jQuery    = require('jquery');
var Fluxxor   = require('fluxxor');
var Constants = require('../constants');

var BrowserStore = Fluxxor.createStore({

  initialize: function () {
    this.root = {
      directories: [],
      files: [],
    };
    this.contents = this.root;
    this.path = [];

    var self = this;
    jQuery.get('http://localhost:8080/ls').then(function (response) {
      self.root = response;
      self.contents = response;
      self.emit('change');
    });

    this.bindActions(
      Constants.OPEN_DIRECTORY, this.handleOpenDirectory,
      Constants.OPEN_PARENT_DIRECTORY, this.handleOpenParentDirectory
    );
  },

  getState: function () {
    return {
      root: this.root,
      contents: this.contents,
      isRoot: this.path.length === 0,
    };
  },

  getContentsFromPath: function () {
    var pathList = this.root;
    for (var i = 0, lenI = this.path.length; i < lenI; i++) {
      var directoryName = this.path[i];

      for (var k = 0, lenK = pathList.directories.length; k < lenK; k++) {
        var item = pathList.directories[k];

        if (item.name === directoryName) {
          pathList = item.contents;
          break;
        }
      }
    }
    return pathList;
  },

  handleOpenDirectory: function (directoryName) {
    var dirList = this.contents.directories;
    for (var i = 0, len = dirList.length; i < len; i++) {
      var dirItem = dirList[i];

      if (dirItem.name === directoryName) {
        this.path.push(directoryName);
        this.contents = dirItem.contents;
        this.emit('change');
        break;
      }
    }
  },

  handleOpenParentDirectory: function () {
    this.path = this.path.slice(0, -1);
    this.contents = this.getContentsFromPath();
    this.emit('change');
  },

});

module.exports = BrowserStore;
