## Professional Programming
Inspired by... \
https://github.com/charlax/professional-programming
https://github.com/mr-mig/every-programmer-should-know


## Bad Code

_Rigid_ - Is there a straight jacket of overbearing types and parameters?

_Fragile_ - Does a small change ripple through the code?

_Immobile_ - Is the code hard to refactor?

_Complex_ - Over-engineered?

_Verbose_ - When you look at it, can you tell what it's trying to do?

## Rob Pike's 5 Rules of Programming

- Rule 1. You can't tell where a program is going to spend its time. Bottlenecks occur in surprising places, so don't try to second guess and put in a speed hack until you've proven that's where the bottleneck is.
- Rule 2. Measure. Don't tune for speed until you've measured, and even then don't unless one part of the code overwhelms the rest.
- Rule 3. Fancy algorithms are slow when n is small, and n is usually small. Fancy algorithms have big constants. Until you know that n is frequently going to be big, don't get fancy. (Even if n does get big, use Rule 2 first.)
- Rule 4. Fancy algorithms are buggier than simple ones, and they're much harder to implement. Use simple algorithms as well as simple data structures.
- Rule 5. Data dominates. If you've chosen the right data structures and organized things well, the algorithms will almost always be self-evident. Data structures, not algorithms, are central to programming.

Pike's rules 1 and 2 restate Tony Hoare's famous maxim "Premature optimization is the root of all evil." Ken Thompson rephrased Pike's rules 3 and 4 as "When in doubt, use brute force.". Rules 3 and 4 are instances of the design philosophy KISS. Rule 5 was previously stated by Fred Brooks in The Mythical Man-Month. Rule 5 is often shortened to "write stupid code that uses smart objects".

### Advice for developers:
https://product.hubspot.com/blog/practical-advice-for-new-software-engineers

#### Signs of a bad programmer
http://www.yacoset.com/Home/signs-that-you-re-a-good-programmer
http://www.yacoset.com/Home/signs-that-you-re-a-bad-programmer
1. Inability to reason about code
Reasoning about code means being able to follow the execution path ("running the program in your head") while knowing what the goal of the code is.

2. Poor understanding of the language's programming model
Languar Programming Models:
	OOP, Functional, Declarative... also Procedural or Imperative

Symptoms
	- Using whatever syntax is necessary to break out of the model, then writing the remainder of the program in their familiar language's style
	- (OOP) Attempting to call non-static functions or variables in uninstantiated classes, and having difficulty understanding why it won't compile
	- (OOP) Writing lots of "xxxxxManager" classes that contain all of the methods for manipulating the fields of objects that have little or no methods of their own
	- (Relational) Treating a relational database as an object store and performing all joins and relation enforcement in client code
	- (Functional) Creating multiple versions of the same algorithm to handle different types or operators, rather than passing high-level functions to a generic implementation
	- (Functional) Manually caching the results of a deterministic function on platforms that do it automatically (such as SQL and Haskell)
	Using cut-n-paste code from someone else's program to deal with I/O and Monads
	- (Declarative) Setting individual values in imperative code rather than using data-binding

3. Deficient research skills / Chronically poor knowledge of the platform's features

Remedies
A programmer can't acquire this kind of knowledge without slowing down, and it's likely that he's been in a rush to get each function working by whatever means necessary. He needs to have the platform's technical reference handy and be able to look through it with minimal effort, which can mean either having a hard copy of it on the desk right next to the keyboard, or having a second monitor dedicated to a browser. To get into the habit initially, he should refactor his old code with the goal of reducing its instruction count by 10:1 or more.

4. Inability to comprehend pointers

Symptoms
- Failure to implement a linked list, or write code that inserts/deletes nodes from linked list or tree without losing data
- Allocating arbitrarily big arrays for variable-length collections and maintaining a separate collection-size counter, rather than using a dynamic data structure
- Inability to find or fix bugs caused by mistakenly performing arithmetic on pointers
- Modifying the dereferenced values from pointers passed as the parameters to a function, and not expecting it to change the values in the scope outside the function
- Making a copy of a pointer, changing the dereferenced value via the copy, then assuming the original pointer still points to the old value
- Serializing a pointer to the disk or network when it should have been the dereferenced value
- Sorting an array of pointers by performing the comparison on the pointers themselves

5. Difficulty seeing through recursion


6. Distrust of code

Symptoms
- Writing IsNull() and IsNotNull(), or IsTrue(bool) and IsFalse(bool) functions
- Checking to see if a boolean-typed variable is something other than true or false

### Signs that you are a mediocre programmer

1. Inability to think in sets

