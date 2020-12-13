package dog

import (
	"fmt"
	"testing"
)

func BenchmarkYears(b *testing.B) {
	for i:=0;i<b.N;i++{
		Years(i)
	}
}

func YearTest(t *testing.T){
	fmt.Println(Years(10))
}