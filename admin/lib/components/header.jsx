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
            onClick={this.onClickSaveBtn}>Save Page</button>
          <button
            className='button button-primary'
            onClick={this.onClickPublishBtn}>Publish Site</button>
        </div>
      </header>
    );
  },

  onClickSaveBtn: function () {
    this.getFlux().actions.save.page();
  },

  onClickPublishBtn: function () {
    this.getFlux().actions.publish.site();
  },

  onPathClick: function (index) {
    var path = '/';
    if (index >= 0) {
      path = this.props.browser.get('path').slice(0, index + 1).join('/');
    }
    this.getFlux().actions.open.path(path);
  },

});

module.exports = Header;
