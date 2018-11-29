const { left, right } = require('./helpers');

const maxHeapify = (A, i, heapSize) => {
    const l = left(i);
    const r = right(i);
    let largest = i;
    if ( l < heapSize && A[l] > A[i]) largest = l;
    if ( r < heapSize && A[r] > A[largest]) largest = r;
    if (largest !== i) {
        const temp = A[i];
        A[i] = A[largest];
        A[largest] = temp;
        maxHeapify(A, largest, heapSize);
    }
};

const buildMaxHeap = (A) => {
    const arr = [...A];
    for (let i = Math.trunc(arr.length/2); i >= 0; i--) {
        maxHeapify(arr, i, arr.length);
    }
    return arr;
};

module.exports = {
    maxHeapify,
    buildMaxHeap
};
