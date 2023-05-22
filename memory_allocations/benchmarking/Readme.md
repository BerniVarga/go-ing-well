# Benchmarking

How to run:   
```
$ go test -bench . -benchtime 3s --benchmem
```

where:  
-> -bench . - runs tests with benchmarking in the current directory   
-> -benchtime 3s - run it for 3 seconds - default is 1 sec  
-> --benchmem - show memory allocations  

Results on an M2/Macbook pro/23GB Mem/8 core:

```
goos: darwin
goarch: arm64
pkg: go-memory-allocations/benchmarking
BenchmarkSprint-8               118206930               30.40 ns/op            5 B/op          1 allocs/op
BenchmarkSprintfEmpty-8         135700570               26.40 ns/op            5 B/op          1 allocs/op
BenchmarkSprintfFormatted-8     60274164                58.74 ns/op           24 B/op          2 allocs/op
```
  
-> 2nd coloumn: the number of cycles run in 3 sec  
-> 3rd coloumn: how much time was needed for one operation  
-> 4th coloumn: the number of byte/operation  
-> 5th coloumn: the number of allocations/op  


For individual test:
```
$ go test -bench BenchmarkSprintfFormatted -benchtime 3s --benchmem
```
   
Create memory profiles:

```
$ go test -bench BenchmarkSprintfFormatted -benchtime 3s --benchmem --memprofile mem.out --gcflags -m=1


# learning-material/memory_allocations/benchmarking [learning-material/memory_allocations/benchmarking.test]
./example_test.go:10:22: b does not escape
./example_test.go:14:17: ... argument does not escape
./example_test.go:14:18: "Hello" escapes to heap
./example_test.go:19:28: b does not escape
./example_test.go:28:32: b does not escape
./example_test.go:32:18: ... argument does not escape
./example_test.go:32:31: i escapes to heap
# learning-material/memory_allocations/benchmarking.test
_testmain.go:41:6: can inline init.0
_testmain.go:49:24: inlining call to testing.MainStart
_testmain.go:49:42: testdeps.TestDeps{} escapes to heap
_testmain.go:49:24: &testing.M{...} escapes to heap
goos: darwin
goarch: arm64
pkg: learning-material/memory_allocations/benchmarking
BenchmarkSprintfFormatted-8     61523005                59.36 ns/op           24 B/op          2 allocs/op
PASS
ok      learning-material/memory_allocations/benchmarking       6.620s

```


The command produced an escape analysis report before running my profiles.
It also created the following files:   
-> mem.out: is the memory profile data that i requested   
-> benchmarking.test: is the test binary rgar was built by the compiler to run the benchmarks   

The test binary is not removed after the memory profile is created, because if I want to use the profiling tools later, it can provide me with extra information.


Now I can use the `pprof` tool to analyse my sample.

```
$ go tool pprof benchmarking.test mem.out
File: benchmarking.test
Type: alloc_space
Time: May 22, 2023 at 1:28pm (BST)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) 

```   
I type in `list` or `weblist` and the function name I want to check.
List will give me a command line tool, weblist a web UI to see my results. The latter requires Graphviz to be installed on my machine.

```
$ (pprof) list fmt.Sprintf

Total: 2.48GB
ROUTINE ======================== fmt.Sprintf in /usr/local/go/src/fmt/print.go
    1.63GB     1.63GB (flat, cum) 65.74% of Total
         .          .    237:func Sprintf(format string, a ...any) string {
         .          .    238:   p := newPrinter()
         .          .    239:   p.doPrintf(format, a)
    1.63GB     1.63GB    240:   s := string(p.buf)
         .          .    241:   p.free()
         .          .    242:   return s
         .          .    243:}
         .          .    244:
         .          .    245:// Appendf formats according to a format specifier, appends the result to the byte
```

In the results we would see flat and cummulative allocations.
In my example they are the same.
`Flat allocation` means: the value on the heap is the allocation of the line represented on the file.   
`Cummmulative allocation` means: the values allocating to the heap are represented inside the call chain originating from that line in the function call   


Note:

We might want to look at how inlining is set up when we are creating a profile. 
For the web tool, it is enabled (which is the default behaviour for our compiler too). However, for the command line tool it could be different. 

Turn it on and off with the `noinlines`:  
```
$ go tool pprof --noinlines benchmarking.test mem.out
```
