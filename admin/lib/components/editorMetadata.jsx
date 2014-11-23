'use strict';

var React           = require('react');
var PureRenderMixin = require('react/addons').addons.PureRenderMixin;
var Fluxxor         = require('fluxxor');
var FluxMixin       = Fluxxor.FluxMixin(React);

var InputText     = require('./inputText');
var InputTextarea = require('./inputTextarea');
var InputCheckbox = require('./inputCheckbox');

var EditorMetadata = React.createClass({
  mixins: [
    FluxMixin,
    PureRenderMixin,
  ],

  propTypes: {
    metadata: React.PropTypes.object.isRequired,
  },

  render: function () {
    var meta = this.props.metadata;

    var data = {
      title: meta.get('title', ''),
      description: meta.get('description', ''),
      date: meta.get('date', ''),
      type: meta.get('type', ''),
    };

    // for (var key in meta) {
    //   data[key] = meta[key];
    // }

    var inputs = [];
    for (var key in data) {
      inputs.push(
        <InputText
          key={key}
          label={key}
          value={data[key]}
          onChange={this.onChange.bind(this, key)}
        />
      );
    }

    return (
      <div className='editor-metadata'>
        <h2>Details</h2>
        {inputs}
        <InputCheckbox label='Draft' />
      </div>
    );
  },

  onChange: function (key, e) {
    var value = e.target.value;
    var meta = this.props.metadata.set(key, value);
    this.getFlux().actions.update.metadata(meta);
  },

});

module.exports = EditorMetadata;
