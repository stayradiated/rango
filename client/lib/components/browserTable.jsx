'use strict';

var React           = require('react');
var Fluxxor         = require('fluxxor');
var FluxMixin       = Fluxxor.FluxMixin(React);
var PureRenderMixin = require('react/addons').addons.PureRenderMixin;

var Constants        = require('../constants');
var BrowserDirectory = require('./browserDirectory');
var BrowserFile      = require('./browserFile');

var BrowserTable = React.createClass({
  mixins: [
    FluxMixin,
    PureRenderMixin,
  ],

  propTypes: {
    browser: React.PropTypes.object.isRequired,
  },

  render: function () {
    var rows = [];
    var contents = this.props.browser.get('contents');
    var selected = this.props.browser.get('selected');

    contents.get('directories').forEach(function (item) {
      rows.push(
        <BrowserDirectory
          key={item.get('name')}
          item={item}
          selected={selected.contains(item)}
        />
      );
    });

    contents.get('pages').forEach(function (item) {
      rows.push(
        <BrowserFile
          key={item.get('name')}
          item={item}
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
    this.getFlux().actions.deselectAll();
  },

});

module.exports = BrowserTable;
