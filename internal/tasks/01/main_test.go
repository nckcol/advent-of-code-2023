package main

import (
	"fmt"
	"testing"
)

func BenchmarkParseCalibrationNumber(b *testing.B) {
	source := []string{
		"t4",
		"jb4one96",
		"24xrt3twosix2",
		"44444444444444",
		"8fiveggdtrfjvrpd7six7",
		"1eightkbsixrhhphnxmjlf",
		"five9565three3nineseven",
		"onefivefivefour94eightfour",
		"6sixtj6threethree2sevenone",
		"bmfljlbbttxlvxzrfnnp319six",
		"jqllbjqndlkbxkeightdrbhjjd3",
		"jmkqqblqnxfivetwo8485eightone",
		"mlplpjkndlflk1nineninetndsqjnpmvzhkeight",
		"gqktqlbkbeightninehvfql3lbfllrnrblqchfmn6pknq",
	}

	for n, value := range source {
		b.Run(fmt.Sprintf("value_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				parseCalibrationNumber(value)
			}
		})
	}
}
