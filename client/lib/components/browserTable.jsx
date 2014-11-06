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
      var dirItem = contents.directories[i];
      rows.push(
        <BrowserDirectory
          key={dirItem.name}
          item={dirItem}
          onClick={this.onDirectoryClick.bind(null, dirItem.name)}
        />
      );
    }

    for (var i = 0, len = contents.files.length; i < len; i++) {
      var fileItem = contents.files[i];
      rows.push(
        <BrowserFile
          key={fileItem.name}
          item={fileItem}
          onClick={this.onFileClick.bind(null, fileItem.name)}
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

  onFileClick: function (name) {
    this.getFlux().actions.openFile(name);
  },

  onDirectoryClick: function (name) {
    this.getFlux().actions.openDirectory(name);
  },

  onOpenParentClick: function () {
    this.getFlux().actions.openParentDirectory();
  },

});

module.exports = BrowserTable;
