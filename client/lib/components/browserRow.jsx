'use strict';

var React = require('react');

var BrowserRow = React.createClass({

  propTypes: {
    name: React.PropTypes.string.isRequired,
    folder: React.PropTypes.string.isRequired,
    lastModified: React.PropTypes.string.isRequired,
  },

  render: function () {
    var icon = (this.props.type === 'folder' ? '📁' : '📄');

    return (
      <tr>
        <td className='type'>{ icon }</td>
        <td className='name'>{ this.props.name }</td>
        <td className='last-modified'>{ this.props.lastModified }</td>
        <td className='contents'>11 Pages</td>
        <td className='draft'></td>
      </tr>
    );
  },

});

module.exports = BrowserRow;
