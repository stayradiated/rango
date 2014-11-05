'use strict';

var React           = require('react');
var Fluxxor         = require('fluxxor');
var FluxMixin       = Fluxxor.FluxMixin(React);
var StoreWatchMixin = Fluxxor.StoreWatchMixin;

var Header  = require('./header');
var Browser = require('./browser');
var Editor  = require('./editor');

var App = React.createClass({

  mixins: [FluxMixin, StoreWatchMixin('AppStore')],

  getStateFromFlux: function () {
    var flux = this.getFlux();
    return {
      app: flux.store('AppStore').getState(),
    };
  },

  render: function () {
    return (
      <div className='app'>
        <Header
          name={this.state.app.name}
          onChangeName={this.onChangeName}
        />
        <Editor />
        <Browser />
      </div>
    );
  },

  onChangeName: function () {
    var name = window.prompt('Enter a name');
    this.getFlux().actions.changeName(name);
  },

});

module.exports = App;
