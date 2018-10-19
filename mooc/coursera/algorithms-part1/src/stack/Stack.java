import edu.princeton.cs.algs4.StdOut;

import java.util.Iterator;

public class Stack<Item> implements Iterable<Item> {
    private Node first;

    @Override
    public Iterator<Item> iterator() {
        return new ListIterator();
    }

    private class ListIterator implements Iterator<Item> {
        private Node current = first;

        @Override
        public boolean hasNext() {
            return current != null;
        }

        @Override
        public Item next() {
            Item item = current.item;
            current = current.next;
            return item;
        }

        @Override
        public void remove() {
            /*
            Not Support
             */
        }
    }

    private class Node {
        private Item item;
        private Node next;
    }

    public void push(Item item) {
        Node node = new Node();
        node.item = item;
        node.next = first;
        first = node;
    }

    public Item pop() {
        Node node = first;
        first = first.next;
        return node.item;
    }

    public boolean isEmpty() {
        return first == null;
    }

    public static void main(String[] args) {
        Stack<String> stack = new Stack<>();
        stack.push("1");
        stack.push("2");
        stack.push("3");
        StdOut.println(stack.pop());
        StdOut.println(stack.pop());
        StdOut.println(stack.pop());

        stack.push("a");
        stack.push("b");
        stack.push("c");

        Iterator<String> i = stack.iterator();
        while(i.hasNext()) {
            String s = i.next();
            StdOut.println(s);
        }

    }
}

