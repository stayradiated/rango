'use strict';

var React = require('react');
var Fluxxor   = require('fluxxor');
var FluxMixin = Fluxxor.FluxMixin(React);

var BrowserDirectory = require('./browserDirectory');
var BrowserFile = require('./browserFile');

var BrowserTable = React.createClass({
  mixins: [FluxMixin],

  propTypes: {
    browser: React.PropTypes.object.isRequired,
  },

  render: function () {
    var rows = [];
    var contents = this.props.browser.contents;

    if (! this.props.browser.isRoot) {
      rows.push(
        <BrowserDirectory
          key='..'
          item={{name: '..'}}
          onClick={this.onOpenParentClick}
        />
      );
    };

    for (var i = 0, len = contents.directories.length; i < len; i++) {
      var dir = contents.directories[i];
      rows.push(
        <BrowserDirectory
          key={dir.name}
          item={dir}
          onClick={this.onDirectoryClick.bind(null, dir)}
        />
      );
    }

    for (var i = 0, len = contents.pages.length; i < len; i++) {
      var page = contents.pages[i];
      rows.push(
        <BrowserFile
          key={page.name}
          item={page}
          onClick={this.onFileClick.bind(null, page)}
        />
      );
    }

    return (
      <div className='browser-table'>
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

  onFileClick: function (file) {
    this.getFlux().actions.openPage(file.path);
  },

  onDirectoryClick: function (dir) {
    this.getFlux().actions.openDirectory(dir.path);
  },

  onOpenParentClick: function () {
    this.getFlux().actions.openParentDirectory();
  },

});

module.exports = BrowserTable;
