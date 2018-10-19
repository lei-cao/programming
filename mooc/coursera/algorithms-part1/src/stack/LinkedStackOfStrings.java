/* *****************************************************************************
 *  Name:
 *  Date:
 *  Description:
 **************************************************************************** */

import edu.princeton.cs.algs4.StdOut;

public class LinkedStackOfStrings {

    private class Node {
        String item;
        Node next;

        public Node(String str) {
            item = str;
        }
    }

    private Node first = null;
    private int size;

    public LinkedStackOfStrings() {
    }

    void push(String item) {
        Node oldFirst = first;
        first = new Node(item);
        first.next = oldFirst;
        size++;
    }

    String pop() {
        String item = first.item;
        first = first.next;
        size--;
        return item;
    }

    boolean isEmpty() {
        return first == null;
    }

    int size() {
        return size;
    }

    // test client (optional)
    public static void main(String[] args) {
        LinkedStackOfStrings stack = new LinkedStackOfStrings();

        StdOut.println(stack.size());
        stack.push("a");
        stack.push("b");
        stack.push("c");
        StdOut.println(stack.size());
        StdOut.println(stack.pop());
        StdOut.println(stack.pop());
        StdOut.println(stack.pop());
        StdOut.println(stack.isEmpty());
        stack.push("a");
        stack.push("b");
        stack.push("c");
        StdOut.println(stack.pop());
        StdOut.println(stack.pop());
        StdOut.println(stack.pop());

    }

}
