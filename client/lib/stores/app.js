'use strict';

var jQuery    = require('jquery');
var Fluxxor   = require('fluxxor');
var Constants = require('../constants');

var AppStore = Fluxxor.createStore({

  initialize: function () {
    this.route = Constants.ROUTE_BROWSER;

    this.bindActions(
      Constants.OPEN_FILE, this.handleOpenFile
    );
  },

  getState: function () {
    return {
      route: this.route,
    };
  },

  handleOpenFile: function () {
    this.route = Constants.ROUTE_EDITOR;
    this.emit('change');
  },

});

module.exports = AppStore;
