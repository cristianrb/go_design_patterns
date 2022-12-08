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
### Adapter
- Determine the API you have and the API you need.
- Create a component which aggregates (has a pointer to, ...) the adapter.
- Intermediate representations can pile up: use caching and other optimizations.

### Bridge
- Decouple abstraction from implementation.
- Both can exist as hierarchies.
- A stronger form of encapsulation.

### Builder
 - A builder is a separate component used for building an object.
 - To make builder fluent, return the receiver - allows chaining.
 - Different facets of an object can be built with different builders working in tandem via a common struct.

### Chain of Responsibility
 - Chain of Responsibility can be implemented as a linked list of pointers or a centralized construct.
 - Enlist objects in the chain, possibly controlling their order.
 - Control object removal from chain.

### Command
 - Encapsulate all details of an operation in a separate object.
 - Define functions for applying the command (either in hte command itself, or elsewhere)
 - Optionally define instructions for undoing the command.
 - Can create composite commands (a.k.a. macros)

### Composite
 - Objects can use other objects via composition.
 - Some composed and singular objects need similar/identical behaviors.
 - Composite design pattern lets us treat both types of objects uniformly.
 - Iteration supported with the Iterator design pattern.

### Decorator
 - A decorator embeds the decorator object.
 - Adds utility fields and methods to augment the objects features.
 - Often used to emulate multiple inheritance (may require extra work)

### Facade
 - Build a Facade to provide a simplified API over a set of components.
 - May wish to (optionally) expose internals through the facade.
 - May allow users to escalate to use more complex APIs if they need to.

### Factory
 - A factory function (a.k.a constructor) is a helper function for making struct instances.
 - A factory is any entity that can take care of object creation.
 - Can be a function or a dedicated struct.

### Flyweight
 - Store common data externally.
 - Specify an index or a pointer into the external data store.
 - Define the idea of ranges on homogeneous collections and store data related to those ranges.

### Interpreter
 - Barring simple cases,an interpreter act in two stages.
 - Lexing turns text into a set of tokens.
 - Parsing tokens into meaningful constructs (AST = Abstract Syntax Tree)
 - Parsed data can then be traversed using the Visitor pattern.

### Iterator
 - An iterator specifies how you can traverse an object.
 - Moves along the iterated collection, indicating when last element has been reached.
 - Not idiomatic in Go (no standard Iterable interface)

### Mediator
 - Create the mediator and have each object in the system point to it.
 - Mediator engages in bidirectional communication with its connected components.
 - Mediator has methods the components can call.
 - Components have methods the mediator can call.
 - Even processing (e.g., Rx) libraries make communication easier to implement.

### Memento
 - Mementos are used to roll back states arbitrarily.
 - A memento is simply a token/handle with (typically) no methods of its own.
 - A memento is not required to expose directly the state(s) to which it reverts the system.
 - Can be used to implement undo/redo.

### Observer
 - Observer is an intrusive approach.
 - Must provide a way of clients to subscribe.
 - Event data sent from observable to all observers.
 - Data represented as interface{} (or any)
 - Unsubscription is possible.
 - More info with channels and goroutines: https://eli.thegreenplace.net/2020/pubsub-using-channels-in-go/

### Prototype
 - To implement a prototype, partially construct an object and store it somewhere.
 - Deep copy the prototype.
 - Customize the resulting instance.
 - A prototype factory provides a convenient API for using prototypes.

### Proxy
 - A proxy has the same interface as the underlying object.
 - To create a proxy, simply replicate the existing interface of an object.
 - Add relevant functionality to the redefined methods.
 - Different proxies (communication, logging, caching, etc.) have completely different behaviors.

### Singleton
 - Lazy one-time initialization using sync.Once.
 - Adhere to DIP: depend on interfaces, not concrete types.

### State
 - Given sufficient complexity, it pays to formally define possible states and events/triggers.
 - Can define:
   - State entry/exit behaviours.
   - Action when a particular event causes a transition.
   - Guard conditions enabling/disabling a transition.
   - Default action when no transitions are found for an event.

### Strategy
 - Define an algorithmat a high level.
 - Define the interface you expect each strategy to follow.
 - Support the injection of the strategy into the high-level algorithm.

### Template Method
 - Very similar to Strategy
 - Typical implementation:
   - Define an interface with common operations.
   - Make use of those operations inside a function.
 - Alternative functional approach:
   - Make a function that takes several functions.
   - Can pass in functions that capture local state.
   - No need for either structs or interfaces.

### Visitor
 - Propagate an Accept(v *Visitor) method throughout the entire hierarchy.
 - Create a visitor with VisitFoo(f Foo), VisitBar(b Bar), ... for each element in the hierarchy.
 - Each Accept() simply calls Visitor.VisitXxx(self)