'use strict';

var Fluxxor = require('fluxxor');
var Constants = require('../constants');

var AppStore = Fluxxor.createStore({

  initialize: function (options) {
    this.name = options.name || '';

    this.bindActions(
      Constants.CHANGE_NAME, this.handleChangeName
    );
  },

  getState: function () {
    return {
      name: this.name,
    };
  },

  handleChangeName: function (payload) {
    this.name = payload.name;
    this.emit('change');
  },

});

module.exports = AppStore;
