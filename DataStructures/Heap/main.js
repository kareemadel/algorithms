const { buildMaxHeap } = require('./buildHeap');
const heapSort = require('./heapSort');

const A = [2, 17, 1, 19, 100, 3, 36, 7, 25, 90];
console.log(`${buildMaxHeap(A)}`);
console.log(`${heapSort(A)}`);