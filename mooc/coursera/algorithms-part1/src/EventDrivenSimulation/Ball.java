import edu.princeton.cs.algs4.StdDraw;
import edu.princeton.cs.algs4.StdRandom;

public class Ball {
  private double rx, ry;
  private double vx, vy;
  private final double radius = 0.3;
  private double width = 50;
  private double height = 50;

  public Ball() {
    rx = StdRandom.uniform(0, 50);
    ry = StdRandom.uniform(0, 50);
    vx = 0.5;
    vy = 0.5;
  }

  public void move(double dt) {
    if (rx + vx * dt < radius || rx + vx * dt > 50.0 - radius) {
      vx = -vx;
    }

    if (ry + vy * dt < radius || ry + vy * dt > 50.0 - radius) {
      vy = -vy;
    }

    rx = rx + vx * dt;
    ry = ry + vy * dt;
  }

  public void draw() {
    StdDraw.filledCircle(rx, ry, radius);
  }

  public double timeToHit(Ball that) {
    double t = 0;
  }

  public double timeToHitVerticalWall() {
    return (50 - ry) / vy;
  }

  public double timeToHitHorizontalWall() {
    return (50 - rx) / vx;
  }

  public void bounceOff(Ball that) {}

  public void bounceOffVerticalWall() {}

  public void bounceOffHorizontalWall() {}

  private double getDistance(Ball that) {
    double dx = rx - that.rx;
    double dy = ry - that.ry;

    return Math.sqrt(Math.pow(dx, 2) + Math.pow(dy, 2));
  }
}
