'use strict';

var React           = require('react');
var PureRenderMixin = require('react/addons').addons.PureRenderMixin;

var BrowserSidebar = require('./browserSidebar');
var BrowserTable   = require('./browserTable');

var Browser = React.createClass({
  mixins: [PureRenderMixin],

  propTypes: {
    browser: React.PropTypes.object.isRequired,
  },

  render: function () {
    return (
      <div className='route browser'>
        <BrowserSidebar browser={this.props.browser} />
        <BrowserTable browser={this.props.browser} />
      </div>
    );
  },

});

module.exports = Browser;
