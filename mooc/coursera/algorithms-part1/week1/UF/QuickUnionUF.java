import edu.princeton.cs.algs4.In;
import edu.princeton.cs.algs4.StdOut;

public class QuickUnionUF {
    private int[] id;

    public QuickUnionUF(int N) {
        id = new int[N];
        for (int i = 0; i < N; i++) {
            id[i] = i;
        }
    }

    public boolean connected(int p, int q) {
        return root(p) == root(q);
    }


    public void union(int p, int q) {
        id[root(p)] = root(q);
    }

    private int root(int i) {
        while(i != id[i]) {
            i = id[i];
        }
        return i;
    }

    public static void main(String[] args) {
        In in = new In("/Users/l.cao/programming/java/algs4-data/tinyUF.txt");
        int N = in.readInt();

        QuickUnionUF uf = new QuickUnionUF(N);

        while (!in.isEmpty()) {
            int p = in.readInt();
            int q = in.readInt();
            if (uf.connected(p, q)) {
                continue;
            }
            uf.union(p, q);
        }

        for (int i = 0; i < uf.id.length; i++) {
            StdOut.println(uf.id[i]);
        }
    }
}
