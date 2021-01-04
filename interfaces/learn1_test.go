package interfaces

import (
	"testing"
)

type DataTest struct {
	react   React
	ErrorMessageArea string
	ErrorMessagePerimeter string
	ExpectationArea  float64
	ExpectationPerimeter float64
}

func TestReactArea(t *testing.T) {
	data := []DataTest{
		DataTest{
			react: React{
				1.0,
				2.0,
			},
			ExpectationArea:  2.0,
			ErrorMessageArea: "the result is not 2.0",
			ExpectationPerimeter: 4.0,
			ErrorMessagePerimeter: "the result is not 4.0",

		},
		DataTest{
			react: React{
				-1.0,
				2.0,
			},
			ExpectationArea:  -2.0,
			ErrorMessageArea: "the result is not -2.0",
			ExpectationPerimeter: -4.0,
			ErrorMessagePerimeter: "the result is not -4.0",
		},
		DataTest{
			react: React{
				0.0,
				0.0,
			},
			ExpectationArea:  0.0,
			ErrorMessageArea: "the result is not 0.0",
			ExpectationPerimeter: 0.0,
			ErrorMessagePerimeter: "the result is not 0.0",
		},
	}
	for _,datas := range data{
		if datas.react.Area() != datas.ExpectationArea {
			t.Error(datas.ErrorMessageArea)
		}
	}



}
