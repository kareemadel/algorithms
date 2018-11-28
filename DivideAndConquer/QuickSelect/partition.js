function generateRandomNumber(min_value , max_value) 
{
    return Math.floor(Math.random() * (max_value - min_value) + min_value) ;
}

function partition(A, min, max) {
    let pivotIndex = generateRandomNumber(min, max);
    const pivot = A[pivotIndex];
    let temp = pivot, lastSmallerIndex = min;
    A[pivotIndex] = A[min];
    A[min] = temp;
    for (let i = min + 1; i <= max; i++) {
        if (A[i] < pivot) {
            lastSmallerIndex++;
            if (lastSmallerIndex !== i) {
                temp = A[i];
                A[i] = A[lastSmallerIndex];
                A[lastSmallerIndex] = temp;
            }
        }
    }
    A[min] = A[lastSmallerIndex];
    pivotIndex = lastSmallerIndex;
    A[pivotIndex] = pivot;
    return pivotIndex;
}

module.exports = partition;