# Interface Implementation
Interfaces are implemented implicitly.

A type never declares that it implements a given interface. If an interface exists and a type has the proper methods defined, then the type automatically fulfills that interface.

>A quick way of checking whether a struct implements an >interface is to declare that a function takes an interface as >an argument. If the function can take the struct as an ?>argument, then the struct implements the interface.

## Assignment
At Textio we have full-time employees and contract employees. We have been tasked with making a more general employee interface so that dealing with different employee types is simpler.

Run the code. You should see an error indicating the contractor type does not fulfill the employee interface.

Add the missing getSalary method to the contractor type so that it fulfills the employee interface.

A contractor's salary is their hourly pay multiplied by how many hours they work per year.