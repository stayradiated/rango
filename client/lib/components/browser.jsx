'use strict';

var React = require('react');

var BrowserSidebar = require('./browserSidebar');
var BrowserTable = require('./browserTable');

var Browser = React.createClass({

  propTypes: {
    browser: React.PropTypes.object.isRequired,
  },

  render: function () {
    return (
      <div className='browser'>
        <BrowserSidebar />
        <BrowserTable browser={this.props.browser} />
      </div>
    );
  },

});

module.exports = Browser;
