var gulp = require('gulp');

gulp.task('build', function () {
    return gulp.src(__dirname + '/client/dist/static/**/*').pipe(gulp.dest(__dirname + "/public"));
});

gulp.task('css', function () {
    return gulp.src(__dirname + '/client/src/client/public/style.css').pipe(gulp.dest(__dirname + "/public/css"));
});

gulp.task('js', function () {
    return gulp.src(__dirname + '/client/src/client/public/app.bundle.js').pipe(gulp.dest(__dirname + "/public/js"));
});