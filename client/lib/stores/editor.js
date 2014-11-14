'use strict';

var Fluxxor = require('fluxxor');

var Constants = require('../constants');
var Rango     = require('../api');

var EditorStore = Fluxxor.createStore({

  initialize: function () {
    // a stack storing the current path stored in an array
    // /usr/local/bin => ['usr', 'local', 'bin']
    this.path = '';

    // the directories and pages in the current path
    this.page = {
      metadata: {},
      content: '',
    };

    // listen to actions
    this.bindActions(
      Constants.OPEN_PAGE, this.handleOpenPage
    );
  },

  getState: function () {
    return {
      page: this.page,
    };
  },

  fetchPage: function () {
    var self = this;

    // empty content while we are waiting
    this.page = {
      metadata: {},
      content: '',
    };
    self.emit('change');

    Rango.readPage(this.path).then(function (page) {
      self.page = page ;
      self.emit('change');
    });
  },

  handleOpenPage: function (filename) {
    this.path = filename;
    this.fetchPage();
  },

});

module.exports = EditorStore;
