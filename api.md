# Rango API

> This is a brief document outlining the available commands that can be
> accessed through the Rango server.

## Directories

### Creating a Directory

**URL:**

```
POST /api/dir/{path:.*}
```

Where `path` is the path to a directory that the new directory will be created
in.

**Properties:**

- `dir (object)` - a directory to create
    - `name (string0` - the name of the directory

**Example:**

```
POST /api/dir/

{
    "dir": {
        "name": "foo"
    }
}
```

**Response**

```
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

### Renaming a Directory

**URL:**

```
PUT /api/rename/{path:.*}
```

Where `path` is the path to a directory that you want to you rename.

**Properties:**

- `dir (object)` - a directory to create
    - `name (string0` - the new name of the directory

**Example:**

```
POST /api/dir/foo

{
    "dir": {
        "name": "bar"
    }
}
```

**Response**

```
HTTP/1.1 204 No Content
Location: http://localhost:8080/api/dir/bar
Content-Type: application/json
```

```
HTTP/1.1 404 Not Found
Location: http://localhost:8080/api/dir/abc
Content-Type: application/json

{
    "errors": {
    }
}
```

## Pages

## Commands

### Copy a Page or Directory

### Move a Page or Directory

### Delete a Page or Directory
