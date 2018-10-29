import edu.princeton.cs.algs4.ResizingArrayQueue;

import java.util.Comparator;

public class BruteCollinearPoints {
  private int numberOfSegments;
  private final ResizingArrayQueue<LineSegment> segmentList;

  // finds all line segments containing 4 points
  public BruteCollinearPoints(Point[] points) {
    if (points == null) {
      throw (new IllegalArgumentException());
    }

    for (int i = 0; i < points.length; i++) {
      if (points[i] == null) {
        throw (new IllegalArgumentException());
      }
      for (int j = i + 1; j < points.length; j++) {
        if (points[i].compareTo(points[j]) == 0) {
          throw (new IllegalArgumentException());
        }
      }
    }

    segmentList = new ResizingArrayQueue<>();
    for (int i = 0; i < points.length; i++) {
      Comparator<Point> slopeOrder = points[i].slopeOrder();
      for (int j = 0; j < points.length; j++) {
        if (points[i].compareTo(points[j]) >= 0) {
          continue;
        }
        for (int k = 0; k < points.length; k++) {
          if (points[j].compareTo(points[k]) >= 0) {
            continue;
          }
          if (slopeOrder.compare(points[j], points[k]) != 0) {
            continue;
          }
          for (int m = 0; m < points.length; m++) {
            if (points[k].compareTo(points[m]) >= 0) {
              continue;
            }
            if (slopeOrder.compare(points[k], points[m]) == 0) {
              segmentList.enqueue(new LineSegment(points[i], points[m]));
              numberOfSegments++;
            }
          }
        }
      }
    }
  }

  // the number of line segments
  public int numberOfSegments() {
    return numberOfSegments;
  }

  // the line segments
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
