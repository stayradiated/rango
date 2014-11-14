'use strict';

var Immutable = require('immutable');
var Fluxxor   = require('fluxxor');
var Path      = require('path')
var jQuery    = require('jquery');

var Constants = require('../constants');
var Rango     = require('../api');

var BrowserStore = Fluxxor.createStore({

  initialize: function () {

    this.state = Immutable.fromJS({

      // a list storing the current path stored in an array
      // /usr/local/bin => ['usr', 'local', 'bin']
      path: [],

      // the directories and pages in the current path
      contents: {
        directories: [],
        pages: [],
      },

      // a set of selected pages and directories
      selected: Immutable.Set(),

    });

    // fetch contents of root directory
    this.fetchContents();

    this.bindActions({
      OPEN_PATH:              'handleOpenPath',
      OPEN_DIRECTORY:         'handleOpenDirectory',

      CREATE_PAGE:            'handleCreatePage',
      CREATE_DIRECTORY:       'handleCreateDirectory',

      OPEN_PARENT_DIRECTORY:  'handleOpenParentDirectory',
      SELECT_FILE:            'handleSelectFile',
      DESELECT_ALL:           'handleDeselectAll',
      REMOVE_SELECTED_FILES:  'handleRemoveSelectedFiles',
    });
  },

  // get the current path as a string, also handles root directory properly
  getPath: function () {
    var path = this.state.get('path');

    if (path.size === 0) {
      return '/';
    } else {
      return path.join('/');
    }
  },

  // get contents of current path from server
  fetchContents: function () {
    var self = this;
    return Rango.readDir(this.getPath()).then(function (data) {
      self.state = self.state.merge({
        contents: Immutable.fromJS(data),
        selected: self.state.get('selected').clear(),
      });
      self.emit('change');
    });
  },

  // change the path
  handleOpenPath: function (newPath) {
    this.state = this.state.update('path', function (path) {
      if (newPath === '/') {
        return path.clear();
      }
      return Immutable.List(newPath.split('/'));
    });

    this.fetchContents();
  },

  // open a sub-directory
  handleOpenDirectory: function (dirName) {
    this.state = this.state.update('path', function (path) {
      return path.push(dirName);
    });
    this.fetchContents();
  },

  // open the parent directory
  handleOpenParentDirectory: function () {
    this.state = this.state.update('path', function (path) {
      return path.pop();
    });
    this.fetchContents();
  },

  // create a new directory on the server
  handleCreateDirectory: function () {
    var self = this;

    var name = window.prompt('Enter a name for the new directory');
    if (! name) { return; }

    var path = Path.join(this.getPath(), name);

    return Rango.createDir(path).then(function () {
      self.fetchContents();
    });
  },

  // create a new page on the server
  handleCreatePage: function () {
    var self = this;

    var name = window.prompt('Enter a title for the new page');
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

  // select a file
  handleSelectFile: function (file) {
    this.state = this.state.update('selected', function (selected) {
      return selected.clear().add(file);
    });
    this.emit('change');
  },

  // deselect all files
  handleDeselectAll: function () {
    this.state = this.state.update('selected', function (selected) {
      return selected.clear();
    });
    this.emit('change');
  },

  // delete the selected files
  handleRemoveSelectedFiles: function () {
    if (! window.confirm('Are you sure you want to delete the selected files?')) {
      return;
    }

    var self = this;
    var path = this.getPath();

    var deferreds = this.state.get('selected').map(function (file) {
      return Rango.destroy(file.get('path'));
    }).toArray();

    jQuery.when.apply(jQuery, deferreds).done(function () {
      self.fetchContents();
    });
  },

});

module.exports = BrowserStore;
