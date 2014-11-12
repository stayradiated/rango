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
  mixins: [FluxMixin, StoreWatchMixin('App', 'Browser', 'Editor')],

  getStateFromFlux: function () {
    var flux = this.getFlux();
    return {
      app: flux.store('App').getState(),
      editor: flux.store('Editor').getState(),
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
        view = <Editor editor={this.state.editor} />
        break;
    }

    return (
      <div className='app'>
        <Header browser={this.state.browser} />
        {view}
      </div>
    );
  },

});

module.exports = App;
