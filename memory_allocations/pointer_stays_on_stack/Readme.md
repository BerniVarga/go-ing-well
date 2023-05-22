### Pointer stays on stack

This example is show how pointers can still remain on the stack - despite them being a candidate for living on the heap.
The function `calculateBonus` gets a copy of the `salary` address (called as s). Since s is a pointer, it can be a candidate for being allocated on the heap.
However, since s still points to a valid memory space, multiplier gets allocated on the stack, but the operation modifying the value that s points too (salary in the original main frame) succeeds, and all the values remain on the stack.


⚠️ Note:  
Sharing down usually stays on the stack.


Run the example with printing by:  

```
$ go run main.go        
salary 45000 with address 0x1400004c760
        salary 0x1400004c760 with address 0x1400004c748
        multiplier 80 % with address 0x1400004c730
        bonus 0x1400004c760 with address 0x1400004c748
bonus 36000 with address 0x1400004c760
```

Check the values to have some clarity on what's happening on the stack.


Ask the compiler about the memory allocations it does, with the memory optimisation flag set to level 1 and 2:    

```
$ go build -gcflags -m=1
# learning-material/memory_allocations/pointer_stays_on_stack
./main.go:3:6: can inline main
./main.go:11:21: s does not escape
```

`s` is a candidate, but does not escape to the heap   

```
$ go build -gcflags -m=2
# learning-material/memory_allocations/pointer_stays_on_stack
./main.go:11:6: cannot inline calculateBonus: marked go:noinline
./main.go:3:6: can inline main with cost 78 as: func() { s := 45000; println("salary", s, "with address", &s); calculateBonus(&s); println("bonus", s, "with address", &s) }
./main.go:11:21: s does not escape
```

`s` is a candidate, but it does not escape to the heap   