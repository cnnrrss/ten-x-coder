# SOLID & OOP

SOLID in Go: https://dave.cheney.net/2016/08/20/solid-go-design

## SOLID principles for OOP:

### Single Responsibility Principle

Each method does one thing, well.

No long methods

Only track 7 things at a time

Doug McIlroy’s Unix philosophy; small, sharp tools which combine to solve larger tasks

### Open/Closed Principle
Open for extension, closed for modification. _–Bertrand Meyer_

Classes should only have access on an as needed basis.

In Go we can imbed

```golang
type Cat struct {}
func (c Cat) Legs() int { return 4 }
func (c Cat) PrintLegs() { fmt.Printf("I have %d legs\n", c.Legs())}
type OctoCat struct {}
func (o OctoCat) Legs() int { return 5 }
func main() {
	var oc OctoCat
	fmt.Println(octo.Legs()) // 5
	octo.PrintLegs() // I have 4 legs
}
```


### Liskov Substitution Principle
Code should be easily extensible (and deletable).
Substitution is the purview of Go interfaces

### Interface Segregation Principle

A great rule of thumb for Go is accept interfaces, return structs. _–Jack Lindamood_


### Dependency Inversion Principle

High-level modules should not depend on low-level modules. \
Both should depend on abstractions. \
Abstractions should not depend on details. Details should depend on abstractions. \
_–Robert C. Martin_ aka Uncle Bob

Inversion of control

Inject dependencies to other classes

#### 4 Pillars of OOP

• Abstraction - "hide complexity". It enables you to focus on _what_ the object does not _how_ it does it.

• Encapsulation - Prevents access to implementation details i.e. Wrapping up a data member and a method together into a single unit. 
	• Binding together of data and methods and hence gaining control over whats exposed, how its exposed and data modification.
	• How you define the boundary of your class, data and methods. The four visibility modifiers are public, default(without any keywords), protected and private.

• Polymorphism - "one name/form many uses"
	• Compile-time: compiler identifies which polymorphism form it has to take and execute at compile time. (Overloading).
	• Run-time: compiler identifies which polymorphism form it has to take and execute at runtime but not at compile time. (Overriding).

• Inheritance - "is a" (Example: parent->child relationship). 
	• When a class includes a property of another class it is Inheritance.
