'use strict';

var jQuery    = require('jquery');
var Fluxxor   = require('fluxxor');
var Constants = require('../constants');

var AppStore = Fluxxor.createStore({

  initialize: function () {
    this.route = Constants.ROUTE_BROWSER;

    this.bindActions(
      Constants.OPEN_PATH, this.handleOpenBrowser,
      Constants.OPEN_DIRECTORY, this.handleOpenBrowser,
      Constants.OPEN_PAGE, this.handleOpenEditor
    );
  },

  getState: function () {
    return {
      route: this.route,
    };
  },

  handleOpenBrowser: function () {
    this.route = Constants.ROUTE_BROWSER;
    this.emit('change');
  },

  handleOpenEditor: function () {
    this.route = Constants.ROUTE_EDITOR;
    this.emit('change');
  },

});

module.exports = AppStore;
