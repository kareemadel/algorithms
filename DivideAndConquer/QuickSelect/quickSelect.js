const partition = require('./partition');

function qSelect(A, min, max, r) {
    if ( min === max && r === min) return A[min];
    const pivotIndex = partition(A, min, max);
    if (pivotIndex === r) return A[pivotIndex];
    else if (r > pivotIndex) {
        return qSelect(A, pivotIndex + 1, max, r)
    } else {
        return qSelect(A, min, pivotIndex - 1, r);
    }
}

function select(A, r) {
    if (r < 1 || r > A.length) return null;
    const arr = [...A];
    return qSelect(arr, 0, A.length - 1, r - 1);
}

module.exports = select;