public class HeapSort {
  public static void sort(Comparable[] a) {
    int N = a.length;
    for (int i = N / 2; i >= 1; i--) {
      sink(a, N, i);
    }

    while (N > 1) {
      swap(a, 1, N--);
      sink(a, N, 1);
    }
  }

  private static void sink(Comparable[] a, int N, int i) {
    while (2 * i <= N) {
      int j = 2 * i;
      if (j < N && less(a, j, j + 1)) {
        j++;
      }
      if (!less(a, i, j)) {
        return;
      }
      swap(a, i, j);
      i = j;
    }
  }

  private static boolean less(Comparable[] a, int i, int j) {
    if (a[i - 1].compareTo(a[j - 1]) < 0) {
      return true;
    }
    return false;
  }

  private static void swap(Comparable[] a, int i, int j) {
    Comparable temp = a[i - 1];
    a[i - 1] = a[j - 1];
    a[j - 1] = temp;
  }

  public static void main(String[] args) {
    Integer[] a = ArrayFactory.IntegerArray(20);
    ArrayFactory.printer(a);
    HeapSort.sort(a);
    ArrayFactory.printer(a);
  }
}
