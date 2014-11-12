'use strict';

var React = require('react');
var moment = require('moment');

var actions = require('../actions');

var BrowserRow = React.createClass({

  propTypes: {
    item: React.PropTypes.object.isRequired,
    onClick: React.PropTypes.func.isRequired,
  },

  render: function () {
    var item = this.props.item;
    var lastModified = moment(item.modified_at).format('MMM do YY');

    return (
      <tr onClick={this.props.onClick}>
        <td className='type'>📄</td>
        <td className='name'>{item.name}</td>
        <td className='last-modified'>{lastModified}</td>
        <td className='contents'></td>
        <td className='draft'></td>
      </tr>
    );
  },

});

module.exports = BrowserRow;
