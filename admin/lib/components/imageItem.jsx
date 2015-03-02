'use strict';

var React = require('react');

var ImageItem = React.createClass({

  propTypes: {
    path: React.PropTypes.string.isRequired,
    onRemove: React.PropTypes.func,
  },

  render: function () {
    var path = this.props.path;

    return (
      // <div className='image'>
      <div>
        <img src={'assets/' + path} alt='' />
        <button onClick={this.props.onRemove}>Remove</button>
      </div>
    );
  },

});

module.exports = ImageItem;
