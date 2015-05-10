package Geometry

type Point struct {
    X, Y float64
}

type Rect struct {
    Coords Point
    W, H float64
}

type Celerity struct {
    Ratio Point
}

type Vector struct {
    Coords Point
    Celerity float64
    Direction float64
}

func SayHello() string{
    return "Hello world"
}

func CollisionRR(r1 Rect, r2 Rect) bool{
    if r1.Coords.X + r1.W < r2.Coords.X  || r1.Coords.X > r2.Coords.X + r2.W {
        if r1.Coords.Y + r1.H < r2.Coords.Y || r1.Coords.Y > r2.Coords.Y + r2.H {
            return false;
        }
    }
    return true;
}

func CollisionPR(p Point, r Rect) bool{
    if p.X < r.Coords.X || p.X > r.Coords.X + r.W {
        if p.Y < r.Coords.Y || p.Y > r.Coords.Y + r.H {
            return false;
        }
    }
    return true;
}

func ApplyCelerity(vector Vector) Vector{
    return vector
}

func LineCoefficient(p1 Point, p2 Point) float64{
    return (p1.Y - p2.Y) / (p1.X - p2.X) 
}