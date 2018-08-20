package quick.sort;

import java.util.ArrayList;
import java.util.Collection;
import java.util.Collections;
import java.util.Comparator;
import java.util.concurrent.ThreadLocalRandom;

public class QuickSortArrayList<T> extends ArrayList<T> {

    public QuickSortArrayList () {
        super();
    }

    public QuickSortArrayList (int initialCapacity) {
        super(initialCapacity);
    }

    public QuickSortArrayList (Collection<T> c) {
        super(c);
    }

    // sort the list and returns the number of pivot comparisons
    public long qSort(Comparator comp) {
        return quickSort(comp, 0, this.size() - 1);
    }

    private long quickSort(Comparator comp, int start, int end) {
        if (start >= end) {
            return 0L;
        }
        int partitionIndex = partitionList(comp, start, end);
        long c1 = quickSort(comp, start, partitionIndex - 1);
        long c2 = quickSort(comp, partitionIndex + 1, end);
        return c1 + c2 + (long) (end - start);
    }

    private int partitionList(Comparator comp, int start, int end) {
        // use the middle of start, end and middle element to improve performance for sorted lists
//        int medianIndex = middleOfThree(comp, start, (start + end) / 2, end);
//        Collections.swap(this, start, medianIndex);
        // choose random pivot
        Collections.swap(this, start, ThreadLocalRandom.current().nextInt(start, end + 1));
        int pivotIndex = start;
        T pivot = this.get(pivotIndex);
        int partitionIndex = start;
        for (int i = partitionIndex; i < end; i++) {
            if(comp.compare(this.get(i + 1), pivot) < 0) {
                partitionIndex++;
                Collections.swap(this, partitionIndex, i+1);
            }
        }
        Collections.swap(this, pivotIndex, partitionIndex);
        return partitionIndex;
    }

    private int middleOfThree(Comparator comp, int start, int middle, int end) {
        T startObj = this.get(start);
        T endObj = this.get(end);
        T midObj = this.get(middle);
        if (comp.compare(startObj, midObj) > 0) {
            return assistedMidOfThree(comp, start, (T) startObj, middle, (T) midObj, end, (T) endObj);
        } else {
            return assistedMidOfThree(comp, middle, (T) midObj, start, (T) startObj, end, (T) endObj);
        }
    }

    private int assistedMidOfThree(Comparator comp, int indexBigger, T biggerObj, int indexSmaller, T samllerObj, int indexUnknown, T unknownObj) {
        if (comp.compare(samllerObj, unknownObj) > 0) {
            return indexSmaller;
        } else if (comp.compare(biggerObj, unknownObj) > 0) {
            return indexUnknown;
        } else {
            return indexBigger;
        }
    }
}
