package inversions.counter;

import java.util.ArrayList;
import java.util.Collection;

public class MyArrayList<E> extends ArrayList<E> {
    private GreaterThan<E> predicate;
    private long inversions = 0;

    public MyArrayList () {
        super();
    }

    public MyArrayList (int initialCapacity) {
        super(initialCapacity);
    }

    public MyArrayList (Collection<E> c) {
        super(c);
    }

    public long getInversions() {
        return inversions;
    }

    public MyArrayList<E> sort(GreaterThan<E> predicate) {
        inversions = 0L;
        this.predicate = predicate;
        this.mergeSort(0, this.size() - 1);
        return this;
    }

    private void mergeSort(int start, int end) {
        if (start >= end) {
            return ;
        } else {
            int half = (start + end) / 2;
            this.mergeSort(start, half);
            this.mergeSort(half + 1, end);
            this.merge(start, half, end);
            return ;
        }
    }

    private void merge(int start, int half, int end) {
        int totalSize = end - start + 1;
        ArrayList<E> merged = new ArrayList<>();
        for (int lIndex = start, rIndex = half + 1, tIndex = 0; tIndex < totalSize; tIndex++) {
            if ((rIndex <= end) && (lIndex > half || predicate.greaterThan(this.get(lIndex), this.get(rIndex)))) {
                merged.add(tIndex, this.get(rIndex));
                inversions += half - lIndex + 1;
                rIndex++;
            } else {
                merged.add(tIndex, this.get(lIndex));
                lIndex++;
            }
        }
        for (int i = 0; i < totalSize; i++) {
            this.set(start + i, merged.get(i));
        }
        return;
    }
}
