/* *****************************************************************************
 *  Name:
 *  Date:
 *  Description:
 **************************************************************************** */

import edu.princeton.cs.algs4.StdOut;

public class ResizingArrayStackOfStrings {

    private String[] items;
    private int N = 0;

    public ResizingArrayStackOfStrings() {
        items = new String[1];
    }

    void push(String item) {
        if (items.length == N) {
            resize(items.length*2);
        }

        items[N++] = item;
    }

    private void resize(int size) {
        String[] copy = new String[size];
        // use N, don't use items.length or size
        for (int i = 0; i < N; i++) {
            copy[i] = items[i];
        }
        items = copy;
    }

    String pop() {
        // Empty
        String temp = items[--N];
        items[N] = null;
        if (N>0 && N == items.length/4) {
            resize(items.length/2);
        }
        return temp;
    }

    boolean isEmpty() {
        return N <= 0;
    }

    int size() {
        return N;
    }

    // test client (optional)
    public static void main(String[] args) {
        ResizingArrayStackOfStrings stack = new ResizingArrayStackOfStrings();


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
