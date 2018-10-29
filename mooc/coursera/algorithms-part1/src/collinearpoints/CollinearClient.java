import edu.princeton.cs.algs4.In;
import edu.princeton.cs.algs4.StdDraw;
import edu.princeton.cs.algs4.StdOut;

public class CollinearClient {
  /** Unit tests the Point data type. */
  public static void main(String[] args) {
    /* YOUR CODE HERE */

    StdDraw.clear();
    StdDraw.enableDoubleBuffering();
    StdDraw.setXscale(0, 32768);
    StdDraw.setYscale(0, 32768);

    StdDraw.setPenRadius(0.003);
    StdDraw.setPenColor(StdDraw.BLACK);
    In in = null;

    //                in = new In("./collinear/equidistant.txt");
    //                in = new In("./collinear/grid4x4.txt");
    //        in = new In("./collinear/horizontal5.txt");
    //        in = new In("./collinear/rs1423.txt");
    //    in = new In("./collinear/input8.txt");
//                in = new In("./collinear/input8.txt");
    //            in = new In("./collinear/input50.txt");
    //        in = new In("./collinear/input80.txt");
    //        in = new In("./collinear/input400.txt");
    //        in = new In("./collinear/input1000.txt");
    //        in = new In("./collinear/input2000.txt");
    //        in = new In("./collinear/input3000.txt");
    //        in = new In("./collinear/input4000.txt");
    //        in = new In("./collinear/input5000.txt");
    //        in = new In("./collinear/input6000.txt");
    //        in = new In("./collinear/input8000.txt");
    //        in = new In("./collinear/input10000.txt");
            in = new In("./collinear/mystery10089.txt");
    //        in = new In("./collinear/random23.txt");
    //        in = new In("./collinear/random38.txt");
    //        in = new In("./collinear/random152.txt");
//      in = new In("./collinear/kw1260.txt");
//                in = new In("./collinear/rs1423.txt");
//    in = new In("./collinear/vertical.txt");
//            in = new In("./collinear/vertical5.txt");
    //            in = new In("./collinear/vertical100.txt");

    int n = in.readInt();
    Point[] ps = new Point[n];
    for (int i = 0; i < n; i++) {
      int x = in.readInt();
      int y = in.readInt();
      ps[i] = new Point(x, y);
      StdDraw.point(x, y);
      StdDraw.show();
    }

    StdDraw.setPenRadius(0.001);
    StdDraw.setPenColor(StdDraw.RED);

    /*
    BruteCollinearPoints collinear = new BruteCollinearPoints(ps);
    LineSegment[] ss = collinear.segments();
    for (int i = 0; i < ss.length; i++) {
        StdOut.println(ss[i]);
        ss[i].draw();
        StdDraw.show();
    }
    StdOut.println("xxxxxxxxx");
    */

    FastCollinearPoints collinear2 = new FastCollinearPoints(ps);
    LineSegment[] ss2 = collinear2.segments();
    for (int i = 0; i < ss2.length; i++) {
      StdOut.println(ss2[i]);
      ss2[i].draw();
      StdDraw.show();
    }
    StdOut.println("xxxxxxxxx");
    //        StdOut.println(ss.length);
    StdOut.println(ss2.length);
  }
}
