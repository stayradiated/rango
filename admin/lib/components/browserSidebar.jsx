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
          <hr />
          {
            selected === 1 ? (
              <li>Rename Selected Folder</li>
            ) : null
          }
          {
            selected >= 1 ? (
              <li onClick={this.removeSelected}>Delete Selected Items</li>
            ) : null
          }
        </ul>
      </div>
    );
  },

  createFolder: function () {
    this.getFlux().actions.create.directory();
  },

  createPage: function () {
    this.getFlux().actions.create.page();
  },

  removeSelected: function () {
    this.getFlux().actions.remove.selected();
  },

  /*
  openParent: function () {
    this.getFlux().actions.open.parent();
  },
  */

});

module.exports = BrowserSidebar;
