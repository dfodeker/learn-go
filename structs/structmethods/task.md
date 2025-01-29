# Assignment: Struct Methods
Let's clean up Textio's authentication logic. We store our user's authentication data inside an authenticationInfo struct.
 We need a method that can take that data and return a basic authorization string.

The format of the string should be:

```
Authorization: Basic USERNAME:PASSWORD

```

Create a method on the authenticationInfo struct called getBasicAuth that returns the formatted string.

```go
//example usage
type rect struct {
  width int
  height int
}

// area has a receiver of (r rect)
// rect is the struct
// r is the placeholder
func (r rect) area() int {
  return r.width * r.height
}

var r = rect{
  width: 5,
  height: 10,
}

fmt.Println(r.area())
// prints 50
```