Funny enough, visualizing a card dealer cutting a deck of cards and interleaving the two stacks together by flipping through them with his thumbs can jolt the mind into thinking about sets and how you can operate on them in bulk. Other stimulating visualizations are:

- freeway traffic passing through an array of toll booths (parallel processing)
- springs joining to form streams joining to form creeks joining to form rivers (parallel reduce/aggregate functions)
- a newspaper printing press (coroutines, pipelines)
- the zipper tag on a jacket pulling the zipper teeth together (simple joins)
- transfer RNA picking up amino acids and joining messenger RNA within a ribosome to become a protein (multi-stage function-driven joins, see animation)
- the above happening simultaneously in billions of cells in an orange tree to convert air, water and sunlight into orange juice (Map/Reduce on large distributed clusters)

If you are writing a program that works with collections, think about all the supplemental data and records that your functions need to work on each element and use Map functions to join them together in pairs before you have your Reduce function applied to each pair.

2. Lack of critical thinking

Unless you criticize your own ideas and look for flaws in your own thinking, you will miss problems that can be fixed before you even start coding. If you also fail to criticize your own code once written, you will only learn at the vastly slower pace of trial and error. This problem originates in both lazy thinking and egocentric thinking, so its symptoms seem to come from two different directions.

3. Pinball Programming

When you tilt the board just right, pull back the pin to just the right distance, and hit the flipper buttons in the right sequence, then the program runs flawlessly with the flow of execution bouncing off conditionals and careening unchecked toward the next state transition.

Remedies
You will need to make yourself familiar with the mechanisms on your platform that help make programs robust and ductile. There are three basic kinds:

- those which stop the program before any damage is done when something unexpected happens, then helps you identify what went wrong (type systems, assertions, exceptions, etc.),
- those which direct program flow to whatever code best handles the contingency (try-catch blocks, multiple dispatch, event driven programming, etc.),
- those which pause the thread until all your ducks are in a row (WaitUntil commands, mutexes and semaphores, SyncLocks, etc.)
- There is also a fourth, Unit Testing, which you use at design time.

4. Unfamiliar with the principles of security
If the following symptoms weren't so dangerous they'd be little more than an issue of fit-n-finish for most programs, meaning they don't make you a bad programmer, just a programmer who shouldn't work on network programs or secure systems until he's done a bit of homework.

Symptoms
- Storing exploitable information (names, card numbers, passwords, etc.) in plaintext
- Storing exploitable information with ineffective encryption (symmetric ciphers with the password compiled into the program; trivial passwords; any "decoder-ring", homebrew, proprietary or unproven ciphers)
- Programs or installations that don't limit their privileges before accepting network connections or interpreting input from untrusted sources
- Not performing bounds checking or input validation, especially when using unmanaged languages
- Constructing SQL queries by string concatenation with unvalidated or unescaped input
- Invoking programs named by user input
- Code that tries to prevent an exploit from working by searching for the exploit's signature
- Credit card numbers or passwords that are stored in an unsalted hash

5. Code is a mess
	- Doesn't follow a consistent naming convention
	- Doesn't use indentation, or uses inconsistent indentation
	- Doesn't make use of whitespace elsewhere, such as between methods (or expressions, see "ANDY=NO")
	- Large chunks of code are left commented-out

### Signs that you shouldn't be a signs-that-you-re-a-bad-programmer
1. Inability to determine the order of program execution
2. Insufficient ability to think abstractly
Symptoms
	- Difficulty comprehending the difference between objects and classes
	- Difficulty implementing design patterns for your program
	- Difficulty writing functions with low cohesion
	- Incompetence with Regular Expressions
	- Lisp is opaque to you
3. Collyer Brothers syndrome
Symptoms
	- Unwilling to throw away anything, including garbage
	- Unwilling to delete anything, be it code or comments
	- The urge to build booby-traps for defense against trespassers
	- Unwilling to communicate with other people
	- Poor organization skills
4. Dysfunctional sense of causality


### How to be a Good Programmer
3. Encyclopedic grasp of the platform
Encyclopedic knowledge takes decades to acquire, but every Guru in the world got there by doing roughly the same three things each day:
	- Struggling to solve problems they find to be difficult
	- Writing about how they solved difficult problems
	- Reflecting on how they solved difficult problems
4. Thinks In Code

6. Creates their own tools
Symptoms
	- Has set up an automated build server
	- Has written their own benchmark or specialized profiler
	- Maintains an open-source project on GitHub
	- Has re-invented LISP at least once
	- Knows what a Domain Specific Language is, and has designed and written an interpreter for one
	- Extends their IDE/Editor with custom macros
	- There's a Radio Shack project enclosure on their desk with a bunch of 7-segment displays showing the number of issues assigned to them in the bug tracker