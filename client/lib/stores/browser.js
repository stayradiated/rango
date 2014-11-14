'use strict';

var Immutable = require('immutable');
var Fluxxor   = require('fluxxor');
var Path      = require('path')
var jQuery    = require('jquery');

var Constants = require('../constants');
var Rango     = require('../api');

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

    // a set of selected pages and directories
    this.selected = Immutable.Set();

    // fetch contents of root directory
    this.fetchContents();

    // listen to actions
    this.bindActions(
      Constants.OPEN_PATH, this.handleOpenPath,
      Constants.OPEN_DIRECTORY, this.handleOpenDirectory,

      Constants.CREATE_PAGE, this.handleCreatePage,
      Constants.CREATE_DIRECTORY, this.handleCreateDirectory,

      Constants.OPEN_PARENT_DIRECTORY, this.handleOpenParentDirectory,
      Constants.SELECT_FILE, this.handleSelectFile,
      Constants.DESELECT_ALL, this.handleDeselectAll,
      Constants.REMOVE_SELECTED_FILES, this.handleRemoveSelectedFiles
    );
  },

  getState: function () {
    return {
      path:      this.path,
      contents:  this.contents,
      selected:  this.selected,
      isRoot:    this.path.length === 0,
    };
  },

  getPath: function () {
    return this.path.join('/') || '/';
  },

  fetchContents: function () {
    var self = this;
    return Rango.readDir(this.getPath()).then(function (data) {
      self.contents = data;
      self.selected = self.selected.clear();
      self.emit('change');
    });
  },

  handleOpenPath: function (path) {
    this.path = path.split('/');

    // remove empty strings from end of path
    var i = this.path.length - 1;
    while (i >= 0 && this.path[i].length <= 0) {
      this.path.pop();
      i -= 1;
    }

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

  handleCreateDirectory: function () {
    var self = this;

    var name = window.prompt('Enter a name for the new directory');
    if (! name) { return; }

    var path = Path.join(this.getPath(), name);

    return Rango.createDir(path).then(function () {
      self.fetchContents();
    });
  },

  handleCreatePage: function () {
    var self = this;

    var name = window.prompt('Enter a name for the new directory');
    if (! name) { return; }

    return Rango.createPage(this.getPath(), {
      metadata: {
        title: name,
      },
      content: ''
    }).then(function (res) {
      if (res.success !== true) {
        throw new Error('Could not create directory');
      }
      self.fetchContents();
    });
  },

  handleSelectFile: function (fileName) {
    this.selected = this.selected.clear().add(fileName);
    this.emit('change');
  },

  handleDeselectAll: function () {
    this.selected = this.selected.clear();
    this.emit('change');
  },

  handleRemoveSelectedFiles: function () {
    if (! window.confirm('Are you sure you want to delete the selected files?')) {
      return;
    }

    var self = this;
    var path = this.path.join('/');

    var deferreds = this.selected.map(function (fileName) {
      return Rango.destroy(Path.join(path, fileName));
    }).toArray();

    jQuery.when.apply(jQuery, deferreds).done(function () {
      self.fetchContents();
    });
  },

});

module.exports = BrowserStore;
