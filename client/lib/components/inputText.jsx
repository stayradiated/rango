'use strict';

var React = require('react');

var InputText = React.createClass({

  propTypes: {
    label: React.PropTypes.string.isRequired,
    value: React.PropTypes.string.isRequired,
  },

  render: function () {
    return (
      <div className='input input-text'>
        <label>{ this.props.label }</label>
        <input type='text' defaultValue={ this.props.value } />
      </div>
    );
  },

});

module.exports = InputText;
