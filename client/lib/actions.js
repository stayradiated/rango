'use strict';

var Constants = require('./constants');

module.exports = {
  changeName: function (name) {
    this.dispatch(Constants.CHANGE_NAME, {name: name});
  },
};
