package main

import (
    "fmt"
    "math"
    "server/Geometry"
)

func main() {
    // p1 := Geometry.Point{X:0, Y:0}
    // p2 := Geometry.Point{X:10, Y:10}
    // r1 := Geometry.Rect{Coords:p1, W:5, H:5}
    // r2 := Geometry.Rect{Coords:p2, W:5, H:5}
    p3 := Geometry.Point{X:5, Y:5}
    p4 := Geometry.GetPointFromAngle(-math.Pi / 2)
    // p5 := Geometry.GetPointFromAngle(math.Pi)
    // p6 := Geometry.GetPointFromAngle(-math.Pi / 2)
    // p7 := Geometry.GetPointFromAngle(-math.Pi)

    fmt.Println(p3.X, p3.Y)
    fmt.Println(p4.X, p4.Y)
    // fmt.Println(p5.X, p5.Y)
    // fmt.Println(p6.X, p6.Y)
    // fmt.Println(p7.X, p7.Y)
    p8 := Geometry.Point{X:p3.X+p4.X, Y:p3.Y+p4.Y}
    fmt.Println(p8.X, p8.Y)
    fmt.Println(Geometry.LineCoefficient(p3, p8))
    if math.IsInf(Geometry.LineCoefficient(p3, p8), 0) {
        fmt.Println("Hello, World!")
    }
}