import edu.princeton.cs.algs4.StdIn;
import edu.princeton.cs.algs4.StdRandom;
import edu.princeton.cs.algs4.StdStats;

public class PercolationStats {
    private final double[] thresholds;

    // perform trials independent experiments on an n-by-n grid
    public PercolationStats(int n, int trials) {
        if (n <= 0 || trials <= 0) {
            throw (new IllegalArgumentException(String.format("n: %d, and trials: %d need to greater than 0", n, trials)));
        }

        thresholds = new double[trials];

        for (int i = 0; i < trials; i++) {
            Percolation pero = new Percolation(n);
            while (!pero.percolates()) {
                int row = StdRandom.uniform(1, n + 1);
                int col = StdRandom.uniform(1, n + 1);
                pero.open(row, col);
            }
            double size = n * n;
            thresholds[i] = pero.numberOfOpenSites() / size;
        }
    }

    // sample mean of percolation threshold
    public double mean() {
        return StdStats.mean(thresholds);
    }

    // sample standard deviation of percolation threshold
    public double stddev() {
        return StdStats.stddev(thresholds);
    }

    // low  endpoint of 95% confidence interval
    public double confidenceLo() {
        return StdStats.min(thresholds);
    }

    // high endpoint of 95% confidence interval
    public double confidenceHi() {
        return StdStats.max(thresholds);
    }

    // test client (described below)
    public static void main(String[] args) {
        int n;
        int trials;

        if (args.length == 2) {
            n = Integer.parseInt(args[0]);
            trials = Integer.parseInt(args[1]);
        } else {
            throw (new IllegalArgumentException(String.format("Please input two integers.")));
        }

        PercolationStats stats = new PercolationStats(n, trials);

        System.out.printf("mean = %f\n", stats.mean());
        System.out.printf("stddev = %f\n", stats.stddev());
        System.out.printf("95%% confidence interval = [%f, %f]\n", stats.confidenceLo(), stats.confidenceHi());
    }

}
