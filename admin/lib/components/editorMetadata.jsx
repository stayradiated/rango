'use strict';

var React           = require('react');
var PureRenderMixin = require('react/addons').addons.PureRenderMixin;
var Fluxxor         = require('fluxxor');
var FluxMixin       = Fluxxor.FluxMixin(React);

var InputText     = require('./inputText');
var InputTextarea = require('./inputTextarea');
var InputCheckbox = require('./inputCheckbox');
var InputDropdown = require('./inputDropdown');

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
      title: {
        type: InputText,
        value: meta.get('title', ''),
      },
      part: {
        type: InputText,
        value: meta.get('part', ''),
      },
      description: {
        type: InputText,
        value: meta.get('description', ''),
      },
      date: {
        type: InputText,
        value: meta.get('date', ''),
      },
      type: {
        type: InputDropdown,
        value: meta.get('type', ''),
      },
    };

    // for (var key in meta) {
    //   data[key] = meta[key];
    // }

    var inputs = [];
    for (var key in data) {
      inputs.push(
        React.createElement(data[key].type, {
          key: key,
          label: key,
          value: data[key].value,
          onChange: this.onChange.bind(this, key),
        })
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
