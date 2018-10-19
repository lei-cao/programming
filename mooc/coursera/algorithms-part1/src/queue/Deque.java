import edu.princeton.cs.algs4.StdOut;

import java.util.Iterator;
import java.util.NoSuchElementException;

public class Deque<Item> implements Iterable<Item> {
    private Node head, tail;
    private int size = 0;

    private class ListIterator implements Iterator<Item> {
        private Node current = head.next;

        @Override
        public boolean hasNext() {
            return current != tail;
        }

        @Override
        public Item next() {
            if (!hasNext()) {
                throw (new NoSuchElementException("Can't remove from empty queue"));
            }
            Item i = current.item;
            current = current.next;
            return i;
        }

        @Override
        public void remove() {
            throw (new UnsupportedOperationException("Remove method is not implemented."));
        }
    }

    private class Node {
        private Item item;
        private Node next;
        private Node prev;
    }

    // construct an empty deque
    public Deque() {
        head = new Node();
        tail = new Node();
        head.next = tail;
        tail.prev = head;
    }

    // is the deque empty?
    public boolean isEmpty() {
        return size == 0;
    }

    // return the number of items on the deque
    public int size() {
        return size;
    }

    // add the item to the front
    public void addFirst(Item item) {
        if (item == null) {
            throw (new IllegalArgumentException("Input can't be null."));
        }
        Node node = new Node();
        node.item = item;
        node.next = head.next;
        node.prev = head;
        head.next.prev = node;
        head.next = node;
        size++;
    }

    // add the item to the end
    public void addLast(Item item) {
        if (item == null) {
            throw (new IllegalArgumentException("Input can't be null."));
        }
        Node node = new Node();
        node.item = item;
        node.next = tail;
        tail.prev.next = node;
        node.prev = tail.prev;
        tail.prev = node;
        size++;
    }

    // remove and return the item from the front
    public Item removeFirst() {
        if (isEmpty()) {
            throw (new NoSuchElementException("Can't remove from empty queue"));
        }
        Item item = head.next.item;
        Node temp = head.next;
        head.next.next.prev = head;
        head.next = head.next.next;
        temp.next = null;
        temp.prev = null;
        size--;
        return item;
    }

    // remove and return the item from the end
    public Item removeLast() {
        if (isEmpty()) {
            throw (new NoSuchElementException("Can't remove from empty queue"));
        }
        Item item = tail.prev.item;
        Node temp = tail.prev;
        tail.prev.prev.next = tail;
        tail.prev = tail.prev.prev;
        temp.prev = null;
        temp.next = null;
        size--;
        return item;
    }

    public Iterator<Item> iterator() {
        return new ListIterator();
    }

    // unit testing (optional)
    public static void main(String[] args) {
        Deque<String> deque = new Deque<>();
        deque.addFirst("c");
        deque.addFirst("b");
        deque.addFirst("a");
        deque.addFirst("a");
        deque.addFirst("a");
        deque.addFirst("a");
        deque.removeFirst();
        deque.removeFirst();
        deque.removeFirst();
        deque.addLast("x");
        deque.addLast("y");
        deque.addLast("z");
        deque.addLast("z");
        deque.addLast("z");
        deque.addLast("z");
        deque.removeLast();
        deque.removeLast();
        deque.removeLast();

        Iterator<String> iterator = deque.iterator();
        while (iterator.hasNext()) {
            String i = iterator.next();
            StdOut.println(i);
        }
    }


}
