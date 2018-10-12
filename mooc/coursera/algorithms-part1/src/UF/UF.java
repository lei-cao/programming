/* *****************************************************************************
 *  Name:
 *  Date:
 *  Description:
 **************************************************************************** */

import edu.princeton.cs.algs4.In;
import edu.princeton.cs.algs4.StdOut;

public class UF {

    // initialize union-find data structure with N objects (0 to N-1)
    public UF(int N) {
    }

    // add connection between p and q
    public void union(int p, int q) {

    }

    // are p and q in the same component?
    public boolean connected(int p, int q) {
        return false;
    }

    // component identifier for p (0 to N-1)
    public int find(int p) {
        return 0;
    }

    // number of components
    public int count() {
        return 0;
    }

    // test client (optional)
    public static void main(String[] args) {
        In in = new In(args[0]);
        int N = in.readInt();

        UF uf = new UF(N);
        while (!in.isEmpty()) {
            int p = in.readInt();
            int q = in.readInt();
            if (uf.connected(p, q)) continue;
            uf.union(p, q);
            StdOut.println(p + " " + q);
        }
    }

}

