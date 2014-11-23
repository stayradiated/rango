'use strict'; 

var jQuery    = require('jquery');
var fastclick = require('fastclick');
var React     = require('react');
var Fluxxor   = require('fluxxor');

var actions      = require('./actions');
var App          = require('./components/app');
var AppStore     = require('./stores/app');
var ConfigStore  = require('./stores/config');
var EditorStore  = require('./stores/editor');
var BrowserStore = require('./stores/browser');

jQuery(function () {
  fastclick(document.body);

  var stores = {
    App: new AppStore(),
    Config: new ConfigStore(),
    Editor: new EditorStore(),
    Browser: new BrowserStore(),
  };

  var flux = new Fluxxor.Flux(stores, actions);
  React.render(<App flux={flux} />, document.body);
});

// Trigger React dev tools
if (typeof window !== 'undefined') {
  window.React = React;
}
