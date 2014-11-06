'use strict';

var jQuery    = require('jquery');
var fastclick = require('fastclick');
var React     = require('react');
var Fluxxor   = require('fluxxor');

var App          = require('./components/app');
var BrowserStore = require('./stores/browser');
var AppStore = require('./stores/app');
var actions      = require('./actions');

jQuery(function () {
  fastclick(document.body);

  var stores = {
    App: new AppStore(),
    Browser: new BrowserStore(),
  };

  var flux = new Fluxxor.Flux(stores, actions);
  React.render(<App flux={flux} />, document.body);
});

// Trigger React dev tools
if (typeof window !== 'undefined') {
  window.React = React;
}
