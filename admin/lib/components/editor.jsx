'use strict';

var React           = require('react');
var PureRenderMixin = require('react/addons').addons.PureRenderMixin;

var EditorMetadata = require('./editorMetadata');
var EditorContent  = require('./editorContent');
var EditorImages   = require('./editorImages');

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
        <EditorMetadata metadata={this.props.editor.getIn(['page', 'metadata'])} />
        {/* <EditorContent content={this.props.editor.getIn(['page', 'content'])} */}
        <EditorImages metadata={this.props.editor.getIn(['page', 'metadata'])} />
      </div>
    );
  },

});

module.exports = Editor;
