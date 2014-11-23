'use strict';

var jQuery    = require('jquery');
var Fluxxor   = require('fluxxor');
var Immutable = require('immutable');

var api = require('../api');

var AppStore = Fluxxor.createStore({

  actions: {
    OPEN_PATH:       'handleOpenBrowser',
    OPEN_DIRECTORY:  'handleOpenBrowser',
    OPEN_PAGE:       'handleOpenEditor',
    PUBLISH_SITE:    'handlePublishSite',
  },

  initialize: function () {
    this.state = Immutable.fromJS({
      route: 'ROUTE_BROWSER',
    });
  },

  handleOpenBrowser: function () {
    this.state = this.state.set('route', 'ROUTE_BROWSER');
    this.emit('change');
  },

  handleOpenEditor: function () {
    this.state = this.state.set('route', 'ROUTE_EDITOR');
    this.emit('change');
  },

  handlePublishSite: function () {
    api.publishSite().then(function (res) {
      console.log(res);
    });
  },

});

module.exports = AppStore;
