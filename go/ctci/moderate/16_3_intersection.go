package main

import (
	"errors"
	"fmt"
)

// Given two straight line segments (represented as a start point and an end point)
// compute the point of intersection, if any.
// Slope can only be the same if they're the same line.

// Point is a point on the grid
type Point struct {
	X float64 // could also do floats here..
	Y float64
}

// NewPoint ...
func NewPoint(x, y float64) Point {
	return Point{X: x, Y: y}
}

// Line is a line
type Line struct {
	Slope      float64
	YIntercept float64
}

// NewLine ...
func NewLine(start, end Point) Line {
	slope := (end.Y - start.Y) / (end.X - start.X)
	yint := start.Y - slope*start.X
	return Line{Slope: slope, YIntercept: yint}
}

func swap(pt1, pt2 *Point) {
	pt2.X, pt2.Y, pt1.X, pt1.Y = pt1.X, pt1.Y, pt2.X, pt2.Y
}

func intersection(start1, end1, start2, end2 Point) (Point, error) {
	if start1.X > end1.X {
		swap(&start1, &end1)
	}
	if start2.X > end2.X {
		swap(&start2, &end2)
	}
	if start1.X > start2.X {
		swap(&start1, &start2)
		swap(&end1, &end2)
	}

	line1, line2 := NewLine(start1, end1), NewLine(start2, end2)
	if line1.Slope == line2.Slope {
		return Point{}, errors.New("The lines do not intersect")

	}
	x := (line2.YIntercept - line1.YIntercept) / (line1.Slope - line2.Slope)
	y := eval(line1, x)
	return Point{x, y}, nil
}

func eval(l Line, x float64) float64 {
	return l.Slope*x + l.YIntercept
}

func main() {
	pt1, pt2 := NewPoint(2.0, 3.0), NewPoint(4.0, 2.0)
	pt3, pt4 := NewPoint(2.0, 3.0), NewPoint(4.0, 2.0)
	if result, err := intersection(pt1, pt2, pt3, pt4); err == nil {
		fmt.Println(result)
	} else {
		fmt.Println("The lines do not intersect")
	}
}
