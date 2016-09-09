package main

import (
	. "github.com/smartystreets/goconvey"
)

func main() {
	Convey("The value should be greater by one", func() {
		So(2, ShouldEqual, 2)
	})
}
