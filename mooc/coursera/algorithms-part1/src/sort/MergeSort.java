import edu.princeton.cs.algs4.StdOut;

public class MergeSort {

    public void sort(int[] a) {

        int[] temp = new int[a.length];
        for (int i = 0; i < a.length; i++) {
            temp[i] = a[i];
        }

        mergeSort(temp, 0, a.length - 1, a);


    }

    private void mergeSort(int[] from, int low, int high, int[] to) {
        // This condition is there is only one ele
        // high < low is avoid empty array
        if (high <= low) {
            return;
        }

        // The mid index between two indexes
        // 4,5 -> 4
        // 6,7,8 -> 7
//        int mid = low + (high-low) / 2;
        int mid = (low + high) >>> 1;

        mergeSort(to, low, mid, from);
        mergeSort(to, mid + 1, high, from);

        if (from[mid + 1] > from[mid]) {
            return;
        }
        merge(from, low, mid, mid + 1, high, to);
    }

    private void merge(int[] from, int low1, int high1, int low2, int high2, int[] to) {
        // 0, 0, 1, 1
        int i = low1;
        int j = low2;
        // Condition of k is loop from low -> high before split (low1 -> high1 (mid) -> low2 (mid +1) -> high2)
        for (int k = low1; k <= high2; k++) {
            // i<= high1 && j<= high2, when there is only one elem, low1 == high1 && low2 == high2
            if (low1 <= high1 && low2 <= high2) {
                if (from[low1] <= from[low2]) {
                    to[k] = from[low1];
                    low1++;
                } else {
                    to[k] = from[low2];
                    low2++;
                }
            } else if (low1 <= high1) {
                to[k] = from[low1];
                low1++;
            } else if (low2 <= high2) {
                to[k] = from[low2];
                low2++;
            } else {
                break;
            }
        }

    }

    public static void main(String[] args) {
        int[] a = ArrayFactory.intArray(10);
        ArrayFactory.printer(a);

        MergeSort sort = new MergeSort();
        sort.sort(a);
        StdOut.println("xxxxxxxxx");
        ArrayFactory.printer(a);
    }
}
