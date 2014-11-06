'use strict';

var React           = require('react');
var Fluxxor         = require('fluxxor');
var FluxMixin       = Fluxxor.FluxMixin(React);
var StoreWatchMixin = Fluxxor.StoreWatchMixin;

var Header    = require('./header');
var Browser   = require('./browser');
var Editor    = require('./editor');
var Constants = require('../constants');

var App = React.createClass({
  mixins: [FluxMixin, StoreWatchMixin('App', 'Browser')],

  getStateFromFlux: function () {
    var flux = this.getFlux();
    return {
      app: flux.store('App').getState(),
      browser: flux.store('Browser').getState(),
    };
  },

  render: function () {
    var view = null;

    switch (this.state.app.route) {
      case Constants.ROUTE_BROWSER:
        view = <Browser browser={this.state.browser} />
        break;
      case Constants.ROUTE_EDITOR:
        view = <Editor />
        break;
    }

    return (
      <div className='app'>
        <Header />
        {view}
      </div>
    );
  },

});

module.exports = App;
