'use strict';

var Fluxxor = require('fluxxor');

var Constants = require('../constants');
var Rango     = require('../api');

var ConfigStore = Fluxxor.createStore({

  initialize: function () {
    this.config = {
      types: {
        default: {
        },
      },
    };

    // fetch contents of root directory
    this.fetchConfig();

    // listen to actions
    this.bindActions(
    );
  },

  getState: function () {
    return {
      config: this.config,
    };
  },

  fetchConfig: function () { 
    var self = this;
    return Rango.readConfig().then(function (config) {
      self.config = config;
      self.emit('change');
    });
  },

});

module.exports = ConfigStore;
