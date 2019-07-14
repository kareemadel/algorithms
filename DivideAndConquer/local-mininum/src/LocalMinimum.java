
public class LocalMinimum {
    public static void main(String[] args) {
        Grid g = Grid.getGrid(30, 1000);
        System.out.println(g);
        Point localMinimum = findLocalMin(g);
        System.out.println(localMinimum);
    }

    static Point findLocalMin(Grid g) {
        Point min = new Point();
        int[] columns = {0, (g.getWidth()-1)/2, g.getWidth()-1};
        int[] rows = {0, (g.getHeight()-1)/2, g.getHeight()-1};
        for(int j : columns) {
            for(int i = 0; i < g.getHeight(); i++) {
                if (g.isLocalMin(i, j)) {
                    min.setLocation(i, j);
                    return min;
                } else if (g.getValue(i, j) < g.getValue(min)) {
                    min.setLocation(i, j);
                }
            }
        }
        for(int i : rows) {
            for(int j = 0; j < g.getWidth(); j++) {
                if (g.isLocalMin(i, j)) {
                    min.setLocation(i, j);
                    return min;
                } else if (g.getValue(i, j) < g.getValue(min)) {
                    min.setLocation(i, j);
                }
            }
        }
        Point smallestNeighbor = new Point(min);
        Point right = g.getRightNeighbor(min);
        Point left = g.getLeftNeighbor(min);
        Point up = g.getTopNeighbor(min);
        Point bottom = g.getBottomNeighbor(min);
        if (right != null && g.getValue(right) < g.getValue(smallestNeighbor)) {
            smallestNeighbor.setLocation(right);
        }
        if (left != null && g.getValue(left) < g.getValue(smallestNeighbor)) {
            smallestNeighbor.setLocation(left);
        }
        if (up != null && g.getValue(up) < g.getValue(smallestNeighbor)) {
            smallestNeighbor.setLocation(up);
        }
        if (bottom != null && g.getValue(bottom) < g.getValue(smallestNeighbor)) {
            smallestNeighbor.setLocation(bottom);
        }
        int i0, j0, i1, j1;
        if (smallestNeighbor.getColumn() < columns[1]) {
            j0 = 0;
            j1 = columns[1]-1;
        } else {
            j0 = columns[1]+1;
            j1 = columns[2]-1;
        }
        if (smallestNeighbor.getRow() < rows[1]) {
            i0 = 0;
            i1 = rows[1]-1;
        } else {
            i0 = rows[1]+1;
            i1 = rows[2]-1;
        }
        return findLocalMin(g.subGrid(i0, j0, i1, j1));
    }
}
