# OOP & Software Engineering Design Principles — Complete Syllabus

> A comprehensive guide to every OOP concept and design principle used in real-world software engineering.
> Each concept is mapped across **C++**, **Python**, and **Go** so you can see how different languages handle the same idea.

---

## Part 1: The 4 Pillars of OOP

These are the foundational concepts that every OOP language is built on.

### 1.1 Encapsulation (Data Hiding)
- [ ] What is encapsulation?
- [ ] Bundling data + methods that operate on that data together
- [ ] Access modifiers / visibility controls
  - **C++:** `public`, `private`, `protected` keywords
  - **Python:** `_single_underscore` (convention), `__double_underscore` (name mangling)
  - **Go:** Capitalized = exported (public), lowercase = unexported (private to package)
- [ ] Getters and Setters
- [ ] Why hiding internal state matters (preventing invalid states)

### 1.2 Abstraction (Hiding Complexity)
- [ ] What is abstraction?
- [ ] Showing only "what it does", hiding "how it works"
- [ ] Abstract classes
  - **C++:** Classes with `virtual` functions (pure virtual = `virtual void foo() = 0`)
  - **Python:** `ABC` (Abstract Base Class) from `abc` module
  - **Go:** No abstract classes. Uses **Interfaces** instead
- [ ] Interfaces
  - **C++:** Pure virtual classes (all methods are `= 0`)
  - **Python:** `ABC` with `@abstractmethod`
  - **Go:** `type MyInterface interface { Method() }` — implicit implementation (duck typing)

### 1.3 Inheritance ("Is-a" Relationship)
- [ ] What is inheritance?
- [ ] Parent/Base class → Child/Derived class
- [ ] Types of inheritance:
  - **Single inheritance** (one parent) — All 3 languages
  - **Multiple inheritance** (multiple parents) — C++ ✅, Python ✅, Go ❌
  - **Multilevel inheritance** (chain: A → B → C)
  - **Hierarchical inheritance** (one parent, many children)
  - **Hybrid inheritance**
