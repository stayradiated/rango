'use strict';

var React     = require('react');
var Fluxxor   = require('fluxxor');
var FluxMixin = Fluxxor.FluxMixin(React);

var Constants        = require('../constants');
var BrowserDirectory = require('./browserDirectory');
var BrowserFile      = require('./browserFile');

var BrowserTable = React.createClass({
  mixins: [FluxMixin],

  propTypes: {
    browser: React.PropTypes.object.isRequired,
  },

  render: function () {
    var rows = [];
    var contents = this.props.browser.contents;
    var selected = this.props.browser.selected;

    for (var i = 0, len = contents.directories.length; i < len; i++) {
      var dir = contents.directories[i];
      rows.push(
        <BrowserDirectory
          key={dir.name}
          item={dir}
          selected={selected.contains(dir.name)}
        />
      );
    }

    for (var i = 0, len = contents.pages.length; i < len; i++) {
      var page = contents.pages[i];
      rows.push(
        <BrowserFile
          key={page.name}
          item={page}
          selected={selected.contains(page.name)}
        />
      );
    }

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
