'use strict';

var React = require('react');

var BrowserSidebar = require('./browserSidebar');
var BrowserTable = require('./browserTable');

var Browser = React.createClass({

  propTypes: {
  },

  render: function () {
    return (
      <div className='browser'>
        <BrowserSidebar />
        <BrowserTable />
      </div>
    );
  },

});

module.exports = Browser;
