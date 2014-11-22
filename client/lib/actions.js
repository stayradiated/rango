'use strict';

var _ = require('lodash');

_.extend(exports, {
  open: {
    directory: dispatch('OPEN_DIRECTORY'),
    parent: dispatch('OPEN_PARENT_DIRECTORY'),
    page: dispatch('OPEN_PAGE'),
    path: dispatch('OPEN_PATH'),
  },
  create: {
    page: dispatch('CREATE_PAGE'),
    directory: dispatch('CREATE_DIRECTORY'),
  },
  update: {
    page: dispatch('UPDATE_PAGE'),
    content: dispatch('UPDATE_CONTENT'),
    metadata: dispatch('UPDATE_METADATA'),
  },
  select: {
    file: dispatch('SELECT_FILE'),
    none: dispatch('DESELECT_ALL'),
  },
  remove: {
    selected: dispatch('REMOVE_SELECTED_FILES'),
  },
  save: {
    page: dispatch('SAVE_PAGE'),
  },
  publish: {
    site: dispatch('PUBLISH_SITE'),
  },
});

function dispatch (event) {
  return function (args) {
    this.dispatch(event, args);
  };
}
