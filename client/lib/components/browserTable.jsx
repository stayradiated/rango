'use strict';

var React = require('react');

var BrowserRow = require('./browserRow');

var BrowserTable = React.createClass({

  propTypes: {
  },

  render: function () {
    return (
      <div className='browser-table'>
        <table>
          <thead>
            <tr>
              <th></th>
              <th>Name</th>
              <th>Last Modified</th>
              <th>Contents</th>
              <th>Draft</th>
            </tr>
          </thead>
          <tbody>
            <BrowserRow type='folder' name='Gallery' lastModified='27th Sep 2014' />
            <BrowserRow type='folder' name='Blog' lastModified='1st Oct 2014' />
            <BrowserRow type='page' name='About' lastModified='9th Sep 2014' />
            <BrowserRow type='page' name='Bio' lastModified='7th Sep 2014' />
            <BrowserRow type='page' name='Contact' lastModified='1st Sep 2014' />
          </tbody>
        </table>
      </div>
    );
  },

});

module.exports = BrowserTable;
