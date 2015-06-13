package lib

import (
	"log"
	"time"

	"github.com/kirillrdy/nadeshiko"
	"github.com/kirillrdy/nadeshiko/html"
	"github.com/kirillrdy/train/melbourne"
	"github.com/kirillrdy/train/model"
	"github.com/sparkymat/webdsl/css"
)

var projection model.Projection

var svgMap css.Id = "svg-map"

func RenderMap(document *nadeshiko.Document, projection model.Projection) {
	now := time.Now()
	svgPaths := melbourne.RailMap(svgMap, projection)
	log.Printf("svg gen took %s", time.Since(now))

	now = time.Now()
	document.JQuery(svgMap).Empty()
	for _, path := range svgPaths[0:1] {
		document.JQuery(svgMap).HTML(path)
	}

	log.Printf("append took %s", time.Since(now))

	log.Printf("size %d", len(svgPaths))
}

func Handler(document *nadeshiko.Document) {

	screen := model.Screen{Width: 507, Height: 324}

	projection := model.Projection{Original: melbourne.Boundaries, Destination: screen.ToRectangle()}

	div := html.Div().Id(svgMap)

	svg_element := html.Svg().Id(svgMap).WidthFloat(projection.Destination.Max.X).HeightFloat(projection.Destination.Max.Y)
	document.JQuery(css.Body).Append(svg_element)

	RenderMap(document, projection)

	var zoomInButton css.Id = "zoom-in-button"
	document.JQuery(css.Body).Append(html.Button().Id(zoomInButton).Text("Zoom in"))

	document.JQuery(zoomInButton).Click(func() {
		projection = projection.Pan(1, 1)
		RenderMap(document, projection)
	})

	go func() {
	}()

	//button := `<input type="button" id="add" value="Add Train">`
	//conneciton.JQuery("body").Append(button)

	//conneciton.JQuery("#add").Click(func() {
	//	limit += 1
	//	conneciton.JQuery("svg").Remove()
	//	activity.AddMap(conneciton, limit)
	//})

	//go func() {
	//	for {
	//		time.Sleep(10 * time.Millisecond)
	//		for _, train := range activity.trains {
	//			train.OneFrame()
	//		}

	//		conneciton.StartBuffer()

	//		for _, train := range activity.trains {
	//			train.Draw(conneciton)
	//		}

	//		conneciton.FlushBuffer()

	//		var non_dead_head []*Train

	//		for _, train := range activity.trains {
	//			if train.ShouldBeRemoved() {
	//				train.RemoveFromPage(conneciton)
	//			} else {
	//				non_dead_head = append(non_dead_head, train)
	//			}
	//		}
	//		activity.trains = non_dead_head
	//	}
	//}()

}

//func (activity *MainActivity) addAtrain(path []Point, conneciton *nadeshiko.Connection) {
//	for i := 0; i < 100; i++ {
//		activity.addTrainToPath(path, conneciton)
//		time.Sleep(120 * time.Millisecond)
//	}
//}
//
//func (activity *MainActivity) addTrainToPath(path []Point, conneciton *nadeshiko.Connection) {
//
//	train := NewTrain(path)
//	train.AppendToPage(conneciton)
//	activity.trains = append(activity.trains, &train)
//
//}
