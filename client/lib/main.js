var $ = require('jquery');
var fastclick = require('fastclick');
var React = require('react');

var App = require('./components/app.react');

$(function () {
  fastclick(document.body);
  React.renderComponent(new App(), document.body);
});

// Trigger React dev tools
if (typeof window !== 'undefined') {
  window.React = React;
}
