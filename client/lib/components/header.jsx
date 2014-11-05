'use strict';

var React = require('react');

var Header = React.createClass({
  
  propTypes: {
  },

  render: function () {
    return (
      <header className='app-header'>
        <h1>Rango > About</h1>
        <div className='button-group'>
          <button
            className='button'
            onClick={this.onClickSettingsBtn}>Settings</button>
          <button
            className='button button-primary'
            Click={this.onClickPublishBtn}>Publish</button>
        </div>
      </header>
    );
  },

  onClickSettingsBtn: function () {
  },

  onClickPublishBtn: function () {
  },

});

module.exports = Header;
