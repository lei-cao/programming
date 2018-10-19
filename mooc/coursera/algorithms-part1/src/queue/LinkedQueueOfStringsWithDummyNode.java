/* *****************************************************************************
 *  Name:
 *  Date:
 *  Description:
 **************************************************************************** */

import edu.princeton.cs.algs4.StdOut;

public class LinkedQueueOfStringsWithDummyNode {

    private class Node {
        String item;
        Node next;

        public Node(String str) {
            item = str;
        }
        public Node() {}
    }

    private Node first, last;
    private int size;

    public LinkedQueueOfStringsWithDummyNode() {
        first = new Node();
        last = first;
    }

    public void enqueue(String item) {
        Node node = new Node(item);
        last.next = node;
        last = node;
        size++;
    }

    public String dequeue() {
        String item = first.next.item;
        first.next = first.next.next;
        // Special case for empty queue
        if (isEmpty()) {
            last = first;
        }
        size--;
        return item;
    }

    boolean isEmpty() {
        return first.next == null;
    }

    int size() {
        return size;
    }

    // test client (optional)
    public static void main(String[] args) {
        LinkedQueueOfStringsWithDummyNode queue = new LinkedQueueOfStringsWithDummyNode();

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
