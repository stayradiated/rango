'use strict';

var React = require('react');

var InputText = require('./inputText');
var InputTextarea = require('./inputTextarea');
var InputCheckbox = require('./inputCheckbox');

var EditorMetadata = React.createClass({

  propTypes: {
    page: React.PropTypes.object.isRequired,
  },

  render: function () {
    var meta = this.props.page.metadata;

    var data = {
      title: meta.title || '',
      description: meta.description || '',
      date: meta.date || '',
      type: meta.type || '',
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
