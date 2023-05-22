### Examples with strings

Tried to come up with a no-heap allocation example for string manipulations - and it turned out more hard than i thought so.

First of all, I was using the `println` function everywhere, since `fmt.Println` created its variables on the heap.

After that, I realised that joining the strings, unless it's done as byte appends, or joins that use the string builder below - it will do some allocations on the heap.

Happy playing.