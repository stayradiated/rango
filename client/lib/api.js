'use strict';

var _ = require('lodash');
var jQuery = require('jquery');
var Path = require('path');

var BASE_URL = 'api/'
var DIR = 'dir';
var PAGE = 'page';
var CONFIG = 'config';

_.extend(exports, {

  // Directories

  readDir: function (path) {
    return get(DIR, path).then(function (res) {
      return res.data;
    });
  },

  createDir: function (path, dirname) {
    return post(DIR, path, {
      dir: {
        name: dirname,
      },
    }).then(function (res) {
      return res.dir;
    });
  },

  updateDir: function (path, props) {
    return put(DIR, path, {
      dir: props,
    }).then(function (res) {
      return res.dir;
    });
  },

  deleteDir: function (path) {
    return del(path);
  },


  // Pages

  readPage: function (path) {
    return get(PAGE, path).then(function (res) {
      return res.page;
    });
  },

  createPage: function (dirPath, page) {
    return post(PAGE, dirPath, {
      page: {
        meta: JSON.stringify(page.metadata),
        content: page.content,
      },
    }).then(function (res) {
      return res.page;
    });
  },

  updatePage: function (path, page) {
    return put(PAGE, path, {
      page: {
        meta: JSON.stringify(page.metadata),
        content: page.content,
      },
    }).then(function (res) {
      return res.page;
    });
  },

  destroy: function (path) {
    return del(path);
  },

  
  // Config

  readConfig: function () {
    return get(CONFIG, '');
  },

  updateConfig: function (config) {
    return put(CONFIG, '', { config: config });
  },


  // Files

  copy: function (src, dst) {
    return post('copy', src, { destination: dst });
  },

  move: function (src, dst) {
    return post('copy', src, { destination: dst }).then(function () {
      return del(src);
    });
  },

});

function get (method, path) {
  return jQuery.ajax({
    type: 'get',
    url: Path.join(BASE_URL, method, path),
    dataType: 'json',
  });
}

function post (method, path, params) {
  return jQuery.ajax({
    type: 'post',
    url: Path.join(BASE_URL, method, path),
    dataType: 'json',
    data: params,
  });
}

function put (method, path, params) {
  return jQuery.ajax({
    type: 'put',
    url: Path.join(BASE_URL, method, path),
    dataType: 'json',
    data: params,
  });
}

function del (path) {
  return jQuery.ajax({
    type: 'delete',
    url: Path.join(BASE_URL, path),
    dataType: 'json',
  });
}
