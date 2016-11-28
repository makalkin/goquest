var gulp = require('gulp');

gulp.task('build', function () {
    return  gulp.src(__dirname + '/client/dist/static/**/*').stream.pipe(gulp.dest(__dirname + "/public"));
});