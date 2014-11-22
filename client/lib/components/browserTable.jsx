'use strict';

var React           = require('react');
var Fluxxor         = require('fluxxor');
var FluxMixin       = Fluxxor.FluxMixin(React);
var PureRenderMixin = require('react/addons').addons.PureRenderMixin;

var BrowserRow = require('./browserRow');

var BrowserTable = React.createClass({
  mixins: [
    FluxMixin,
    PureRenderMixin,
  ],

  propTypes: {
    browser: React.PropTypes.object.isRequired,
  },

  render: function () {
    var contents = this.props.browser.get('contents');
    var selected = this.props.browser.get('selected');

    var rows = contents.toArray().map(function (item) {
      return (
        <BrowserRow
          item={item}
          key={item.get('name')}
          selected={selected.contains(item)}
        />
      );
    });

    return (
      <div onClick={this.onClick} className='browser-table'>
        <table>
          <thead>
            <tr>
              <th></th>
              <th>Name</th>
              <th>Last Modified</th>
              <th>Contents</th>
              <th>Draft</th>
            </tr>
          </thead>
          <tbody>
            {rows}
          </tbody>
        </table>
      </div>
    );
  },

  onClick: function () {
    this.getFlux().actions.select.none();
  },

});

module.exports = BrowserTable;
