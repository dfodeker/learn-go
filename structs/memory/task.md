## Assignment: Memory
Our over-engineering boss is at it again.
He's heard about memory layout and wants to squeeze every last byte out of our structs.

Run the tests to see the current size of the structs, then update the struct definitions to minimize memory usage.


### Should I Panic?
To be honest, you should not stress about memory layout. However, if you have a specific reason to be concerned about memory usage, aligning the fields by size (largest to smallest) can help. You can also use the reflect package to debug the memory layout of a struct:
```go
typ := reflect.TypeOf(stats{})
fmt.Printf("Struct is %d bytes\n", typ.Size())
```