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

      // a stack storing the current path stored in an array
      // /usr/local/bin => ['usr', 'local', 'bin']
      path: Immutable.List(),

      // the directories and pages in the current path
      contents: Immutable.fromJS({
        directories: [],
        pages: [],
      }),

      // a set of selected pages and directories
      selected: Immutable.Set(),

    });

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

  getPath: function () {
    var path = this.state.get('path');

    if (path.size === 0) {
      return '/';
    } else {
      return path.join('/');
    }
  },

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

  handleOpenPath: function (newPath) {
    this.state = this.state.update('path', function (path) {
      if (newPath === '/') {
        return path.clear();
      }
      return Immutable.List(newPath.split('/'));
    });

    this.fetchContents();
  },

  handleOpenDirectory: function (dirName) {
    this.state = this.state.update('path', function (path) {
      return path.push(dirName);
    });
    this.fetchContents();
  },

  handleOpenParentDirectory: function () {
    this.state = this.state.update('path', function (path) {
      return path.pop();
    });
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

  handleSelectFile: function (file) {
    this.state = this.state.update('selected', function (selected) {
      return selected.clear().add(file);
    });
    this.emit('change');
  },

  handleDeselectAll: function () {
    this.state = this.state.update('selected', function (selected) {
      return selected.clear();
    });
    this.emit('change');
  },

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
