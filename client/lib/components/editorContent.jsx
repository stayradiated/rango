'use strict';

var React = require('react');

var EditorContent = React.createClass({

  propTypes: {
    page: React.PropTypes.object.isRequired,
  },

  render: function () {
    return (
      <div className='editor-content'>
        <div className='textarea'>
          <textarea value={this.props.page.content} readOnly />
        </div>
      </div>
    );
  },

});

module.exports = EditorContent;
