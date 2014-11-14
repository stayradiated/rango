'use strict';

var React           = require('react');
var PureRenderMixin = require('react/addons').addons.PureRenderMixin;

var InputText     = require('./inputText');
var InputTextarea = require('./inputTextarea');
var InputCheckbox = require('./inputCheckbox');

var EditorMetadata = React.createClass({
  mixins: [
    PureRenderMixin,
  ],

  propTypes: {
    page: React.PropTypes.object.isRequired,
  },

  render: function () {
    var meta = this.props.page.get('metadata');

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
        <InputText key={key} label={key} value={data[key]} />
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

});

module.exports = EditorMetadata;
