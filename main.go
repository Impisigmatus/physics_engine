package main

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"time"

	"github.com/Impisigmatus/physics_engine/internal/geometry"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		CallerPrettyfier: func(frame *runtime.Frame) (string, string) {
			file := frame.File[len(path.Dir(os.Args[0]))+1:]
			line := frame.Line
			return "", fmt.Sprintf("%s:%d", file, line)
		},
	})
}

func main() {
	var begin time.Time

	edgeA := geometry.NewLine(&geometry.Point{X: 0, Y: 0}, &geometry.Point{X: 5, Y: 5})
	edgeB := geometry.NewLine(&geometry.Point{X: 0, Y: 3}, &geometry.Point{X: 5, Y: 2})

	begin = time.Now()
	if point := geometry.LineCross(edgeA, edgeB); point != nil {
		logrus.Infof("Cross[%s]: %v", time.Since(begin), *point)
	} else {
		logrus.Infof("Cross[%s]: No", time.Since(begin))
	}

	begin = time.Now()
	if point := geometry.LineIntersection(edgeA, edgeB); point != nil {
		logrus.Infof("Intersects[%s]: %v", time.Since(begin), *point)
	} else {
		logrus.Infof("Intersects[%s]: No", time.Since(begin))
	}
}
