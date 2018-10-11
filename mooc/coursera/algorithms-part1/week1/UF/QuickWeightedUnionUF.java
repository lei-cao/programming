import edu.princeton.cs.algs4.In;
import edu.princeton.cs.algs4.StdOut;

public class QuickWeightedUnionUF {
    private int[] id;
    private int[] size;

    public QuickWeightedUnionUF(int N) {
        id = new int[N];
        size = new int[N];
        for (int i = 0; i < N; i++) {
            id[i] = i;
            size[i] = 1;
        }
    }

    public boolean connected(int p, int q) {
        return root(p) == root(q);
    }


    public void union(int p, int q) {
        int rootP = root(p);
        int rootQ = root(q);
        if (rootP == rootQ) {
            return;
        }
        if (size[rootP] > size[rootQ]) {
            id[rootP] = rootQ;
            size[rootP] = size[rootP] + size[rootQ];
        } else {
            id[rootQ] = rootP;
            size[rootQ] = size[rootP] + size[rootQ];
        }
    }

    private int root(int i) {
        while (i != id[i]) {
            id[i] = id[id[i]];
            i = id[i];
        }
        return i;
    }

    public static void main(String[] args) {
        In in = new In("/Users/l.cao/programming/java/algs4-data/tinyUF.txt");
        int N = in.readInt();

        QuickWeightedUnionUF uf = new QuickWeightedUnionUF(N);

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
