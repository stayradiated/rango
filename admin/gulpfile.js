'use strict';

var gulp         = require('gulp');
var source       = require('vinyl-source-stream');
var browserSync  = require('browser-sync');
var sass         = require('gulp-sass');
var autoprefixer = require('gulp-autoprefixer');
var reactify     = require('reactify');
var browserify   = require('browserify');
var watchify     = require('watchify');
var uglify       = require('gulp-uglify');

gulp.task('default', ['lib', 'style'], function () {
  gulp.watch('./style/**/*.scss', ['style']);
});

gulp.task('server', ['default'], function () {
  return browserSync({
    notify: false,
    ghostMode: false,
    server: {
      baseDir: './dist/',
    }
  });
}); 

gulp.task('proxy', ['default'], function () {
  return browserSync({
    notify: false,
    open: false,
    ghostMode: false,
    proxy: 'localhost:8080',
  });
});

gulp.task('lib', function () {
  var bundler = watchify(browserify({
    cache: {},
    packageCache: {},
    fullPaths: true,
    extensions: '.jsx'
  }));

  bundler.add('./lib/main.jsx');
  bundler.transform(reactify);

  bundler.on('update', rebundle);

  function rebundle () {
    console.log('rebundling');
    return bundler.bundle()
      .on('error', function (err) {
        console.log(err.message);
      })
      .pipe(source('main.js'))
      .pipe(gulp.dest('./dist'))
      .pipe(browserSync.reload({stream: true}));
  }

  return rebundle();
});

gulp.task('style', function () {
  return gulp.src('./style/main.scss')
    .pipe(sass({errLogToConsole: true, outputStyle: 'compressed'}))
    .pipe(autoprefixer())
    .pipe(gulp.dest('./dist'))
    .pipe(browserSync.reload({stream: true}));
});

gulp.task('minify', function () {
  return gulp.src('./dist/*.js')
    .pipe(uglify())
    .pipe(gulp.dest('./dist'));
});
