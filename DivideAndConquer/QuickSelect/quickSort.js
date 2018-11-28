const partition = require('./partition');

function qSort(A, min, max) {
    if (min >= max) return A;
    const pivotIndex = partition(A, min, max);
    qSort(A, min, pivotIndex - 1);
    qSort(A, pivotIndex + 1, max);
    return A;
}

function sort(A) {
    const arr = [...A];
    qSort(arr, 0, A.length - 1);
    return arr;
}

module.exports = sort;