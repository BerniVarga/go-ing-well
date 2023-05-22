### Memory allocations in Go

A program's memory consist of the stack and heap memory. 
Both are taken out of the RAM, but while the stack is cheap, heap can be expensive to handle. Let's see some whys.

The `stack memory` is a sequential memory space, used in small blocks (2k frame initial size for each go routine). It is not a shared memory, every go routine uses it's own stack, and while we create variables on the stack, we basically just move up and down the stack.
Stacks are called self-cleaning, since each time a function exists, the used stack frame is considered invalid, and every time a new function gets called, the stack space is reused. Go's way of using zero is a huge help in making this very easy to do so.
Whenever a new stack frame is taken, that gets taken out of the "heap" space. There is such a thing as maximum heap size, however, it is very large, so most machines shouldn't reach it. 


The `heap memory` is a shared memory space, and it's not a sequential space by default. Due to this, someone needs to monitor this space, and clean those variables which no longer have any reference to them. This is the job of the garbage collector.
The garbage collector occassionally needs to stop the program and analyse the memory data, and do the cleaning - which causes latency to the whole program, not only to the parts that are creating garbage. This is why we say that creating variables on the heap is expensive. 
In fact, only variables that live on the heap are considered allocations in Go.

The examples might make more sense if you look at the in this order:   
[1] variables that stay on stack   
[2] pointers that stay on stack   
[3] pointers that "escape" to the heap   
[4] strings    


Credit:  
The examples in this repository were highly influenced by all I learnt from [Ardanlabs](https://github.com/ardanlabs/gotraining/). 



