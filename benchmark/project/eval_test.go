package project

import (
	"testing"

	"github.com/larry618/eval"
	"github.com/larry618/eval_lab/benchmark"
)

func Benchmark_eval(b *testing.B) {
	params := benchmark.CreateParams()

	cc := eval.NewCompileConfig(eval.RegisterSelKeys(params))

	ctx := eval.NewCtxWithMap(cc, params)

	s := `
(and
  (or
    (= Origin "MOW")
    (= Country "RU"))
  (or
    (>= Value 100)
    (= Adults 1)))
`

	program, err := eval.Compile(cc, s)

	var out eval.Value

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, err = program.Eval(ctx)
	}
	b.StopTimer()

	if err != nil {
		b.Fatal(err)
	}
	if !out.(bool) {
		b.Fail()
	}
}
