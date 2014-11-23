'use strict';

var React           = require('react');
var Fluxxor         = require('fluxxor');
var FluxMixin       = Fluxxor.FluxMixin(React);
var PureRenderMixin = require('react/addons').addons.PureRenderMixin;
var CodeMirror      = require('react-code-mirror');
var Marked          = require('marked');

// load markdown syntax for codemirror
require('codemirror/mode/markdown/markdown');

var EditorContent = React.createClass({
  mixins: [
    FluxMixin,
    PureRenderMixin
  ],

  propTypes: {
    content: React.PropTypes.string.isRequired,
  },

  render: function () {
    // var markdown = Marked(this.props.content);

    // <div
    //   className='editor-preview'
    //   dangerouslySetInnerHTML={{__html: markdown}}
    // />

    return (
      <div className='editor-content'>
        <CodeMirror
          className='editor-code'
          value={this.props.content}
          mode='markdown'
          theme='base16-solarized'
          lineNumbers={false}
          lineWrapping={true}
          onChange={this.onChange}
        />
      </div>
    );
  },

  onChange: function (e) {
    var value = e.target.value;
    this.getFlux().actions.update.content(value);
  },

});

module.exports = EditorContent;
