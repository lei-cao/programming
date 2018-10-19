import edu.princeton.cs.algs4.StdIn;
import edu.princeton.cs.algs4.StdOut;

import java.util.NoSuchElementException;

/**
 * Permutation.
 */
public class Permutation {

    /**
     * Takes an integer k as a command-line argument;
     * reads in a sequence of strings from standard input using StdIn.
     * readString(); and prints exactly k of them, uniformly at random.
     * Print each item from the sequence at most once.
     */
    public static void main(String[] args) {
        int k = 0;
        if (args[0] != null) {
            k = Integer.parseInt(args[0]);
        }
        RandomizedQueue<String> queue = new RandomizedQueue<>();
        String s;
        while (true) {
            try {
                s = StdIn.readString();
            } catch (NoSuchElementException e) {
                break;
            }
            queue.enqueue(s);
        }

        for (int i = 0; i < k; i++) {
            StdOut.println(queue.dequeue());
        }

    }
}
