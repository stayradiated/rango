'use strict';

var React = require('react');

var InputText = require('./inputText');
var InputCheckbox = require('./inputCheckbox');

var EditorMetadata = React.createClass({

  propTypes: {
  },

  render: function () {
    return (
      <div className='editor-metadata'>
        <h2>Details</h2>
        <InputText label='Title' value='About' />
        <InputText label='Description' value='The Rango editor is quite ...' />
        <InputText label='Date' value='2nd Oct 2014' />
        <InputText label='Slug' value='about' />
        <InputText label='Type' value='Article' />
        <InputCheckbox label='Draft' />
      </div>
    );
  },

});

module.exports = EditorMetadata;
