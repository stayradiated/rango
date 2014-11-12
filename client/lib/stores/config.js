'use strict';

var jQuery    = require('jquery');
var Fluxxor   = require('fluxxor');
var Constants = require('../constants');

var CONFIG_URL = './api/config';

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

    jQuery.get(CONFIG_URL).then(function (response) {
      self.config = response;
      console.log(response);
      self.emit('change');
    });
  },

});

module.exports = ConfigStore;
