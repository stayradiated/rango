# Rango API

> This is a brief document outlining the available commands that can be
> accessed through the Rango server.

## Directories

### Creating a Directory

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
Content-Type: application/json
Accept: application/json

{
    "dir": {
        "name": "foo"
    }
}
```

**Example Response**

``` http
HTTP/1.1 201 Created
Location: http://localhost:8080/api/dir/foo
Content-Type: application/json

{
    "dir": {
        name: "foo",
        path: "/foo"
    }
}
```

**Errors**

```
HTTP/1.1 409 Conflict
Content-Type: application/json

{
    "errors": {
        // a directory with that name already exists
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
Content-Type: application/json
Accept: application/json

{
    "dir": {
        "name": "bar"
    }
}
```

**Example Response**

``` http
HTTP/1.1 204 No Content
```

**Errors**

``` http
HTTP/1.1 404 Not Found
Content-Type: application/json

{
    "errors": {
        // couldn't find directory specified by path
    }
}
```

``` http
HTTP/1.1 409 Conflict
Content-Type: application/json

{
    "errors": {
        // a directory with that name already exists
    }
}
```

## Pages

## Commands

### Copy a Page or Directory

### Delete a Page or Directory

**URL**

``` http
DELETE /api/{path:.*}
```

Where `path` is the path to a directory or page that you want to you delete.

**Example Request**

``` http
DELETE /api/dir/foo/bar.md
```

**Example Response**

``` http
HTTP/1.1 204 No Content
```

**Errors**

``` http
HTTP/1.1 404 Not Found
Content-Type: application/json

{
    "errors": {
        // couldn't find directory specified by path
    }
}
```
