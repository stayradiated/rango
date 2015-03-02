'use strict';

/**
 * TODO: Split this file up into multiple components.
 */

var React           = require('react');
var PureRenderMixin = require('react/addons').addons.PureRenderMixin;
var Sortable        = require('react-anything-sortable');
var Immutable       = require('immutable');
var Fluxxor         = require('fluxxor');
var FluxMixin       = Fluxxor.FluxMixin(React);

var ImageItem    = require('./imageItem');
var SortableItem = require('./sortableItem');

var EditorImages = React.createClass({

  mixins: [
    FluxMixin,
    PureRenderMixin,
  ],

  propTypes: {
    // images: React.PropTypes.arrayOf(React.PropTypes.string).isRequired,
    metadata: React.PropTypes.object.isRequired,
  },

  render: function () {
    var meta = this.props.metadata;
    var key = 'images';

    var imageArr = meta.get('images');
    var images = [];
    var key = '';

    if (imageArr != null) {
      // concat image paths to make parent key
      key = imageArr.toArray().join('+');

      images = imageArr.map(function (path, index) {
        return (
          <SortableItem className='image' key={path} sortData={path}>
          <ImageItem
            key={path}
            path={path}
            onRemove={this.handleRemove.bind(null, path)}
          />
          </SortableItem>
        );
      }, this).toArray();
    }

    return (
      <div className='editor-images'>
        <div className='images'>
          <Sortable key={key} onSort={this.handleSort}>
            {images}
          </Sortable>
        </div>
      </div>
    );
  },

  handleSort: function (imageArray) {
    var imageList = Immutable.fromJS(imageArray);
    var meta = this.props.metadata.set('images', imageList);
    this.getFlux().actions.update.metadata(meta);
  },

  handleRemove: function (path) {
    var meta = this.props.metadata;

    if (meta.has('images')) {
      meta = meta.update('images', function (images) {
        var index = images.indexOf(path);
        if (index >= 0) {
          images = images.delete(index);
        }
        return images;
      });

      this.getFlux().actions.update.metadata(meta);
    }
  },

});

module.exports = EditorImages;
