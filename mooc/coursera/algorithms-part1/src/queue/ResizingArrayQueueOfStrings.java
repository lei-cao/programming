/* *****************************************************************************
 *  Name:
 *  Date:
 *  Description:
 **************************************************************************** */

import edu.princeton.cs.algs4.StdOut;

public class ResizingArrayQueueOfStrings {

    private String[] items;
    private int first, last;
    private int capacity = 1;

    public ResizingArrayQueueOfStrings() {
        items = new String[capacity];
        first = last = 0;
    }

    public void enqueue(String item) {
        if (last == capacity) {
            capacity = capacity * 2;
            resize(capacity);
        }
        items[last++] = item;
    }

    public String dequeue() {
        String item = items[first];
        items[first] = null;
        first++;
        return item;
    }

    public boolean isEmpty() {
        return first == last;
    }

    private void resize(int size) {
        StdOut.println("Resizing");
        StdOut.println(size);
        String[] copy;
        copy = new String[size];
        int length = last - first;
        for (int i = 0; i < last - first; i++) {
            copy[i] = items[first+i];
        }
        items = copy;
        first = 0;
        last = length;
    }

    int size() {
        return items.length;
    }

    // test client (optional)
    public static void main(String[] args) {
        ResizingArrayQueueOfStrings queue = new ResizingArrayQueueOfStrings();

        StdOut.println(queue.isEmpty());
        queue.enqueue("1");
        queue.enqueue("2");
        queue.enqueue("3");
        queue.enqueue("4");
        queue.enqueue("5");
        queue.enqueue("6");
        StdOut.println(queue.isEmpty());
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
        queue.enqueue("5");
        queue.enqueue("6");
        StdOut.println(queue.dequeue());
        StdOut.println(queue.dequeue());
        StdOut.println(queue.dequeue());

        StdOut.println(queue.dequeue());
        StdOut.println(queue.dequeue());
        StdOut.println(queue.dequeue());

    }

}
