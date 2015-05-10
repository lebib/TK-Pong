package Geometry
import "math"

////////////////////////////////////////////////////////////////////////////////

type Point struct  {
    X, Y float64
}

type Rect struct {
    Point
    W, H float64
}

func (r1 *Rect) CollisionRR(r2 Rect) bool {
    if r1.X + r1.W < r2.X || r1.X > r2.X + r2.W {
        if r1.Y + r1.H < r2.Y || r1.Y > r2.Y + r2.H {
            return false;
        }
    } 
    return true
}

func (r *Rect) CollisionPR(p Point) bool {
    if p.X < r.X || p.X > r.X + r.W {
        if p.Y < r.Y || p.Y > r.Y + r.H {
            return false;
        }
    }
    return true;
}

////////////////////////////////////////////////////////////////////////////////

func Round(val float64, precision int) float64 {
    buffer := math.Pow(10, float64(precision))
    return math.Floor(val * buffer + 0.5) / buffer
}

func GetPointFromAngle(angle float64, precision int) Point {
    return Point{ Round(math.Cos(angle), precision),
        Round(math.Sin(angle), precision),
    }
}

func LineCoefficient(p1, p2 Point) float64 {
    return (p2.Y - p1.Y) / (p2.X - p1.X)
}

////////////////////////////////////////////////////////////////////////////////


// func GetPointFromAngle(angle float64) Point {
//     point := Point{X: math.Cos(angle), Y: math.Sin(angle)}
//     return Point{X: RoundPlus(point.X, 3), Y: RoundPlus(point.Y, 3)}
// }

// func ApplyCelerity(vector Vector) Vector{
//     vector.Coords.X = vector.Coords.X + vector.Celerity.X
//     vector.Coords.Y = vector.Coords.Y + vector.Celerity.Y
//     return vector
// }

// func ApplyBounce(vector Vector, point Point) Vector {
//     return vector
// }

// func LineCoefficient(p1 Point, p2 Point) float64 {
//     return (p2.Y - p1.Y) / (p2.X - p1.X) 
// }