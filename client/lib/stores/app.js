var Reflux = require('reflux');

var actions = require('../actions');

var state = {
  name: 'Boiler: Web App'
};

var appStore = Reflux.createStore({

  init: function () {
    this.listenTo(actions.changeName, this.changeName);
  },

  changeName: function (name) {
    state.name = name;
    this.trigger(state);
  },

  getName: function () {
    return state.name;
  }

});

module.exports = appStore;
