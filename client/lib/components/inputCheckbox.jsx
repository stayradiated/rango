'use strict';

var React = require('react');

var InputCheckbox = React.createClass({

  propTypes: {
    label: React.PropTypes.string.isRequired,
  },

  render: function () {
    return (
      <div className='input input-checkbox'>
        <input type='checkbox' />
        <label>{ this.props.label }</label>
      </div>
    );
  },

});

module.exports = InputCheckbox;
