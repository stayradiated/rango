'use strict';

var React = require('react');
var Fluxxor   = require('fluxxor');
var FluxMixin = Fluxxor.FluxMixin(React);

var Header = React.createClass({
  mixins: [FluxMixin],
  
  propTypes: {
    browser: React.PropTypes.object.isRequired,
  },

  render: function () {
    var title = this.props.browser.path.join(' > ');

    return (
      <header className='app-header'>
        <h1><a href='#' onClick={this.onPathClick}>Rango</a> > {title}</h1>
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

  onPathClick: function () {
    this.getFlux().actions.openPath([]);
  },

});

module.exports = Header;
