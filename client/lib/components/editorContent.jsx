'use strict';

var React = require('react');

var EditorContent = React.createClass({

  propTypes: {
  },

  render: function () {
    return (
      <div className='editor-content'>
        <div className='textarea'>
          <textarea defaultValue='hello world' />
        </div>
      </div>
    );
  },

});

module.exports = EditorContent;
