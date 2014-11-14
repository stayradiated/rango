rango
=====

A web frontend for [hugo](https://gohugo.io).

It's designed to make it easy to manage a small site, even for people with
little computer experience.

## Installation

```
$ go get -u -v github.com/stayradiated/rango
$ cd $GOPATH/src/github.com/stayradiated/rango
$ cd client
$ npm install
$ gulp
$ cd ..
$ go build
$ ./rango
```

## Using with Apache

Based on [this
tutorial](http://www.jeffreybolle.com/blog/run-google-go-web-apps-behind-apache).

1. Create a folder named `admin` or `rango` or whatever.
2. Create a `.htaccess` inside that folder with the following content:
3. Enable apache modules: `proxy`, `proxy_http`, `rewrite`

```
RewriteEngine on
RewriteRule ^(.*)$ http://localhost:8080/$1 [P,L]
```
