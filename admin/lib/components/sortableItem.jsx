'use strict';

var React           = require('react');
var PureRenderMixin = require('react/addons').addons.PureRenderMixin;
var SortableItemMixin = require('react-anything-sortable/SortableItemMixin');

var ImageItem = React.createClass({
  mixins: [
    SortableItemMixin,
    PureRenderMixin,
  ],

  getDefaultProps: function () {
    return {
      className: '',
    }
  },

  render: function () {
    return this.renderWithSortable(
      <div className={this.props.className}>
        {this.props.children}
      </div>
    );
  },

});

module.exports = ImageItem;
