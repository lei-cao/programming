import edu.princeton.cs.algs4.StdDraw;

public class BouncingBalls {
  public static void main(String[] args) {
    int N = Integer.parseInt(args[0]);

    Ball[] balls = new Ball[N];

    StdDraw.enableDoubleBuffering();

    for (int i = 0; i < N; i++) {
      balls[i] = new Ball();
    }

    while (true) {
      StdDraw.setXscale(-0.2 , 50.2 );
      StdDraw.setYscale(-0.2 , 50.2 );   // leave a border to write text
//      StdDraw.filledSquare(N / 1.0, N / 1.0, N / 1.0);
      StdDraw.clear();

      for (int i = 0; i < N; i++) {
        balls[i].move(0.5);
        balls[i].draw();
      }

      StdDraw.show();
      StdDraw.pause(10);
    }
  }
}
