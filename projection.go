package main

type Projection struct {
	Original    Rectangle
	Destination Rectangle
}

func (projection Projection) Transform(point Point) (tranformed Point) {
	//TODO for future perhaps cache something
	tranformed.x = (point.x-projection.Original.Min.x)/(projection.Original.Max.x-projection.Original.Min.x)*(projection.Destination.Max.x-projection.Destination.Min.x) + projection.Destination.Min.x
	tranformed.y = (point.y-projection.Original.Min.y)/(projection.Original.Max.y-projection.Original.Min.y)*(projection.Destination.Max.y-projection.Destination.Min.y) + projection.Destination.Min.y

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
