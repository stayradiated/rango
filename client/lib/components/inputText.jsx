'use strict';

var React           = require('react');
var PureRenderMixin = require('react/addons').addons.PureRenderMixin;

var InputText = React.createClass({
  mixins: [
    PureRenderMixin,
  ],

  propTypes: {
    label: React.PropTypes.string.isRequired,
    value: React.PropTypes.string.isRequired,
    onChange: React.PropTypes.func,
  },

  render: function () {
    return (
      <div className='input input-text'>
        <label>{ this.props.label }</label>
        <input
          type='text'
          ref='input'
          value={this.props.value}
          onChange={this.props.onChange}
        />
      </div>
    );
  },

});

module.exports = InputText;
