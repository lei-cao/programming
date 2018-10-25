import edu.princeton.cs.algs4.StdOut;

public class InsertionSort {
    public void sort(int[] a) {
        for (int i = 0; i < a.length; i++) {
            for (int j = i; j > 0; j--) {
                if (a[j] < a[j-1]) {
                    int temp = a[j];
                    a[j] = a[j-1];
                    a[j-1] = temp;
                } else {
                    break;
                }
            }
        }
    }

    public static void main(String[] args) {
        int[] a = ArrayFactory.intArray(10);

        ArrayFactory.printer(a);
        InsertionSort sort = new InsertionSort();
        sort.sort(a);
        StdOut.println("xxxxxxxxx");
        ArrayFactory.printer(a);
    }
}
