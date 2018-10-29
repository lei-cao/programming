import edu.princeton.cs.algs4.Queue;

import java.util.Arrays;
import java.util.Comparator;

public class FastCollinearPoints {
  private int numberOfSegments = 0;
  private Queue<LineSegment> segmentList;

  /**
   * Finds all line segments containing 4 points.
   *
   * @param points Points
   */
  public FastCollinearPoints(Point[] points) {
    // Input check
    if (points == null) {
      throw (new IllegalArgumentException());
    }

    segmentList = new Queue<>();

    if (points.length == 0) {
      return;
    }


    Point[] pointList = new Point[points.length];
    // Input check
    for (int i = 0; i < points.length; i++) {
      if (points[i] == null) {
        throw (new IllegalArgumentException());
      }
      for (int j = i + 1; j < points.length; j++) {
        if (points[i].compareTo(points[j]) == 0) {
          throw (new IllegalArgumentException());
        }
      }

      pointList[i] = points[i];
    }

    // create list out of origin array, remove the process element
    for (int i = 0; i < points.length; i++) {
      process(points[i], pointList);
    }
  }

  private void process(Point point, Point[] points) {
    Comparator<Point> comparator = point.slopeOrder();
    Arrays.sort(points, comparator);

    int start;
    int end;
    for (start = 0; start < points.length; start = end) {
      int count = 1;
      for (end = start + 1; end < points.length; end++) {
        if (comparator.compare(points[start], points[end]) == 0) {
          count++;
          continue;
        }
        break;
      }
      if (count > 2) {
        Point[] seg = new Point[count + 1];
        seg[0] = point;
        for (int i = 1; i < count + 1; i++) {
          seg[i] = points[start + i - 1];
        }
        Arrays.sort(seg);
        if (point.compareTo(seg[0]) == 0) {
          segmentList.enqueue(new LineSegment(seg[0], seg[seg.length - 1]));
          numberOfSegments++;
        }
      }
    }
  }

  // the number of line segments
  public int numberOfSegments() {
    return numberOfSegments;
  }

  /** All line segments. @ return */
  public LineSegment[] segments() {
    LineSegment[] seg = new LineSegment[segmentList.size()];
    int i = 0;
    for (LineSegment ls : segmentList) {
      seg[i] = ls;
      i++;
    }
    return seg;
  }
}
