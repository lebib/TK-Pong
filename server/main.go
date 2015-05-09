package main

import ("fmt")

type Rect struct {
    x, y, w, h int32
}

type Point struct {
    x, y, alpha int32
}

type Vector struct {
    x, y, c, d int32
}

func detectCollisionR2R(r1 Rect, r2 Rect) bool {
    if r1.x + r1.w <= r2.x || r1.x >= r2.x + r2.w {
        if r1.y + r1.h <= r2.y || r1.y >= r2.y + r2.h {
            return true;
        }
    }
    return false;
}

func bounce(v Vector) void {

}

func main() {
    fmt.Println("hello world")
    r1 := Rect{x:0, y:0, w:10, h:10}
    r2 := Rect{x:20, y:20, w:10, h:10}
    test := detectCollisionR2R(r1, r2)
    fmt.Println(test)
}
