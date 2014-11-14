'use strict';

var React           = require('react');
var Fluxxor         = require('fluxxor');
var FluxMixin       = Fluxxor.FluxMixin(React);
var PureRenderMixin = require('react/addons').addons.PureRenderMixin;

var BrowserSidebar = React.createClass({
  mixins: [
    FluxMixin,
    PureRenderMixin,
  ],

  propTypes: {
    browser: React.PropTypes.object.isRequired,
  },

  render: function () {
    var selected = this.props.browser.get('selected').size;

    return (
      <div className='browser-sidebar'>
        <ul>
          <li onClick={this.createFolder}>Create Folder</li>
          <li onClick={this.createPage}>Create Page</li>
          {
            selected === 1 ? (
              <li>Rename Selected Folder</li>
            ) : null
          }
          {
            selected === 1 ? (
              <li onClick={this.removeSelected}>Delete Selected Files</li>
            ) : null
          }
        </ul>
      </div>
    );
  },

  createFolder: function () {
    this.getFlux().actions.createDirectory();
  },

  createPage: function () {
    this.getFlux().actions.createPage();
  },

  removeSelected: function () {
    this.getFlux().actions.removeSelectedFiles();
  },

  /*
  openParent: function () {
    this.getFlux().actions.openParentDirectory();
  },
  */

});

module.exports = BrowserSidebar;
