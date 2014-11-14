'use strict';

var React           = require('react');
var PureRenderMixin = require('react/addons').addons.PureRenderMixin;
var Fluxxor         = require('fluxxor');
var FluxMixin       = Fluxxor.FluxMixin(React);

var Header = React.createClass({
  mixins: [
    FluxMixin,
    PureRenderMixin,
  ],
  
  propTypes: {
    browser: React.PropTypes.object.isRequired,
  },

  render: function () {
    return (
      <header className='app-header'>
        <nav>
          <h1 onClick={this.onPathClick.bind(this, -1)}>Rango</h1>
          {
            this.props.browser.get('path').map(function (name, i) {
              return (
                <span key={i} onClick={this.onPathClick.bind(this, i)}>
                  {name}
                </span>
              );
            }, this).toArray()
          }
        </nav>
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

  onPathClick: function (index) {
    var path = '/';
    if (index >= 0) {
      path = this.props.browser.get('path').slice(0, index + 1).join('/');
    }
    this.getFlux().actions.openPath(path);
  },

});

module.exports = Header;
