package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSpec(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Given some integer with a starting value", t, func() {
		x := 1

		Convey("When the integer is incremented", func() {
			y := x + 2

			Convey("The value should be greater by one", func() {
				So(y, ShouldEqual, 2)
			})
		})
	})
}
