'use strict';

var $ = require('jquery');
var fastclick = require('fastclick');
var React = require('react');
var Fluxxor = require('fluxxor');

var App = require('./components/app');
var AppStore = require('./stores/app');
var actions = require('./actions');

$(function () {
  fastclick(document.body);

  var stores = {
    AppStore: new AppStore({ name: 'Boiler' }),
  };

  var flux = new Fluxxor.Flux(stores, actions);
  React.renderComponent(<App flux={flux} />, document.body);
});

// Trigger React dev tools
if (typeof window !== 'undefined') {
  window.React = React;
}
