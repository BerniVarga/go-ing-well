### Pointers "escape" to the heap

This example is show how pointers escape to the heap. Note: while the term used is "escape", they never actually get moved, they are just allocated on the heap, instead of the stack.   

The function `calculateBonus` gets a copy of the `salary` (called as s) and returns a pointer to the bonus variable that it calculates. 
If the bonus variable has been created on the original stack as in our previous examples, once the function returns, that variables would have lived in an invalid memory space. The danger with this is that the next function call can easily overwrite the value - which is definitely not what we want.

When a variable's address is taken, the compiler runs an algorithm called escape analysis, to check whether that variables is referenced after a function return. If so, it moves it to the heap. (There are other cases, when variables get moved to the heap, but they are not covered in this example.)

Since the variable is created on the heap, it is safe for the function to return, and the main frame will get the address of bonus variable happily living in the heap now.

Run the example with printing by:  

```
$ go run main.go:
salary 45000 with address 0x1400004c758
        salary 45000 with address 0x1400004c748
        multiplier 80 % with address 0x1400004c728
        bonus 36000 with address 0x14000102000
bonus 0x14000102000 with address 0x1400004c760
```


Notice that the address of the bonus variable is not sequeantial anymore - it is taken from a different memory space.
Also notice that the newly taken stack frames always go "down" (check the memory addresses), and the valid memory addresses always grow "up". That's just a convention go uses to keep track of what's valid or invalid.



Ask the compiler about the memory allocations it does, with the memory optimisation flag set to level 1 and 2:   

```
$ go build -gcflags -m=1
# memory_allocations/escapes_to_heap
./main.go:15:2: moved to heap: bonus
```

Bonus got created on the heap.   

```
$ go build -gcflags -m=2
# learning-material/memory_allocations/escapes_to_heap
./main.go:12:6: cannot inline calculateBonus: marked go:noinline
./main.go:3:6: cannot inline main: function too complex: cost 81 exceeds budget 80
./main.go:15:2: bonus escapes to heap:
./main.go:15:2:   flow: ~r0 = &bonus:
./main.go:15:2:     from &bonus (address-of) at ./main.go:22:9
./main.go:15:2:     from return &bonus (return) at ./main.go:22:2
./main.go:15:2: moved to heap: bonus
```

Bonus got created on the heap, since it's variable was taken in line `:22`.