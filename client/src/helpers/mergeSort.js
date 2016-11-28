/**
 * This adds mergeSort method for arrays.
 *
 * @summary Fast and stable sorting.
 *
 * Created by michael on 7/20/2016.
 * Updated by Maks 8/17/2016.
 */

(function () {
  // Add stable merge sort method to Array prototype
  if (!Array.mergeSort) {
    Array.prototype.mergeSort = function (compare) {

      var length = this.length,
          middle = length / 2 >> 0;

      // define default comparison function if none is defined
      if (!compare) {
        compare = function (left, right) {
          return left >= right;
        };
      }

      if (length < 2) {
        return this;
      }

      function merge(left, right, compare) {
        var result = [];
        while(left.length || right.length) {
          result.push( !right.length || left.length && compare(left[0], right[0]) ? left.shift() : right.shift() );
        }
        return result;
      }

      return merge(
        this.slice(0, middle).mergeSort(compare),
        this.slice(middle, length).mergeSort(compare),
        compare
      );
    };
  }
}());
