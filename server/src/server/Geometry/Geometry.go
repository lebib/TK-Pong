package Geometry

type Point struct {
    X, Y float64
}

type Rect struct {
    C Point
    W, H float64
}

func SayHello() string{
    return "Hello world"
}

func CollisionRR(r1 Rect, r2 Rect) bool{
    if r1.C.X + r1.W < r2.C.X  || r1.C.X > r2.C.X + r2.W {
        if r1.C.Y + r1.H < r2.C.Y || r1.C.Y > r2.C.Y + r2.H {
            return false;
        }
    }
    return true;
}

func CollisionPR(p Point, r Rect) bool{
    if(p.X < r.C.X && p.X || r.C.X + r.W) {
        if(p.Y < r.C.Y || p.Y > r.C.Y + h)
        {
            return false;
        }
    }
    return true;
}