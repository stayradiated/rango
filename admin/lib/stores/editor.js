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
    OPEN_PAGE:        'handleOpenPage',
    SAVE_PAGE:        'handleSavePage',
    UPDATE_METADATA:  'handleUpdateMetadata',
    UPDATE_CONTENT:   'handleUpdateContent',
    UPLOAD_FILE:      'handleUploadFile',
  },

  initialize: function () {

    this.state = Immutable.fromJS({

      // path to the currently edited file
      path: '',

      // the directories and pages in the current path
      page: emptyPage,

      // if the page has changed since it was last loaded
      hasChanged: false,

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
      self.state = self.state.set('hasChanged', false);
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
        hasChanged: false,
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
    this.state = this.state.set('hasChanged', true);
    this.emit('change')
  },

  handleUpdateContent: function (content) {
    this.state = this.state.setIn(['page', 'content'], content);
    this.state = this.state.set('hasChanged', true);
    this.emit('change');
  },
  
  handleUploadFile: function (file) {
    console.log('Uploading', file);
    Rango.createAsset(file);
  },

});

module.exports = EditorStore;
