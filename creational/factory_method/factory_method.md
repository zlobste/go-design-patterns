### Factory method :factory:

#### Description
The Factory Method pattern is a design pattern used to define a runtime interface for creating an object. It’s called a factory because it creates various types of objects without necessarily knowing what kind of object it creates or how to create it.

#### Purpose
Allows the sub-classes to choose the type of objects to create at runtime
It provides a simple way of extending the family of objects with minor changes in application code.
Promotes the loose-coupling by eliminating the need to bind application-specific structs into the code.

#### Important aspects when we implement the Factory Method design pattern are:
* Designing the arguments of the factory method.
* Considering an internal object pool that will allow object cache and reuse instead of created from scratch.