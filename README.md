# Go Design Patterns

Collection of design patterns in Golang and SOLID principles.

## SOLID
### Single Responsibility Principle
  - A type should only have one reason to change
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
