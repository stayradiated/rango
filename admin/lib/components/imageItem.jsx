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
        <img src={'assets' + path} alt='' />
        <button
          onMouseDown={this.handleMouseDown}
          onClick={this.props.onRemove}
        >Remove</button>
      </div>
    );
  },

  handleMouseDown: function (event) {
    // stop sortable from firing
    event.stopPropagation();
  },

});

module.exports = ImageItem;
