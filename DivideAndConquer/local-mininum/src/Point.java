public class Point {
    private int row;
    private int column;

    public Point() {
    }

    public Point(int row, int column) {
        this.row = row;
        this.column = column;
    }

    public Point(Point p) {
        this.row = p.getRow();
        this.column = p.getColumn();
    }

    public int getRow() {
        return row;
    }

    public void setRow(int row) {
        this.row = row;
    }

    public int getColumn() {
        return column;
    }

    public void setColumn(int column) {
        this.column = column;
    }

    public void setLocation(int row, int column) {
        this.row = row;
        this.column = column;
    }

    @Override
    public String toString() {
        return "Point{" +
                "row=" + row +
                ", column=" + column +
                '}';
    }

    public void setLocation(Point p) {
        this.row = p.getRow();
        this.column = p.getColumn();
    }
}
