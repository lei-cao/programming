import edu.princeton.cs.algs4.StdOut;

public class FixedCapacityArrayStack<Item> {
    private Item[] s;
    private int N = 0;

    public FixedCapacityArrayStack(int capacity) {
        s = (Item[]) new Object[capacity];
    }

    public Item pop() {
        return s[--N];
    }

    public void push(Item item) {
        s[N++] = item;
    }

    public static void main(String[] args) {
        FixedCapacityArrayStack<String> stack = new FixedCapacityArrayStack<>(4);
        stack.push("1");
        stack.push("2");
        stack.push("3");
        stack.push("3");
        StdOut.println(stack.pop());
        StdOut.println(stack.pop());
        StdOut.println(stack.pop());
        StdOut.println(stack.pop());

    }
}
