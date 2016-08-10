package main
import (
    "fmt"
    "os"
    "image"
    "image/color"
    "image/draw"
    "image/jpeg"
)
func main() {
    f1, err := os.Open("1.jpg")
    if err != nil {
        panic(err)
    }
    defer f1.Close()
    f2, err := os.Open("2.jpg")
    if err != nil {
        panic(err)
    }
    defer f2.Close()
    f3, err := os.Create("3.jpg")
    if err != nil {
        panic(err)
    }
    defer f3.Close()
   
    m1, err := jpeg.Decode(f1)
    if err != nil {
        panic(err)
    }
    bounds := m1.Bounds()
    m2, err := jpeg.Decode(f2)
    if err != nil {
        panic(err)
    }
    m := image.NewRGBA(bounds)
    white := color.RGBA{255, 255, 255, 255}
    draw.Draw(m, bounds, &image.Uniform{white}, image.ZP, draw.Src)
    draw.Draw(m, bounds, m1, image.ZP, draw.Src)
    draw.Draw(m, image.Rect(100, 200, 300, 600), m2, image.Pt(250, 60), draw.Src)
    err = jpeg.Encode(f3, m, &jpeg.Options{90})
    if err != nil {
        panic(err)
    }
    fmt.Printf("ok\n")
}
