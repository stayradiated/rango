'use strict';

var React     = require('react');
var classSet  = require('react/addons').addons.classSet;
var Fluxxor   = require('fluxxor');
var FluxMixin = Fluxxor.FluxMixin(React);

var BrowserRow = React.createClass({
  mixins: [FluxMixin],

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
        <td className='name'>{item.name}</td>
        <td className='last-modified'></td>
        <td className='contents'></td>
        <td className='draft'></td>
      </tr>
    );
  },

  onClick: function (e) {
    e.stopPropagation();
    this.getFlux().actions.selectFile(this.props.item.name);
  },

  onDoubleClick: function () {
    this.getFlux().actions.openPath(this.props.item.path);
  },

});

module.exports = BrowserRow;
