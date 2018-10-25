import edu.princeton.cs.algs4.StdOut;
import edu.princeton.cs.algs4.StdRandom;

public class ArrayFactory {
    public static int[] intArray(int length) {
        int[] a = new int[length];
        for (int i = 0; i < length; i++) {
            a[i] = StdRandom.uniform(0, length);
        }
        return a;
    }

    public static void printer(int[] a) {
        for (int i = 0; i < a.length; i++) {
            StdOut.print(a[i]);
            StdOut.print(", ");
        }
        StdOut.println("");
    }
}
