import edu.princeton.cs.algs4.StdOut;

public class BottomUpMergeSort {

    public void sort(int[] a) {

        int[] temp = new int[a.length];
        for (int i = 0; i < a.length; i++) {
            temp[i] = a[i];
        }

        for (int k = 2;  k/2 <= a.length*2; k = k*2) {
            for (int i = 0; i < a.length; i = i + k) {
                    int low1 = Math.min(i, a.length-1);
                    int high1 = Math.min(i+k/2-1, a.length-1);
                    int low2 = Math.min(high1+1, a.length-1);
                    int high2 = Math.min(low2+k/2-1, a.length-1);
                    merge(temp, low1, high1, low2, high2, a);

            }

            int[] b = temp;
            temp = a;
            a = b;

        }
    }

    private void merge(int[] from, int low1, int high1, int low2, int high2, int[] to) {
        for (int i = low1; i <= high2; i++) {
            if (low1 <= high1 && low2 <= high2) {
                if (from[low1] <= from[low2]) {
                    to[i] = from[low1++];
                } else {
                    to[i] = from[low2++];
                }
            } else if (low1 <= high1) {
                to[i] = from[low1++];
            } else if (low2 <= high2) {
                to[i] = from[low2++];
            } else {
                break;
            }
        }
    }

    public static void main(String[] args) {
        int[] a = ArrayFactory.intArray(100);
        ArrayFactory.printer(a);
        StdOut.println("xxxxxxxxx");

        BottomUpMergeSort sort = new BottomUpMergeSort();
        sort.sort(a);
        ArrayFactory.printer(a);
    }
}