- [ ] Method Overriding (child redefines parent's method)
- [ ] `super()` / base class constructor calls
  - **C++:** `BaseClass::method()` or initializer list
  - **Python:** `super().__init__()`
  - **Go:** ❌ No inheritance. Uses **Composition** instead (see 1.5)
- [ ] Diamond Problem (ambiguity in multiple inheritance)
  - **C++:** Solved with `virtual` inheritance
  - **Python:** Solved with MRO (Method Resolution Order) / C3 Linearization
  - **Go:** Not applicable (no inheritance)

### 1.4 Polymorphism ("Many Forms")
- [ ] What is polymorphism?
- [ ] **Compile-time polymorphism** (Static / Early Binding):
  - Function Overloading (same name, different parameters)
    - **C++:** ✅ Supported natively
    - **Python:** ❌ Not supported (uses default args / `*args` instead)
    - **Go:** ❌ Not supported
  - Operator Overloading
    - **C++:** ✅ `operator+()`, `operator<<()` etc.
    - **Python:** ✅ `__add__`, `__str__`, `__len__` (dunder methods)
    - **Go:** ❌ Not supported
  - Templates / Generics
    - **C++:** `template<typename T>`
    - **Python:** `typing.Generic[T]`
    - **Go:** `func Foo[T any](val T)` (Generics added in Go 1.18)
- [ ] **Runtime polymorphism** (Dynamic / Late Binding):
  - Method Overriding (child provides its own implementation)
    - **C++:** `virtual` keyword + vtable
    - **Python:** Just redefine the method in child class
    - **Go:** Achieved through **Interfaces** (implicit satisfaction)
  - Interface-based polymorphism
    - **Go example:** `io.Reader` — any struct with a `Read()` method satisfies it

### 1.5 Composition ("Has-a" Relationship)
- [ ] Composition vs Inheritance
  - Inheritance = "Dog **is an** Animal"
  - Composition = "Car **has an** Engine"
- [ ] Why Go chose Composition over Inheritance
- [ ] Struct Embedding in Go
  - **Go:** `type EBook struct { Book; FileSize int }` — EBook "has a" Book
  - **C++:** Member objects / `has-a` via class members
  - **Python:** Instance variables holding other objects
- [ ] **Favor Composition over Inheritance** — one of the most important design rules

### 1.6 Association, Aggregation, Composition (The 3 Relationships)
- [ ] **Association:** Two objects are related but independent (Teacher ↔ Student)
- [ ] **Aggregation:** "Has-a" but the child can exist independently (Department has Employees, employees exist without the department)
- [ ] **Composition:** "Has-a" but the child CANNOT exist independently (House has Rooms, rooms don't exist without the house)

---

## Part 2: Design Principles

These are the "rules of thumb" that experienced engineers follow to write clean, maintainable code.

### 2.1 SOLID Principles
- [ ] **S — Single Responsibility Principle (SRP)**
  - A class/function should have only ONE reason to change
  - Example: Don't mix database logic and email-sending in the same function
- [ ] **O — Open/Closed Principle (OCP)**
  - Open for extension, closed for modification
  - Add new behavior by adding new code, not changing existing code
  - Achieved via interfaces and polymorphism
- [ ] **L — Liskov Substitution Principle (LSP)**
  - If function accepts a parent type, any child type must work without breaking
  - In Go: if a function accepts an `io.Reader`, any struct implementing `Read()` must work correctly
- [ ] **I — Interface Segregation Principle (ISP)**
  - Many small, focused interfaces > one giant interface
  - Go is the king of this: `io.Reader` (1 method), `io.Writer` (1 method), `io.ReadWriter` (combines both)
- [ ] **D — Dependency Inversion Principle (DIP)**
  - High-level modules should NOT depend on low-level modules
  - Both should depend on abstractions (interfaces)
  - Example: Your API handler depends on a `Storage` interface, not directly on PostgreSQL

### 2.2 DRY (Don't Repeat Yourself)
- [ ] Every piece of knowledge should have a single, authoritative representation
- [ ] If you copy-paste code, extract it into a function
- [ ] Opposite: **WET** (Write Everything Twice / We Enjoy Typing)

### 2.3 KISS (Keep It Simple, Stupid)
- [ ] Simplicity should be a key goal; avoid unnecessary complexity
- [ ] Go's entire philosophy is built on KISS
- [ ] Don't over-engineer solutions

### 2.4 YAGNI (You Aren't Gonna Need It)
- [ ] Don't build features until you actually need them
- [ ] Avoid premature optimization and over-abstraction
- [ ] Write code for today's requirements, not imaginary future ones

### 2.5 Law of Demeter (Principle of Least Knowledge)
- [ ] An object should only talk to its immediate friends
- [ ] Bad: `order.getCustomer().getAddress().getCity()` (chaining too deep)
- [ ] Good: `order.getShippingCity()` (hide the chain behind a method)

### 2.6 Separation of Concerns (SoC)
- [ ] Different responsibilities should be in different modules/files
- [ ] Example: Controllers handle HTTP, Services handle business logic, Repositories handle database
- [ ] Related to SRP but applied at the architecture level

### 2.7 Principle of Least Astonishment (POLA)
- [ ] Code should behave the way other developers expect it to
- [ ] No surprises — a `Delete()` function should delete, not archive

### 2.8 Fail Fast
- [ ] Detect errors as early as possible and report them immediately
- [ ] Don't let bad data silently flow through your system
- [ ] Go's explicit error handling (`if err != nil`) is a great example of this

---

## Part 3: Design Patterns (The Classics)

Design patterns are reusable solutions to common software design problems. Grouped into 3 categories:

### 3.1 Creational Patterns (How objects are created)
- [ ] **Singleton** — Ensure a class has only ONE instance (e.g., database connection pool)
- [ ] **Factory** — Create objects without specifying the exact class
- [ ] **Builder** — Construct complex objects step by step
- [ ] **Prototype** — Clone existing objects instead of creating new ones

### 3.2 Structural Patterns (How objects are composed)
- [ ] **Adapter** — Make incompatible interfaces work together
- [ ] **Decorator** — Add behavior to objects dynamically (Go: middleware wrapping handlers!)
- [ ] **Facade** — Provide a simple interface to a complex subsystem
- [ ] **Proxy** — Control access to another object

### 3.3 Behavioral Patterns (How objects communicate)
- [ ] **Observer** — When one object changes, all dependents are notified (pub/sub)
- [ ] **Strategy** — Define a family of algorithms, make them interchangeable
- [ ] **Iterator** — Access elements of a collection sequentially
- [ ] **Middleware/Chain of Responsibility** — Pass request through a chain of handlers (very common in Go backends!)

---

## Part 4: Go-Specific OOP Concepts

Go has its own unique approach to OOP. These are Go-specific things you should know:

### 4.1 Structs (Go's replacement for Classes)
- [ ] Defining structs
- [ ] Struct methods (value receivers vs pointer receivers)
  - `func (b Book) GetTitle() string` — value receiver (read-only copy)
  - `func (b *Book) SetTitle(t string)` — pointer receiver (modifies original)
- [ ] Struct tags (`json:"id"`, `db:"user_id"`)
- [ ] Anonymous structs

### 4.2 Interfaces in Go
- [ ] Implicit interface satisfaction (no `implements` keyword!)
- [ ] Empty interface `interface{}` / `any` (like `Object` in Java or `void*` in C++)
- [ ] Type assertion (`val, ok := i.(string)`)
- [ ] Type switch
- [ ] Stringer interface (`String() string` — like `toString()` in Java or `__str__` in Python)
- [ ] Error interface (`Error() string`)

### 4.3 Embedding (Go's Composition)
- [ ] Struct embedding (promoted fields and methods)
- [ ] Interface embedding (combining small interfaces into bigger ones)
  - Example: `io.ReadWriter` embeds `io.Reader` + `io.Writer`

### 4.4 Constructor Pattern
- [ ] Go has no constructors — use `NewXxx()` factory functions instead
  - Example: `func NewBook(title string, price float64) *Book`

---

## Part 5: Other Important Engineering Concepts

### 5.1 Coupling & Cohesion
- [ ] **Tight Coupling** — Modules are heavily dependent on each other (BAD)
- [ ] **Loose Coupling** — Modules are independent, interact via interfaces (GOOD)
- [ ] **High Cohesion** — Related things are grouped together (GOOD)
- [ ] **Low Cohesion** — Unrelated things are mixed together (BAD)

### 5.2 Code Smells
- [ ] God Object / God Function (one thing does everything)
- [ ] Long Parameter Lists
- [ ] Deep Nesting (too many if/else levels)
- [ ] Magic Numbers (use constants instead of `if status == 200`)
- [ ] Dead Code (unused variables and functions)

### 5.3 Refactoring Techniques
- [ ] Extract Function
- [ ] Rename Variable / Function for clarity
- [ ] Replace Magic Number with Named Constant
- [ ] Introduce Parameter Object (group related params into a struct)

---

## Study Order (Recommended)

| Phase | Topics | Priority |
|-------|--------|----------|
| 1 | 4 Pillars (Encapsulation, Abstraction, Inheritance, Polymorphism) | 🔴 Must Know |
| 2 | Composition vs Inheritance | 🔴 Must Know |
| 3 | SOLID Principles | 🔴 Must Know |
| 4 | DRY, KISS, YAGNI | 🟡 Important |
| 5 | Go-Specific OOP (Structs, Interfaces, Embedding) | 🔴 Must Know |
| 6 | Design Patterns (Singleton, Factory, Observer, Middleware) | 🟡 Important |
| 7 | Coupling, Cohesion, Code Smells | 🟢 Good to Know |
| 8 | Association, Aggregation, Composition Relationships | 🟢 Good to Know |

---

> **Note:** You don't need to memorize all of this at once. As you build more projects in Go throughout this tutorial, you will naturally encounter and apply these concepts. Check off items as you learn them!
