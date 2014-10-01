var React = require('react');
var Reflux = require('reflux');

var appStore = require('../stores/app');
var Header = require('./header.react');

var App = React.createClass({

  mixins: [Reflux.ListenerMixin],

  componentDidMount: function () {
    this.listenTo(appStore, this.onChange);
  },

  onChange: function () {
    this.forceUpdate();
  },

  render: function () {
    /* jshint ignore: start */
    return (
      <div className='app'>
        <Header />
      </div>
    );
    /* jshint ignore: end */
  }

});

module.exports = App;
