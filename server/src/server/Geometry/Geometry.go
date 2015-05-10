package Geometry

import "math"

type Point struct {
    X, Y float64
}

type Rect struct {
    Coords Point
    W, H float64
}

type Vector struct {
    Coords Point
    Celerity Point
    Direction float64
}

func Round(f float64) float64 {
    return math.Floor(f + .5)
}

func RoundPlus(f float64, places int) (float64) {
    shift := math.Pow(10, float64(places))
    return Round(f * shift) / shift
}

func CollisionRR(r1 Rect, r2 Rect) bool {
    if r1.Coords.X + r1.W < r2.Coords.X  || r1.Coords.X > r2.Coords.X + r2.W {
        if r1.Coords.Y + r1.H < r2.Coords.Y || r1.Coords.Y > r2.Coords.Y + r2.H {
            return false;
        }
    }
    return true;
}

func CollisionPR(p Point, r Rect) bool {
    if p.X < r.Coords.X || p.X > r.Coords.X + r.W {
        if p.Y < r.Coords.Y || p.Y > r.Coords.Y + r.H {
            return false;
        }
    }
    return true;
}

func GetPointFromAngle(angle float64) Point {
    point := Point{X: math.Cos(angle), Y: math.Sin(angle)}
    return Point{X: RoundPlus(point.X, 3), Y: RoundPlus(point.Y, 3)}
}

func ApplyCelerity(vector Vector) Vector{
    vector.Coords.X = vector.Coords.X + vector.Celerity.X
    vector.Coords.Y = vector.Coords.Y + vector.Celerity.Y
    return vector
}

func ApplyBounce(vector Vector, point Point) Vector {
    return vector
}

func LineCoefficient(p1 Point, p2 Point) float64 {
    return (p2.Y - p1.Y) / (p2.X - p1.X) 
}

