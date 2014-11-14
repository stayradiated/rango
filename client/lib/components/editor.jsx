'use strict';

var React           = require('react');
var PureRenderMixin = require('react/addons').addons.PureRenderMixin;

var EditorMetadata = require('./editorMetadata');
var EditorContent  = require('./editorContent');

var Editor = React.createClass({
  mixins: [
    PureRenderMixin,
  ],

  propTypes: {
    editor: React.PropTypes.object.isRequired,
  },

  render: function () {
    return (
      <div className='route editor'>
        <EditorMetadata page={this.props.editor.get('page')} />
        <EditorContent page={this.props.editor.get('page')} />
      </div>
    );
  },

});

module.exports = Editor;
