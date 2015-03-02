'use strict';

var React           = require('react');
var PureRenderMixin = require('react/addons').addons.PureRenderMixin;

var InputDropdown = React.createClass({
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
        <select
          ref='input'
          value={this.props.value}
          onChange={this.props.onChange}
        >
          <option value="standard">Standard</option>
          <option value="carousel">Carousel</option>
          <option value="contact">Contact</option>
        </select>
      </div>
    );
  },

});

module.exports = InputDropdown;
