'use strict';

var React = require('react');

var BrowserSidebar = React.createClass({

  propTypes: {
  },

  render: function () {
    return (
      <div className='browser-sidebar'>
        <ul>
          <li>Create Folder</li>
          <li>Create Page</li>
          <li>Rename Folder</li>
          <li>Delete Folder</li>
        </ul>
      </div>
    );
  },

});

module.exports = BrowserSidebar;
