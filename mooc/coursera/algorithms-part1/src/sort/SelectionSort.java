import edu.princeton.cs.algs4.StdOut;

public class SelectionSort {
    public void sort(int[] a) {
        for (int i = 0; i < a.length; i++) {
            int min = i;
            for (int j = i + 1; j < a.length; j++) {
                if (a[j] < a[min]) {
                    min = j;
                }
            }
            int temp = a[i];
            a[i] = a[min];
            a[min] = temp;
        }
    }

    public static void main(String[] args) {
        int[] a = ArrayFactory.intArray(10);
        ArrayFactory.printer(a);

        SelectionSort sort = new SelectionSort();
        sort.sort(a);
        StdOut.println("xxxxxxxxx");
        ArrayFactory.printer(a);
    }
}
