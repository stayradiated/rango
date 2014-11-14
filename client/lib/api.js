var jQuery = require('jquery');
var Path = require('path');

var BASE_URL = 'api/'
var DIR = 'dir';
var FILE = 'file';
var PAGE = 'page';
var CONFIG = 'config';

var Rango = {

  createDir: function (path) {
    return post(DIR, path);
  },

  createFile: function (path) {
    return post(FILE, path);
  },

  createPage: function (dirPath, page) {
    return post(PAGE, dirPath, {
      metadata: JSON.stringify(page.metadata),
      content: page.content,
    });
  },

  readDir: function (path) {
    return get(DIR, path);
  },

  readFile: function (path) {
    return get(FILE, path);
  },
  
  readPage: function (path) {
    return get(PAGE, path);
  },

  readConfig: function () {
    return get(CONFIG, '');
  },

  updateFile: function (path, contents) {
    return put(FILE, path, { data: contents });
  },

  updatePage: function (path, page) {
    return put(PAGE, path, { page: page });
  },

  updateConfig: function (config) {
    return put(CONFIG, '', { config: config });
  },

  copy: function (src, dst) {
    return post('copy', src, { destination: dst });
  },

  move: function (src, dst) {
    return post('copy', src, { destination: dst }).then(function () {
      return del(src);
    });
  },

  rename: function (path, name) {
    return put('rename', path, { name: name });
  },

  destroy: function (path) {
    return del(path);
  },

};

function get (method, path) {
  return jQuery.ajax({
    type: 'get',
    url: Path.join(BASE_URL, method, path),
    dataType: 'json',
  });
}

function post (method, path, data) {
  return jQuery.ajax({
    type: 'post',
    url: Path.join(BASE_URL, method, path),
    dataType: 'json',
    data: data,
  });
}

function put (method, path, params) {
  return jQuery.ajax({
    type: 'put',
    url: Path.join(BASE_URL, method, path),
    dataType: 'json',
    data: data,
  });
}

function del (path) {
  return jQuery.ajax({
    type: 'delete',
    url: Path.join(BASE_URL, path),
    dataType: 'json',
  });
}

module.exports = Rango;
