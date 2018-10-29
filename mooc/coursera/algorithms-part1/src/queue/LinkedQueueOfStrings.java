/* *****************************************************************************
 *  Name:
 *  Date:
 *  Description:
 **************************************************************************** */

import edu.princeton.cs.algs4.StdOut;

public class LinkedQueueOfStrings {

    private static class Node {
        String item;
        Node next;

        public Node(String str) {
            item = str;
        }
        public Node() {}
    }

    private Node first, last;
    private int size;

    public void enqueue(String item) {
        Node oldLast = last;
        last = new Node(item);
        // Special case for empty queue
        if (isEmpty()) {
            first = last;
        } else {
            oldLast.next = last;
        }
        size++;
    }

    public String dequeue() {
        String item = first.item;
        first = first.next;
        // Special case for empty queue
        if (isEmpty()) {
            last = first;
        }
        size--;
        return item;
    }

    boolean isEmpty() {
        return first== null;
    }

    int size() {
        return size;
    }

    // test client (optional)
    public static void main(String[] args) {
        LinkedQueueOfStrings queue = new LinkedQueueOfStrings();

        StdOut.println(queue.isEmpty());
        queue.enqueue("1");
        queue.enqueue("2");
        StdOut.println(queue.isEmpty());
        StdOut.println(queue.dequeue());
        StdOut.println(queue.dequeue());

        queue.enqueue("1");
        queue.enqueue("2");
        StdOut.println(queue.dequeue());
        StdOut.println(queue.dequeue());


        queue.enqueue("1");
        StdOut.println(queue.dequeue());
        queue.enqueue("2");
        StdOut.println(queue.dequeue());

    }

}
