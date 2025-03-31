# Update Users
We need a way to differentiate between standard and premium users. 
When a new user is created, they need a membership type, and that type will determine the message character limit.

## Assignment
Create a new struct called Membership, it should have:

A Type string field
A MessageCharLimit integer field
Update the User struct to embed a Membership.

Complete the newUser function. It should return a new User with all the fields set as you would expect based on the inputs.
 If the user is a "premium" member, the MessageCharLimit should be 1000, otherwise, it should only be 100.