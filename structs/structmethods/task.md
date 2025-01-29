# Assignment: Struct Methods
Let's clean up Textio's authentication logic. We store our user's authentication data inside an authenticationInfo struct.
 We need a method that can take that data and return a basic authorization string.

The format of the string should be:

```
Authorization: Basic USERNAME:PASSWORD

```

Create a method on the authenticationInfo struct called getBasicAuth that returns the formatted string.