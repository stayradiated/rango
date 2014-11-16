# Rango API

> This is a brief document outlining the available commands that can be
> accessed through the Rango server.

## Directories

### Reading the contents of a directory

**URL**

``` http
GET /api/dir/{path:.*}
```

Where `path` is the path to an existing directory that you want to get the
contents of.

**Example Request**

``` http
GET /api/dir/foo/bar
```

**Example Response**

``` http
HTTP/1.1 200 OK

{
    "dir": {
        "name": "bar",
        "path": "/foo/bar",
        "type": "directory",
        "size": 145,
        "atime": 1416092956000,
        "mtime": 1416092956000,
        "link": false,
        "contents": [{
            "name": "abc",
            "path": "/foo/bar/abc",
            "type": "directory",
            "size": 136,
            "atime": 1416092956000,
            "mtime": 1416092956000,
            "link": false
        }, {
            "name": "xyz.md",
            "path": "/foo/bar/xyz.md",
            "type": "file",
            "size": 9,
            "atime": 1416092648000,
            "mtime": 1416092170000,
            "link": false
        }]
    }
}
```

**Errors**

```
HTTP/1.1 404 Not Found

{
    "errors": {
        "status": "404",
        "code": "dir_not_found",
        "title": "directory does not exist",
        "detail": "Could not find the directory '/foo/bar'"
    }
}
```

### Creating a directory

**URL**

``` http
POST /api/dir/{path:.*}
```

Where `path` is the path to a directory that the new directory will be created
in.

**Properties**

- `dir (object)`
    - `name (string)` - the name of the directory

**Example Request**

``` http
POST /api/dir/

{
    "dir": {
        "name": "foo"
    }
}
```

**Example Response**

``` http
HTTP/1.1 201 Created

{
    "dir": {
        "name": "foo",
        "path": "/foo",
        "type": "directory",
        "size": 10,
        "atime": 1416092956000,
        "mtime": 1416092956000,
        "link": false
    }
}
```

**Errors**

``` http
HTTP/1.1 400 Bad Request

{
    "errors": {
        "status": "400",
        "code": "malformed_json",
        "title": "could not parse request body",
        "detail": "Could not parse request body as JSON"
    }
}
```

``` http
HTTP/1.1 409 Conflict

{
    "errors": {
        "status": "409",
        "code": "dir_conflict",
        "title": "directory already exists",
        "detail": "Could not create '/foo/bar' as it already exists"
    }
}
```


### Renaming a Directory

**URL**

``` http
PUT /api/rename/{path:.*}
```

Where `path` is the path to a directory that you want to you rename.

**Properties**

- `dir (object)`
    - `name (string)` - the new name of the directory

**Example Request**

``` http
PUT /api/dir/foo

{
    "dir": {
        "name": "bar"
    }
}
```

**Example Response**

``` http
HTTP/1.1 200 OK

{
    "name": "bar",
    "path": "/bar",
    "type": "directory",
    "size": 10,
    "atime": 1416092956000,
    "mtime": 1416092956000,
    "link": false
}
```

**Errors**

``` http
HTTP/1.1 400 Bad Request

{
    "errors": {
        "status": "400",
        "code": "malformed_json",
        "title": "could not parse request body",
        "detail": "Could not parse request body as JSON"
    }
}
```

```
HTTP/1.1 404 Not Found

{
    "errors": {
        "status": "404",
        "code": "dir_not_found",
        "title": "directory does not exist",
        "detail": "Could not find the directory '/foo/bar'"
    }
}
```

``` http
HTTP/1.1 409 Conflict

{
    "errors": {
        "status": "409",
        "code": "dir_conflict",
        "title": "directory already exists",
        "detail": "Could not create '/foo/bar' as it already exists"
    }
}
```

## Pages

## Read a page

**URL**

``` http
GET /api/page/{path:.*}
```

**Example Request**

``` http
GET /api/page/foo/bar/lorem-ipsum.md
```

**Example Response**

``` http
HTTP/1.1 200 OK

{
    "page": {
        "path": "/foo/bar/lorem-ipsum.md",
        "meta": {
            "title": "Lorem Ipsum",
            "description": "bla blah bla"
        },
        "content": "# Hello World"
    }
}
```

**Errors**

```
HTTP/1.1 404 Not Found

{
    "errors": {
        "status": "404",
        "code": "file_not_found",
        "title": "file does not exist",
        "detail": "Could not find the file '/foo/bar/lorem-ipsum.md'"
    }
}
```

### Create a page

**URL**

``` http
POST /api/page/{path:.*}
```

Where `path` is the path to a directory where the new page will be created in.

**Properties**

