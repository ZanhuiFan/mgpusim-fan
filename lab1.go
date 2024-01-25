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

type Runner struct {
    engine            sim.Engine
    GPUDriver         *driver.Driver
    Benchmark         benchmarks.Benchmark
}

func (r *Runner) Init() {
    r.engine, r.GPUDriver = platform.BuildNR9NanoPlatform(1)
}

func (r *Runner) Run() {
    r.Benchmark.Run()
    r.engine.Finished()
}
