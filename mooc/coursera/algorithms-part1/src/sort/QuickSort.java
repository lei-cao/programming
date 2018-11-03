import edu.princeton.cs.algs4.StdOut;

public class QuickSort {
  public static void sort(int[] a) {
    threeWaySort(a, 0, a.length - 1);
  }

  /**
   * lo_________lt_______i________gt____hi;
   * lo is the index of the pivot. lt is the first index where the element == pivot
   * i is the first index where the element haven't checked
   * gt is the last index where the element haven't checked
   * i <-> gt are the elements haven't partitioned.
   * if a[i] > pivot, swap a[i] with a[gt], gt--; only do gt-- because that was a[i] who is checked.
   *    don't do i++ because a[i] (a[gt] is not checked after swap
   * if a[i] < pivot, swap a[i] with a[lt], i++, lt++; a[lt] == pivot, so swap a[i] with a[lt] then do lt++, i++
   * if a[i] == pivot, i++, check next i.
   * @param a
   * @param lo
   * @param hi
   */
  private static void threeWaySort(int[] a, int lo, int hi) {
    if (lo >= hi) {
      return;
    }

    int lt = lo;
    int gt = hi;
    int i = lo;
    int pivot = a[lo];

    while (i <= gt) {
      if (a[i] > pivot) {
        // i doesn't need to ++ here
        swap(a, i, gt--);
      } else if (a[i] < pivot) {
        swap(a, lt++, i++);
      } else {
        i++;
      }
    }

    threeWaySort(a, lo, lt - 1);
    threeWaySort(a, gt + 1, hi);
  }

  private static void sort(int[] a, int lo, int hi) {
    // Need check equal
    if (lo >= hi) {
      return;
    }

    // Return the sorted position
    int j = partition(a, lo, hi);
    // sort the lower part
    sort(a, lo, j - 1);
    // sort the higher part
    sort(a, j + 1, hi);
  }

  private static int partition(int[] a, int lo, int hi) {
    // local indexer. j -> hi + 1 for simplifying the operation.
    // because we chose lo as pivot, it's like i = start-1
    int i = lo, j = hi + 1;

    while (true) {
      // loop i to the first place where a[i] >= pivot
      // equal to pivot is included for swapping.
      while (a[++i] < a[lo]) {
        if (i == hi) {
          break;
        }
      }

      // equal is included for swapping as well.
      while (a[--j] > a[lo]) {
        if (j == lo) {
          break;
        }
      }

      if (i >= j) {
        break;
      }

      swap(a, i, j);
    }

    swap(a, lo, j);
    return j;
  }

  private static void my_sort(int[] a, int start, int end) {
    if (start >= end) {
      return;
    }
    int pivot = end;

    int i = start;
    int j = pivot - 1;
    // 4, 5, 1, 2
    while (i < j) {
      if (a[i] <= a[pivot]) {
        i++;
      }
      if (a[j] > a[pivot]) {
        j--;
      }
      if (a[i] > a[pivot] && a[j] <= a[pivot]) {
        swap(a, i, j);
        i++;
        j--;
      }
    }

    if (a[j] > a[pivot]) {
      swap(a, j, pivot);
      pivot = j;
    } else {
      swap(a, j + 1, pivot);
      pivot = j + 1;
    }

    my_sort(a, start, pivot - 1);
    my_sort(a, pivot + 1, end);
  }

  private static void swap(int[] a, int i, int j) {
    int temp = a[i];
    a[i] = a[j];
    a[j] = temp;
  }

  public static void main(String[] args) {
    int[] a = ArrayFactory.intArray(10);
    //    int[] a = new int[] {4, 5, 1, 7, 2,6,8,9,7};
    ArrayFactory.printer(a);

    QuickSort.sort(a);
    StdOut.println("xxxxxxxxx");
    ArrayFactory.printer(a);
  }
}
