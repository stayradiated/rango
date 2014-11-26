'use strict';

var React           = require('react');
var Fluxxor         = require('fluxxor');
var FluxMixin       = Fluxxor.FluxMixin(React);
var StoreWatchMixin = Fluxxor.StoreWatchMixin;

var Header    = require('./header');
var Browser   = require('./browser');
var Editor    = require('./editor');

var App = React.createClass({
  mixins: [
    FluxMixin,
    StoreWatchMixin('App', 'Browser', 'Editor'),
  ],

  getStateFromFlux: function () {
    var flux = this.getFlux();
    return {
      app: flux.store('App').state,
      editor: flux.store('Editor').state,
      browser: flux.store('Browser').state,
    };
  },

  render: function () {
    var view = null;

    switch (this.state.app.get('route')) {
      case 'BROWSER':
        view = <Browser browser={this.state.browser} />
        break;
      case 'EDITOR':
        view = <Editor editor={this.state.editor} />
        break;
    }

    return (
      <div className='app'>
        <Header
          app={this.state.app}
          editor={this.state.editor}
          browser={this.state.browser}
        />
        {view}
      </div>
    );
  },

});

module.exports = App;
