# Go Design Patterns

Collection of design patterns in Golang and SOLID principles.

## SOLID
### Single Responsibility Principle
  - A type should only have one reason to change.
  - Separation of concerns - different types/packages handling different, independent tasks/problems.
  
### Open-Closed Principle
  - Types should be open for extension but closed for modification.

### Liskov Substitution Principle
  - You should be able to substitute an embedding type in place of its embedded part.

### Interface Segregation Principle
  - Do not put too much into an interface; split into separate interfaces.
  - YAGNI - You Ain't Going to Need It

### Dependency Inversion Principle
  - High-level modules should not depend upon low-level ones; use abstractions.

## Design Patterns
### Builder
 - A builder is a separate component used for building an object.
 - To make builder fluent, return the receiver - allows chaining.
 - Different facets of an object can be built with different builders working in tandem via a common struct.

### Factory
 - A factory function (a.k.a constructor) is a helper function for making struct instances.
 - A factory is any entity that can take care of object creation.
 - Can be a function or a dedicated struct.

### Prototype
 - To implement a prototype, partially construct an object and store it somewhere.
 - Deep copy the prototype.
 - Customize the resulting instance.
 - A prototype factory provides a convenient API for using prototypes.

### Singleton
 - Lazy one-time initialization using sync.Once.
 - Adhere to DIP: depend on interfaces, not concrete types.

### Adapter
 - Determine the API you have and the API you need.
 - Create a component which aggregates (has a pointer to, ...) the adapter.
 - Intermediate representations can pile up: use caching and other optimizations.