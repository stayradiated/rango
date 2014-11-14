'use strict';

var React = require('react');

var EditorMetadata = require('./editorMetadata');
var EditorContent = require('./editorContent');

var Editor = React.createClass({

  propTypes: {
    editor: React.PropTypes.object.isRequired,
  },

  render: function () {
    return (
      <div className='route editor'>
        <EditorMetadata page={this.props.editor.page} />
        <EditorContent page={this.props.editor.page} />
      </div>
    );
  },

});

module.exports = Editor;
