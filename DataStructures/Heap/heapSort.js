const { buildMaxHeap, maxHeapify } = require('./buildHeap');

const heapSort = (A) => {
    const arr = buildMaxHeap(A);
    let heapSize = arr.length;
    let temp;
    for (let i = arr.length - 1; i >= 1; i--) {
        temp = arr[i];
        arr[i] = arr[0];
        arr[0] = temp;
        maxHeapify(arr, 0, --heapSize);
    }
    return arr;
};

module.exports = heapSort;
