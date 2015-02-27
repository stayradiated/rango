'use strict';

var React           = require('react');
var moment          = require('moment');
var classSet        = require('react/addons').addons.classSet;
var PureRenderMixin = require('react/addons').addons.PureRenderMixin;
var Fluxxor         = require('fluxxor');
var FluxMixin       = Fluxxor.FluxMixin(React);

var actions = require('../actions');

var BrowserRow = React.createClass({
  mixins: [
    FluxMixin,
    PureRenderMixin,
  ],

  propTypes: {
    item: React.PropTypes.object.isRequired,
    selected: React.PropTypes.bool.isRequired,
  },

  render: function () {
    var item = this.props.item;
    var lastModified = moment(item.get('modTime') * 1000).format('LT ll');
    var isDir = item.get('isDir');

    var icon = isDir ? (
      <span className='icon-folder' />
    ) : (
    <span className='icon-doc-text'/>
    );

    var classes = classSet({
      selected: this.props.selected,
    });

    return (
      <tr
        className={classes}
        onClick={this.onClick}
        onDoubleClick={this.onDoubleClick}
      >
        <td className='type'>{icon}</td>
        <td className='name'>{item.get('name')}</td>
        <td className='last-modified'>{lastModified}</td>
        <td className='contents'></td>
        <td className='draft'></td>
      </tr>
    );
  },

  onClick: function (e) {
    e.stopPropagation();
    this.getFlux().actions.select.file(this.props.item);
  },

  onDoubleClick: function () {
    var actions = this.getFlux().actions;

    if (this.props.item.get('isDir')) {
      actions.open.path(this.props.item.get('path'));
    } else {
      actions.open.page(this.props.item.get('path'));
    }
  },

});

module.exports = BrowserRow;
