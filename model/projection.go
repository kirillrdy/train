package model

type Projection struct {
	Original    Rectangle
	Destination Rectangle
}

func (projection Projection) Translate(point Point) (tranformed Point) {
	//TODO for future perhaps cache something
	tranformed.X = (point.X-projection.Original.Min.X)/(projection.Original.Max.X-projection.Original.Min.X)*(projection.Destination.Max.X-projection.Destination.Min.X) + projection.Destination.Min.X
	tranformed.Y = (point.Y-projection.Original.Min.Y)/(projection.Original.Max.Y-projection.Original.Min.Y)*(projection.Destination.Max.Y-projection.Destination.Min.Y) + projection.Destination.Min.Y

	return tranformed
}

//func main() {
//	scren := Rectangle{Point{0, 0}, Point{640, 480}}
//	fake_world := Rectangle{Point{-10, -10}, Point{10, 10}}
//	projection := Projection{Original: fake_world, Destination: scren}
//
//	want_to_draw_point := Point{-10, -10}
//
//	fmt.Println(projection.Transform(want_to_draw_point))
//}
