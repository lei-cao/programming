import edu.princeton.cs.algs4.StdOut;
import edu.princeton.cs.algs4.StdRandom;

import java.util.Iterator;
import java.util.NoSuchElementException;

public class RandomizedQueue<Item> implements Iterable<Item> {
    private Item[] s;
    private int N = 0;

    private class ListIterator implements Iterator<Item> {
        private Item[] copy;
        private int copyN;

        public ListIterator() {
            copy = (Item[]) new Object[N];
            for (int i = 0; i < N; i++ ) {
                copy[i] = s[i];
            }
            copyN = N;
        }

        @Override
        public boolean hasNext() {
            return copyN > 0;
        }

        @Override
        public Item next() {
            int i = StdRandom.uniform(0, copyN);
            Item item = copy[i];
            copy[i] = copy[copyN-1];
            copy[--copyN] = null;
            return item;
        }

        @Override
        public void remove() {
            throw (new UnsupportedOperationException("Remove method is not implemented."));
        }
    }

    // construct an empty randomized queue
    public RandomizedQueue() {
        s = (Item[]) new Object[1];
    }

    // is the randomized queue empty?
    public boolean isEmpty() {
        return N == 0;
    }

    // return the number of items on the randomized queue
    public int size() {
        return N;
    }

    // add the item
    public void enqueue(Item item) {
        if (item == null) {
            throw (new IllegalArgumentException("Input can't be null."));
        }
        if (N == s.length) {
            resize(2 * s.length);
        }
        s[N++] = item;
    }

    // remove and return a random item
    public Item dequeue() {
        if (isEmpty()) {
            throw (new NoSuchElementException("Can't remove from empty queue"));
        }
        int i = StdRandom.uniform(0, N);
        Item item = s[i];
        s[i] = s[N - 1];
        s[--N] = null;
        if (N == s.length / 4) {
            resize(s.length / 2);
        }
        return item;
    }

    // return a random item (but do not remove it)
    public Item sample() {
        if (isEmpty()) {
            throw (new NoSuchElementException("Can't remove from empty queue"));
        }
        int i = StdRandom.uniform(0, N);
        return s[i];
    }

    private void resize(int capacity) {
        Item[] copy = (Item[]) new Object[capacity];
        for (int i = 0; i < N; i++) {
            copy[i] = s[i];
        }
        s = copy;
    }

    // return an independent iterator over items in random order
    @Override
    public Iterator<Item> iterator() {
        return new ListIterator();
    }

    // unit testing (optional)
    public static void main(String[] args) {
        RandomizedQueue<String> queue = new RandomizedQueue<>();
        queue.enqueue("1");
        queue.enqueue("2");
        queue.enqueue("3");
        queue.enqueue("4");
        queue.enqueue("5");
        queue.enqueue("6");
        queue.enqueue("7");
        queue.enqueue("8");

        StdOut.println(queue.dequeue());
        StdOut.println(queue.dequeue());
        StdOut.println(queue.dequeue());
        StdOut.println(queue.dequeue());
        StdOut.println(queue.dequeue());
        StdOut.println(queue.dequeue());
        StdOut.println(queue.dequeue());
        StdOut.println(queue.dequeue());

        queue.enqueue("1");
        queue.enqueue("2");
        queue.enqueue("3");
        queue.enqueue("4");
        StdOut.println(queue.sample());
        StdOut.println(queue.sample());
        StdOut.println(queue.sample());
        StdOut.println(queue.sample());

        for (String q : queue) {
            StdOut.println(q);
        }
    }

}
