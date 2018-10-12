import edu.princeton.cs.algs4.In;
import edu.princeton.cs.algs4.StdOut;

public class QuickFindUF {
    private int[] id;

    public QuickFindUF(int N) {
        id = new int[N];
        for (int i = 0; i < N; i++) {
            id[i] = i;
        }
    }

    public boolean connected(int p, int q) {
        return id[p] == id[q];
    }

    public void union(int p, int q) {
        int pid = id[p];
        int qid = id[q];
        for (int i = 0; i < id.length; i++) {
           if (id[i] == pid) {
               id[i] = qid;
           }
        }
    }

    public static void main(String[] args) {
        In in = new In("/Users/l.cao/programming/java/algs4-data/tinyUF.txt");
        int N = in.readInt();

        QuickFindUF uf = new QuickFindUF(N);

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
