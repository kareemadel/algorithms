import java.util.*;

public class Grid {
    private int[][] arr;

    private Grid(int rows, int columns) {
        this.arr = new int[rows][columns];
    }

    public static Grid getGrid(int n, int max) {
        Grid g = new Grid(n, n);
        g.arr = new int[n][n];
        Map<Integer, Boolean> doesExist = new HashMap<>();
        int buffer;
        Random rand = new Random();
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < n; j++) {
                buffer = rand.nextInt(max);
                while (doesExist.containsKey(buffer)) {
                    buffer = rand.nextInt(max);
                }
                g.arr[i][j] = buffer;
                doesExist.put(buffer, true);
            }
        }
        return g;
    }

    public Grid subGrid(int i0, int j0, int i1, int j1) {
        Grid g = new Grid(i1-i0+1, j1-j0+1);
        for(int i = i0; i <= i1; i++) {
            for (int j = j0; j <= j1; j++) {
                g.arr[i][j] = this.arr[i][j];
            }
        }
        return g;
    }

    public int getValue(int i, int j) {
        return arr[i][j];
    }

    public int getValue(Point p) {
        return arr[p.getRow()][p.getColumn()];
    }

    boolean isLocalMin(int i, int j) {
        int val = getValue(i, j);
        Point p = new Point(i, j);
        Point right = getRightNeighbor(p);
        Point left = getLeftNeighbor(p);
        Point up = getTopNeighbor(p);
        Point bottom = getBottomNeighbor(p);
        if (right != null && getValue(right) < val) {
            return false;
        }
        if (left != null && getValue(left) < val) {
            return false;
        }
        if (up != null && getValue(up) < val) {
            return false;
        }
        if (bottom != null && getValue(bottom) < val) {
            return false;
        }
        return true;
    }

    public Point getRightNeighbor(Point p) {
        int i = p.getRow();
        int j = p.getColumn();
        if (j < getWidth()-1 && j >= 0 ) {
            return new Point(i, j+1);
        }
        return null;
    }

    public Point getLeftNeighbor(Point p) {
        int i = p.getRow();
        int j = p.getColumn();
        if (j < getWidth() && j > 0 ) {
            return new Point(i, j-1);
        }
        return null;
    }

    public Point getTopNeighbor(Point p) {
        int i = p.getRow();
        int j = p.getColumn();
        if (i < getHeight() && i > 0 ) {
            return new Point(i-1, j);
        }
        return null;
    }

    public Point getBottomNeighbor(Point p) {
        int i = p.getRow();
        int j = p.getColumn();
        if (i < getHeight()-1 && i >= 0 ) {
            return new Point(i+1, j);
        }
        return null;
    }

    public int getHeight() {
        return arr.length;
    }

    public int getWidth() {
        return arr[0].length;
    }

    @Override
    public String toString() {
        StringJoiner rowJoiner = new StringJoiner("\n", "[\n", "\n]");
        for(int i = 0; i < getHeight(); i++) {
            StringJoiner columnJoiner = new StringJoiner(", ");
            for (int j = 0; j < getWidth(); j++) {
                columnJoiner.add(String.format("%3d", getValue(i, j)));
            }
            rowJoiner.add(columnJoiner.toString());
        }
        return rowJoiner.toString();
    }
}
