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

    this.bindActions(
      Constants.OPEN_PATH, this.handleOpenBrowser,
      Constants.OPEN_DIRECTORY, this.handleOpenBrowser,
      Constants.OPEN_PAGE, this.handleOpenEditor
    );
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
