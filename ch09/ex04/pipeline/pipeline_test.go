package pipeline_test

import (
	"strconv"
	"sync"
	"testing"

	"github.com/progfay/go-training/ch09/ex04/pipeline"
)

var benchcases = []struct {
	name string
	num  int
}{
	{
		name: "1 pipeline",
		num:  1,
	},
	{
		name: "10 pipelines",
		num:  10,
	},
	{
		name: "100 pipelines",
		num:  100,
	},
	{
		name: "1000 pipelines",
		num:  1000,
	},
	{
		name: "10000 pipelines",
		num:  10000,
	},
	{
		name: "100000 pipelines",
		num:  100000,
	},
	{
		name: "1000000 pipelines",
		num:  1000000,
	},
}

func BenchmarkPipeline(b *testing.B) {
	for _, benchcase := range benchcases {
		b.Run(benchcase.name, func(b *testing.B) {
			p := pipeline.New(benchcase.num, strconv.Itoa)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				var wg sync.WaitGroup
				wg.Add(benchcase.num)

				p.Send(func(string) {
					wg.Done()
				})
				wg.Wait()
			}
		})
	}
}
