'use strict'; 

var Constants = {

  OPEN_DIRECTORY:         null,
  OPEN_PARENT_DIRECTORY:  null,
  OPEN_PATH:              null,
  OPEN_PAGE:              null,

  CREATE_PAGE:            null,
  CREATE_DIRECTORY:       null,

  SELECT_FILE:            null,
  DESELECT_ALL:           null,

  REMOVE_SELECTED_FILES:  null,

  SAVE_PAGE:              null,

  ROUTE_BROWSER:          null,
  ROUTE_EDITOR:           null,

};

for (var key in Constants) {
  exports[key] = Constants[key] ? Constants[key] : key;
}
