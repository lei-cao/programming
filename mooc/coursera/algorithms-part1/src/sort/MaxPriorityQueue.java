import edu.princeton.cs.algs4.StdOut;

public class MaxPriorityQueue<Key extends Comparable<Key>> {
  private Key[] pq;
  private int N = 0;

  public MaxPriorityQueue(int cap) {
    pq = (Key[]) new Comparable[cap];
  }

  public void insert(Key key) {
    pq[++N] = key;
    swim(N);
  }

  public Key delMax() {
    Key max = pq[1];
    swap(1, N);
    pq[N--] = null;
    sink(1);
    return max;
  }

  public boolean isEmpty() {
    return N == 0;
  }

  public Key max() {
    return pq[1];
  }

  public int size() {
    return N;
  }

  public void print() {
    for (int i = 0; i < N; i++) {
      StdOut.print(pq[i]);
      StdOut.print(", ");
    }
    StdOut.println("");
  }

  /** */
  private void heapfy() {
    for (int i = 0; i < pq.length; i++) {
      int n = i;
      while (n > 0) {
        int p = parentIndex(n);
        if (pq[n].compareTo(pq[p]) <= 0) {
          return;
        }

        swap(n, p);
        n = p;
      }
    }
  }

  /**
   * @param i
   * @param j
   */
  private void swap(int i, int j) {
    Key temp = pq[i];
    pq[i] = pq[j];
    pq[j] = temp;
  }

  /**
   * @param i
   * @return
   */
  private int parentIndex(int i) {
    return i / 2;
  }

  private int leftIndex(int i) {
    return i * 2;
  }

  private int rightIndex(int i) {
    return i * 2 + 1;
  }

  private void sink(int i) {
    while (leftIndex(i) <= N) {
      int j = leftIndex(i);
      if(j+1<=N && less(j, j+1)) {
        j++;
      }
      if (!less(i, j)) {
        return;
      }
      swap(i, j);
      i = j;
    }
  }

  private void swim(int i) {
    while (i > 1) {
      int parent = parentIndex(i);
      if (less(i, parent)) {
        return;
      }
      swap(i, parent);
      i = parent;
    }
  }

  private boolean less(int i, int j) {
    if (pq[i].compareTo(pq[j]) < 0) {
      return true;
    }
    return false;
  }

  public static void main(String[] args) {
    Integer[] a = new Integer[] {5, 6, 7, 1, 2, 3, 8}; // ,2,3,8,9};
    ArrayFactory.printer(a);
    MaxPriorityQueue<Integer> pq = new MaxPriorityQueue<>(10);
    for (int i = 0; i < a.length; i++) {
      pq.insert(a[i]);
    }

    pq.print();
        StdOut.println(pq.delMax());
        StdOut.println(pq.delMax());
        StdOut.println(pq.delMax());
        StdOut.println(pq.delMax());
        StdOut.println(pq.delMax());
        StdOut.println(pq.delMax());
        StdOut.println(pq.delMax());
    //    StdOut.println(pq.delMax());
  }
}
