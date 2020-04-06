# Clean Code
Notes from the book

## Naming

#### Class Names
Classes and objects should have noun or noun phrase names like Customer, WikiPage, Account, and AddressParser. 

Avoid words like Manager, Processor, Data, or Info in the name of a class. A class name should not be a verb.

#### Method Names
Methods should have verb or verb phrase names like postPayment, deletePage, or save. Accessors, mutators, and predicates should be named for their value and prefixed with get, set, and is according to the javabean standard.

##### Pick One Word per Concept
- fetch, retrieve, get
- controller, manager, driver
- DeviceManager, Protocol Controller

##### Don't Pun
Don't use the same word for two purposes.

##### Use Solution Domain Names
`AccountVisitor` means a great deal to a programmer familiar with the VISITOR pattern.

##### Add Meaningful Context
Imagine that you have variables named firstName, lastName, street, houseNumber, city, state, and zipcode. Taken together it’s pretty clear that they form an address. But what if you just saw the state variable being used alone in a method? Would you automatically infer that it was part of an address?

You can add context by using prefixes: addrFirstName, addrLastName, addrState, and so on. At least readers will understand that these variables are part of a larger structure. Of course, a better solution is to create a class named Address. Then, even the compiler knows that the variables belong to a bigger concept


## Functions
##### Small!
##### Blocks and Indenting
##### Do One Thing
##### One Level of Abstraction per Function
##### The Stepdown Rule
##### Switch Statements
##### Descriptive Names
##### Have No Side Effects
##### Command Query Separation

## Comments
## Formatting
## Objects and Data Structures
#### Data/Object Anti-Symmetry
Again, we see the complimentary nature of these two definitions; they are virtual opposites! This exposes the fundamental dichotomy between objects and data structures:

Procedural code (code using data structures) makes it easy to add new functions without changing the existing data structures. OO code, on the other hand, makes it easy to add new classes without changing existing functions.

The complement is also true:

Procedural code makes it hard to add new data structures because all the functions must change. OO code makes it hard to add new functions because all the classes must change.

#### The Law of Demeter
Module should not know about the innards of he objects it manipulates

## Error Handling

##### Write Your Try-Catch-Finally Statement First

## Boundaries

## Unit Tests

## Classes

## Systems

## Emergence

## Concurrency

## Successive Refinement