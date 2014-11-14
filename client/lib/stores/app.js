'use strict';

var jQuery    = require('jquery');
var Fluxxor   = require('fluxxor');
var Immutable = require('immutable');

var Constants = require('../constants');

var AppStore = Fluxxor.createStore({

  initialize: function () {
    this.state = Immutable.fromJS({
      route: Constants.ROUTE_BROWSER,
    });

    this.bindActions({
      OPEN_PATH:       'handleOpenBrowser',
      OPEN_DIRECTORY:  'handleOpenBrowser',
      OPEN_PAGE:       'handleOpenEditor',
    });
  },

  handleOpenBrowser: function () {
    this.state = this.state.set('route', Constants.ROUTE_BROWSER);
    this.emit('change');
  },

  handleOpenEditor: function () {
    this.state = this.state.set('route', Constants.ROUTE_EDITOR);
    this.emit('change');
  },

});

module.exports = AppStore;
