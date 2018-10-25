import edu.princeton.cs.algs4.StdOut;

public class ShellSort {
    public void sort(int[] a) {
        int N = a.length;
        int h = 1;
        while (h < N / 3) {
            h = 3 * h + 1;
        }

        while (h >= 1) {
            insertSrot(a, h);
            h = h / 3;
        }
    }

    private void insertSrot(int[] a, int h) {
        for (int i = h; i < a.length; i++) {
            for (int j = i; j >= h && a[j] < a[j - h]; j = j - h) {
                int temp = a[j];
                a[j] = a[j - h];
                a[j - h] = temp;
            }
        }
    }

    public static void main(String[] args) {
        int[] a = ArrayFactory.intArray(10);

        ArrayFactory.printer(a);
        ShellSort sort = new ShellSort();
        sort.sort(a);
        ArrayFactory.printer(a);
    }
}


