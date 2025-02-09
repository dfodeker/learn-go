# Assignment
The birthdayMessage and sendingReport structs already have implementations of the getMessage method. The getMessage method returns a string, and any type that implements the method can be considered a message (meaning it implements the message interface).

Add the getMessage() method signature as a requirement on the message interface.
Complete the sendMessage function. It should return:
The content of the message.
The cost of the message, which is the length of the message multiplied by 3.
Notice that your code doesn't care at all about whether a specific message is a birthdayMessage or a sendingReport!

## Tip
The length of a string can be obtained using the len function, which returns the number of bytes.
```go
s := "Hello, World!"
fmt.Println(len(s))
// 13
```