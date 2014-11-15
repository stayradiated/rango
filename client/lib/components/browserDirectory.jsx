'use strict';

var React           = require('react');
var classSet        = require('react/addons').addons.classSet;
var PureRenderMixin = require('react/addons').addons.PureRenderMixin;
var Fluxxor         = require('fluxxor');
var FluxMixin       = Fluxxor.FluxMixin(React);

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

    var classes = classSet({
      selected: this.props.selected,
    });

    return (
      <tr
        className={classes}
        onClick={this.onClick}
        onDoubleClick={this.onDoubleClick}
      >
        <td className='type'>📁</td>
        <td className='name'>{item.get('name')}</td>
        <td className='last-modified'></td>
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
    this.getFlux().actions.open.path(this.props.item.get('path'));
  },

});

module.exports = BrowserRow;