- `page (object)`
    - `meta (object)` - page metadata
        - `title (string)` - the name of the page
        - `...`
    - `content (string)` - page content in markdown

**Example Request**

``` http
POST /api/page/foo

{
    "page": {
        "meta": {
            "title": "Lorem ipsum"
        },
        "content": "# Hello World"
    }
}
```

**Example Response**

A filename will be created based on the title of the page. This will be in
lowercase, and all spaces will be replaced with hypens and have the `.md`
extension appended to the end. E.g. `Lorem Ipsum -> lorem-ipsum.md`

If a file already exists with that filename, a number will be appended to it.
E.g. `lorem-ipsum.md -> lorem-ipsum-1.md -> lorem-ipsum-2.md`

``` http
HTTP/1.1 201 Created

{
    "page": {
        "path": "/foo/lorem-ipsum.md",
        "meta": {
            "title": "Lorem ipsum"
        },
        "content": "# Hello World"
    }
}
```

**Errors**

``` http
HTTP/1.1 400 Bad Request

{
    "errors": {
        "status": "400",
        "code": "malformed_json",
        "title": "could not parse request body",
        "detail": "Could not parse request body as JSON"
    }
}
```

```
HTTP/1.1 404 Not Found

{
    "errors": {
        "status": "404",
        "code": "dir_not_found",
        "title": "directory does not exist",
        "detail": "Could not find the directory '/foo/bar'"
    }
}
```

### Update a Page

**URL**

``` http
PUT /api/page/{path:.*}
```

Where `path` is the path to an existing page that you want to update.

**Properties**

- `page (object)`
    - `meta (object)` - page metadata
        - `title (string)` - the name of the page
        - `...`
    - `content (string)` - page content in markdown

**Example Request**

If the title of the page is changed, the file will be renamed. See "Create a
page" for more details on how the filename is generated.

``` http
PUT /api/page/foo/lorem-ipsum.md

{
    "page": {
        "meta": {
            "title": "Programming News"
        }
    }
}
```

**Example Response**

If the title didn't change:

``` http
HTTP/1.1 204 No Content
```

If the title was changed:

``` http
HTTP/1.1 200 OK

{
    "page": {
        "path": "/foo/programming-news.md",
        "meta": {
            "title": "Programming News"
        },
        "content": "# Hello World"
    }
}
```

**Errors**

``` http
HTTP/1.1 400 Bad Request

{
    "errors": {
        "status": "400",
        "code": "malformed_json",
        "title": "could not parse request body",
        "detail": "Could not parse request body as JSON"
    }
}
```


```
HTTP/1.1 404 Not Found

{
    "errors": {
        "status": "404",
        "code": "file_not_found",
        "title": "file does not exist",
        "detail": "Could not find the file '/foo/bar/lorem-ipsum.md'"
    }
}
```


## Commands

### Copy a file or directory

**URL**

``` http
POST /api/copy/{path:.*}
```

Where `path` is the path to an existing directory or file you want to copy.

**Properties**

- `destination (string)` - path to where the item should be copied.

**Example Request**

``` http
POST /api/copy/foo/bar

{
    "destination": "/baz"
}
```

**Example Response**

```
HTTP/1.1 201 Created

{
    "data": {
        "name": "baz",
        "path": "/baz",
        "type": "directory",
        "size": 136,
        "atime": 1416092956000,
        "mtime": 1416092956000,
        "link": false
    }
}
```

**Errors**

``` http
HTTP/1.1 400 Bad Request

{
    "errors": {
        "status": "400",
        "code": "malformed_json",
        "title": "could not parse request body",
        "detail": "Could not parse request body as JSON"
    }
}
```

```
HTTP/1.1 404 Not Found

{
    "errors": {
        "status": "404",
        "code": "file_not_found",
        "title": "directory does not exist",
        "detail": "Could not find the file '/foo/bar'"
    }
}
```

``` http
HTTP/1.1 409 Conflict

{
    "errors": {
        "status": "409",
        "code": "dir_conflict",
        "title": "directory already exists",
        "detail": "Could not create '/foo/bar' as it already exists"
    }
}
```

**Errors**

### Delete a file or directory

**URL**

``` http
DELETE /api/{path:.*}
```

Where `path` is the path to a directory or page that you want to you delete.

**Example Request**

``` http
DELETE /api/foo/bar.md
```

**Example Response**

``` http
HTTP/1.1 204 No Content
```

**Errors**

```
HTTP/1.1 404 Not Found

{
    "errors": {
        "status": "404",
        "code": "file_not_found",
        "title": "file does not exist",
        "detail": "Could not find the file '/foo/bar'"
    }
}
```
