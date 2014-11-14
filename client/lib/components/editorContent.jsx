'use strict';

var React           = require('react');
var PureRenderMixin = require('react/addons').addons.PureRenderMixin;
var CodeMirror      = require('react-code-mirror');
var Marked          = require('marked');

// load markdown syntax for codemirror
require('codemirror/mode/markdown/markdown');

var EditorContent = React.createClass({
  mixins: [
    PureRenderMixin
  ],

  propTypes: {
    page: React.PropTypes.object.isRequired,
  },

  getInitialState: function (props) {
    return {
      content: (props || this.props).page.get('content'),
    };
  },

  componentWillReceiveProps: function (nextProps) {
    this.setState(this.getInitialState(nextProps));
  },

  render: function () {
    var markdown = Marked(this.state.content);

    return (
      <div className='editor-content'>
        <CodeMirror
          className='editor-code'
          value={this.state.content}
          mode='markdown'
          theme='base16-solarized'
          lineNumbers={false}
          lineWrapping={true}
          onChange={this.onChange}
        />
        <div
          className='editor-preview'
          dangerouslySetInnerHTML={{__html: markdown}}
        />
      </div>
    );
  },

  onChange: function (e) {
    this.setState({
      content: e.target.value
    });
  },

});

module.exports = EditorContent;
