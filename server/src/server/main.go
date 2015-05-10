package main

import (
    "fmt"
    "server/Geometry"
)

func main() {
    fmt.Println("Hello, World!")
    p1 := Geometry.Point{X:0, Y:0}
    p2 := Geometry.Point{X:10, Y:10}

    r1 := Geometry.Rect{Coords:p1, W:5, H:5}
    r2 := Geometry.Rect{Coords:p2, W:5, H:5}

    fmt.Println(Geometry.SayHello())
    fmt.Println(Geometry.CollisionRR(r1, r2))
    fmt.Println(Geometry.LineCoefficient(p1, p2))
}