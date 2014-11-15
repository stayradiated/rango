'use strict';

var Immutable = require('immutable');
var Fluxxor   = require('fluxxor');

var Constants = require('../constants');
var Rango     = require('../api');

var ConfigStore = Fluxxor.createStore({

  actions: {
  },

  initialize: function () {
    this.config = Immutable.fromJS({
      types: {
        default: {
        },
      },
    });

    // fetch config from server
    this.fetchConfig();
  },

  getState: function () {
    return {
      config: this.config,
    };
  },

  fetchConfig: function () { 
    var self = this;
    return Rango.readConfig().then(function (config) {
      self.config = Immutable.fromJS(config);
      self.emit('change');
    });
  },

});

module.exports = ConfigStore;
