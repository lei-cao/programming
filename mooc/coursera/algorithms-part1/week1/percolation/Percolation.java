/* *****************************************************************************
 *  Name:
 *  Date:
 *  Description:
 **************************************************************************** */

import edu.princeton.cs.algs4.WeightedQuickUnionUF;

public class Percolation {
    private final int size;
    private final boolean[] sites;
    private final WeightedQuickUnionUF ufTop;
    private final WeightedQuickUnionUF ufBottom;
    private int openSitesCount;

    // create n-by-n grid, with all sites blocked
    public Percolation(int n) {
        if (n <= 0) {
            throw (new IllegalArgumentException("grid size must be greater than zero."));
        }
        openSitesCount = 0;
        size = n;
        sites = new boolean[n * n];
        ufTop = new WeightedQuickUnionUF(n * n + 2);
        ufBottom = new WeightedQuickUnionUF(n * n + 2);
        for (int i = 0; i < n * n; i++) {
            sites[i] = false;
        }
    }

    // open site (row, col) if it is not open already
    public void open(int row, int col) {
        rangeCheck(row, col);
        if (isOpen(row, col)) {
            return;
        }
        sites[index(row, col)] = true;
        openSitesCount++;
        unionAll(row, col);
    }

    // is site (row, col) open?
    public boolean isOpen(int row, int col) {
        rangeCheck(row, col);
        return sites[index(row, col)];
    }

    // is site (row, col) full?
    public boolean isFull(int row, int col) {
        rangeCheck(row, col);
        return ufTop.connected(0, (row - 1) * size + col);
    }

    // number of open sites
    public int numberOfOpenSites() {
        return openSitesCount;
    }

    private int index(int row, int col) {
        return (row - 1) * size + col - 1;
    }

    private int ufIndex(int row, int col) {
        return (row - 1) * size + col;
    }

    private void unionAll(int row, int col) {
        if (row == 1) {
            ufTop.union(0, ufIndex(row, col));
            ufBottom.union(0, ufIndex(row, col));
        }
        if (row == size) {
            ufBottom.union(size * size + 1, ufIndex(row, col));
        }
        doUnion(row, col, row - 1, col);
        doUnion(row, col, row + 1, col);
        doUnion(row, col, row, col - 1);
        doUnion(row, col, row, col + 1);
    }

    private void doUnion(int row1, int col1, int row2, int col2) {
        if (!inRange(row1, col1) || !inRange(row2, col2)) {
            return;
        }
        if (!isOpen(row1, col1) || !isOpen(row2, col2)) {
            return;
        }
        ufTop.union(ufIndex(row1, col1), ufIndex(row2, col2));
        ufBottom.union(ufIndex(row1, col1), ufIndex(row2, col2));
    }

    // does the system percolate?
    public boolean percolates() {
        return ufBottom.connected(0, size * size + 1);
    }

    // test client (optional)
    public static void main(String[] args) {
        /**
         * Main
         */

    }

    private void rangeCheck(int row, int col) {
        if (!inRange(row, col)) {
            throw (new IllegalArgumentException(String.format("row: %d, col: %d) is out of range %d", row, col, size)));
        }
    }

    private boolean inRange(int row, int col) {
        if (row < 1 || col < 1 || row > size || col > size) {
            return false;
        }
        return true;
    }
}

