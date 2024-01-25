package main

import (
    "flag"

    "github.com/sarchlab/mgpusim/v3/benchmarks/heteromark/fir"
    "github.com/sarchlab/mgpusim/v3/samples/runner"
)

func main() {
    runner := runner.Runner{}
    runner.Init()

    benchmark := fir.NewBenchmark(runner.GPUDriver)
    benchmark.Length = 4096
    runner.Benchmark = benchmark
    runner.Run()
}
