'use strict';

var Immutable = require('immutable');
var Fluxxor   = require('fluxxor');

var Rango     = require('../api');

var emptyPage = Immutable.fromJS({
  content: '',
  metadata: {},
});

var EditorStore = Fluxxor.createStore({

  actions: {
    OPEN_PAGE: 'handleOpenPage',
    SAVE_PAGE: 'handleSavePage',
    UPDATE_METADATA: 'handleUpdateMetadata',
    UPDATE_CONTENT: 'handleUpdateContent',
  },

  initialize: function () {

    this.state = Immutable.fromJS({

      // path to the currently edited file
      path: '',

      // the directories and pages in the current path
      page: emptyPage,

    });
  },

  // fetch the page metadata and content from the server
  fetchPage: function () {
    var self = this;

    // empty content while we are waiting
    this.state = this.state.set('page', emptyPage);
    self.emit('change');

    Rango.readPage(this.state.get('path')).then(function (page) {
      self.state = self.state.set('page', Immutable.fromJS(page));
      self.emit('change');
    });
  },

  // open a page in the editor
  handleOpenPage: function (filename) {
    this.state = this.state.set('path', filename);
    this.fetchPage();
  },

  // save the current page back to the server
  handleSavePage: function () {
    var self = this;

    Rango.updatePage(
      this.state.get('path'),
      this.state.get('page').toJS()
    ).then(function (page) {
      self.state = Immutable.fromJS({
        path: page.path,
        page: {
          content: page.content,
          metadata: page.metadata,
        },
      })
      self.emit('change');
    })
  },

  handleUpdateMetadata: function (metadata) {
    this.state = this.state.setIn(['page', 'metadata'], metadata);
    this.emit('change')
  },

  handleUpdateContent: function (content) {
    this.state = this.state.setIn(['page', 'content'], content);
    this.emit('change');
  },

});

module.exports = EditorStore;
