### Variable stays on stack

This example is supposed to show how Go uses the stack to allocate variables.
The function `calculateBonus` gets a copy of the `salary` value (called as s).
Creates that value on its stack frame, together with the multiplier and bonus variables, then returns the bonus value.

⚠️ Note:  
No cleaning is happening, since the stack is self-cleaning in Go - the next function call (`println` or `addALittleExtra`) will overwrite the previous stack addresses. Since the variables in Go always have a zero value (even if not directly initialised), we won't have a problem with reading from addresses that had a value at the same stack address before.


Run the example with printing by:  

```
$ go run main.go
salary 45000 at address 0x1400004c758
        salary 45000 with address 0x1400004c748
        multiplier 80 % with address 0x1400004c728
        bonus 36000 with address 0x1400004c730
bonus 36000 at address 0x1400004c760
        bonus 43200 with address 0x1400004c748
bonus 43200 at address 0x1400004c760

```

Ask the compiler about the memory allocations it does, with the memory optimisation flag set to level 2:    


```
go build -gcflags -m=2
# learning-material/memory_allocations/variable_stays_on_stack
./main.go:15:6: cannot inline calculateBonus: marked go:noinline
./main.go:28:6: cannot inline addALittleExtra: marked go:noinline
./main.go:3:6: cannot inline main: function too complex: cost 149 exceeds budget 80
```

What you see is some `:inlining` (inlining should help with optimisation by sometimes moving a function within another, helping the compiler to keep the variables on the stack). Ask the compiler tells: we disabled the inlining for the two functions we created.
The only function evaluated for inlining is the main, which had a cost of 149, which is over 80, so it couldn't be inlined. 

As you can see in this examples: there were no variable addresses involved, so the compiler didn't identify any variables that could potentially be created on the heap.