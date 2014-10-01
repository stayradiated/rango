var React = require('react');

var actions = require('../actions');
var appStore = require('../stores/app');

var Header = React.createClass({

  changeName: function () {
    var name = window.prompt('Enter a name');
    actions.changeName(name);
  },

  render: function () {
    /* jshint ignore: start */
    return (
      <header>
        <h1>{appStore.getName()}</h1>
        <button onClick={this.changeName}>
          Change Name
        </button>
      </header>
    );
    /* jshint ignore: end */
  }

});

module.exports = Header;
