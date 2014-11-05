'use strict';

var React = require('react');

var EditorMetadata = require('./editorMetadata');
var EditorContent = require('./editorContent');

var Editor = React.createClass({

  propTypes: {
  },

  render: function () {
    return (
      <div className='editor'>
        <EditorMetadata />
        <EditorContent />
      </div>
    );
  },

});

module.exports = Editor;
