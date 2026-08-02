# OOP & Software Engineering Design — Complete Syllabus
### C++ · Python · TypeScript · Go

> **How this is organized (read this first):**
>
> C++, Python and TypeScript are **classical OOP** languages — classes, inheritance, `extends`, constructors, `this`.
> **Go deliberately rejected that model.** No classes, no inheritance, no constructors, no `extends`, no exceptions.
> Mixing them line-by-line makes both confusing, so this syllabus splits into tracks:
>
> | Part | What it covers | Applies to |
> |---|---|---|
> | **A** | Universal concepts — true everywhere, language-agnostic | All 4 |
> | **B** | Classical OOP mechanics | C++ · Python · TypeScript |
> | **C** | Go's model — structs, embedding, implicit interfaces | Go only |
> | **D** | Translation table — same idea, four spellings | All 4 |
> | **E** | Design principles (SOLID, DRY, GRASP…) — language-independent | All 4 |
> | **F** | Design patterns — with per-language notes | All 4 |
> | **G** | Concurrency patterns | All 4 |
> | **H** | Architecture | All 4 |
> | **I** | Errors, testing, code quality | All 4 |
> | **J** | Per-language deep dives | Each separately |
> | **K** | Study plan | — |
>
> Legend: 🔴 Must know · 🟡 Important · 🟢 Good to know · ⚪ Advanced

---
---

# PART A — Universal Concepts
*True in every language. Learn these as ideas first, syntax second.*

## A.1 The Object Model 🔴
- [ ] **Type** (blueprint) vs **Instance** (a concrete value of that type)
- [ ] **State** (data) + **Behaviour** (operations) bundled together
- [ ] **Identity vs equality** — "same object" vs "same value"
- [ ] **The receiver** — how a method knows which instance it acts on
- [ ] **Invariants** — states an object must *never* be in. This is the real reason OOP exists

## A.2 The Four Pillars (as ideas) 🔴
- [ ] **Encapsulation** — hide internal state, expose a controlled surface
- [ ] **Abstraction** — expose *what*, hide *how*
- [ ] **Inheritance** — "is-a", reuse by extension
- [ ] **Polymorphism** — one interface, many implementations
- [ ] **+ Composition** — "has-a", reuse by assembly (the pillar everyone forgets, and the most useful one)

## A.3 Polymorphism — the two kinds 🔴
- [ ] **Static / compile-time** — overloading, operator overloading, generics/templates
- [ ] **Dynamic / runtime** — the actual implementation is chosen while the program runs
- [ ] **Subtype polymorphism** (inheritance-based) vs **structural/duck typing** (shape-based)
- [ ] **Ad-hoc** (overloading) vs **parametric** (generics) vs **subtype** — the formal taxonomy

## A.4 Object Relationships 🟢
- [ ] **Dependency** — "uses-a" temporarily (a method parameter)
- [ ] **Association** — related but independent (Teacher ↔ Student)
- [ ] **Aggregation** — has-a, child outlives the parent (Department ◇— Employee)
- [ ] **Composition** — has-a, child dies with the parent (House ◆— Room)
- [ ] Multiplicity (1:1, 1:N, N:M), UML class-diagram notation
- [ ] Why bidirectional links create coupling

## A.5 Object Lifecycle & Memory 🟡
- [ ] Creation → initialization → use → destruction
- [ ] **Value semantics vs reference semantics** — does assignment copy or alias?
- [ ] **Shallow vs deep copy** — and the bugs that come from getting it wrong
- [ ] Stack vs heap allocation
- [ ] Ownership: who is responsible for cleaning this up?
- [ ] Manual memory (C++) vs garbage collection (Python/TS/Go)
- [ ] Deterministic cleanup (RAII, `defer`, `using`, `with`) vs non-deterministic (finalizers)

## A.6 Type Systems 🟡
- [ ] Static vs dynamic typing; strong vs weak
- [ ] **Nominal typing** (C++, Java: "declares it implements X") vs **structural typing** (Go, TypeScript: "has the right shape")
- [ ] Compile-time vs runtime type checks
- [ ] Type inference; gradual typing (Python hints, TS)
- [ ] Covariance / contravariance ⚪
- [ ] Null safety — `nullptr`, `None`, `null | undefined`, `nil`

---
---

# PART B — Classical OOP
### C++ · Python · TypeScript
*Skip to Part C if you only care about Go — but read this anyway, since 90% of OOP literature assumes this model.*

## B.1 Classes & Objects 🔴
- [ ] Declaring a class, fields, methods
  - **C++:** `class Book { ... };` (`struct` = class with default-public)
  - **Python:** `class Book:`
  - **TypeScript:** `class Book { ... }`
- [ ] The receiver
  - **C++:** implicit `this` pointer
  - **Python:** explicit `self` as first parameter
  - **TypeScript:** `this` — ⚠️ **rebinding hazard**: `this` depends on *call site*; fix with arrow-function class fields or `.bind(this)`
- [ ] Fields vs methods; instance vs class scope

## B.2 Constructors 🔴
- [ ] **C++:** `Book(std::string t) : title_(t) {}` — default / parameterized / delegating / copy / move ctors; `explicit` to block implicit conversion; member-initializer list runs before the body
- [ ] **Python:** `__init__` (initializer) vs `__new__` (allocator); alternate constructors via `@classmethod`
- [ ] **TypeScript:** `constructor(...)`; **parameter properties** — `constructor(private title: string)` auto-assigns; static factory methods as the idiomatic alternative to overloaded ctors
- [ ] Initialization order (base → members → body)
- [ ] Constructors should produce a *valid* object or fail — never a half-built one

## B.3 Destructors & Cleanup 🟡
- [ ] **C++:** `~Book()`; **RAII** — the single most important C++ idea; ⚠️ **virtual destructor required** when deleting through a base pointer
- [ ] **Python:** `__del__` is unreliable — use **context managers** (`__enter__`/`__exit__` + `with`), or `contextlib.contextmanager`
- [ ] **TypeScript:** GC only; explicit `dispose()` / `close()`; `using` declarations (TS 5.2 explicit resource management); `FinalizationRegistry` ⚪
- [ ] Deterministic vs non-deterministic destruction

## B.4 Encapsulation & Access Control 🔴
- [ ] **C++:** `public` / `private` / `protected`; `friend` classes & functions; `const` methods; `mutable`
- [ ] **Python:** `_single` (convention: internal), `__double` (name mangling) — **no real privacy**, "we're all consenting adults"
- [ ] **TypeScript:** `public` / `private` / `protected` / `readonly` — ⚠️ **compile-time only, erased at runtime**; `#privateField` is *real* runtime privacy (ES2022)
- [ ] **Properties / computed accessors**
  - **C++:** just methods
  - **Python:** `@property`, `@x.setter`, descriptors
  - **TypeScript:** `get x()` / `set x(v)`
- [ ] Getters/setters — and when *not* to write them (see Tell Don't Ask, E.12)

## B.5 Inheritance 🔴
- [ ] Syntax
  - **C++:** `class EBook : public Book { };` (also `protected`/`private` inheritance ⚪)
  - **Python:** `class EBook(Book):`
  - **TypeScript:** `class EBook extends Book { }` — **single inheritance only**
- [ ] **Multiple inheritance:** C++ ✅ · Python ✅ · TypeScript ❌ (mixins via functions/`implements` instead)
- [ ] Types: single, multilevel, hierarchical, hybrid
- [ ] **Method overriding** vs **method hiding/shadowing**
- [ ] Calling the parent
  - **C++:** `Base::method()`, member-initializer list
  - **Python:** `super().__init__()` — cooperative multiple inheritance
  - **TypeScript:** `super.method()`; ⚠️ `super()` **must** be called before touching `this` in a derived constructor
- [ ] **Diamond problem**
  - **C++:** `virtual` base classes
  - **Python:** **MRO** / C3 linearization (`Cls.__mro__`)
  - **TypeScript:** N/A (single inheritance)
- [ ] **Mixins:** Python (multiple inheritance) · TypeScript (mixin factory functions returning classes) · C++ (CRTP) ⚪
- [ ] Preventing extension: `final` (C++/TS-via-lint), `@final` (Python typing)
- [ ] **Why inheritance is overused** — fragile base class, yo-yo problem, deep hierarchies

## B.6 Abstraction: Abstract Classes & Interfaces 🔴
- [ ] **Abstract class**
  - **C++:** pure virtual — `virtual void Draw() = 0;`
  - **Python:** `from abc import ABC, abstractmethod`
  - **TypeScript:** `abstract class Shape { abstract draw(): void }`
- [ ] **Interface**
  - **C++:** class with all methods pure virtual; C++20 **concepts** for compile-time constraints
  - **Python:** `ABC` (nominal) or `typing.Protocol` (**structural** — closest to Go/TS)
  - **TypeScript:** `interface Shape { draw(): void }` + `implements` — **structurally typed**, `implements` is only a check, not a requirement
- [ ] Abstract class vs interface — when to use which (state+behaviour vs pure contract)
- [ ] TS-only: `type` vs `interface`, declaration merging, index signatures

## B.7 Polymorphism in Practice 🔴
**Compile-time**
- [ ] **Overloading:** C++ ✅ real · Python ❌ (defaults, `*args`, `functools.singledispatch`) · TypeScript ⚠️ *signature* overloads only — one implementation
- [ ] **Operator overloading:** C++ ✅ (`operator+`, `operator<<`, `<=>`) · Python ✅ dunders (`__add__`, `__eq__`, `__len__`) · TypeScript ❌
- [ ] **Generics:** C++ `template<typename T>` (+ specialization, SFINAE, concepts) · Python `Generic[T]`, `TypeVar` · TypeScript `<T>`, `extends` constraints, conditional/mapped types, `keyof`

**Runtime**
- [ ] **C++:** `virtual` + **vtable/vptr**; ⚠️ **object slicing** when assigning derived to base by value; `override` and `final` specifiers — always write `override`
- [ ] **Python:** every method is virtual; dispatch = dict lookup along the MRO
- [ ] **TypeScript:** prototype-chain dispatch; all methods virtual
- [ ] Runtime type checks: `dynamic_cast`/`typeid` (C++) · `isinstance` (Python) · `instanceof` + **type guards / discriminated unions** (TS)

## B.8 Copy Semantics 🟡
- [ ] **C++:** copy ctor, copy assignment, **move ctor/assignment**, `std::move`, copy elision, **Rule of 0/3/5**
- [ ] **Python:** `copy.copy()` vs `copy.deepcopy()`, `__copy__`/`__deepcopy__`; ⚠️ **mutable default argument** trap
- [ ] **TypeScript:** spread `{...obj}` is shallow; `structuredClone()` for deep; `Readonly<T>`, `as const` for immutability
- [ ] Immutability & value objects across all three

## B.9 Static / Class-level Members 🟡
- [ ] **C++:** `static` data members and methods; `constexpr`, `enum class`
- [ ] **Python:** class attributes, `@staticmethod` vs `@classmethod` (know the difference cold), `enum.Enum`
- [ ] **TypeScript:** `static` members, `static` blocks, `enum` (and why `as const` unions are often better)
- [ ] Class-level state = shared mutable state = be careful

---
---

# PART C — Go's Model
### Everything above, rebuilt from different primitives

> Go has **no classes, no inheritance, no constructors, no `extends`, no exceptions, no overloading, no operator overloading**.
> It still does OOP — encapsulation, abstraction, polymorphism, composition — just with structs, embedding and implicit interfaces.

## C.1 Structs — Go's "class" 🔴
- [ ] Struct definition, struct literals, **zero value is usable** (a huge design idea)
- [ ] Methods on structs: `func (b Book) Title() string`
- [ ] **Value receiver vs pointer receiver** — the rules, and when each is right
- [ ] **Method sets** — what `T` satisfies vs what `*T` satisfies (the #1 interface confusion)
- [ ] Struct tags — `json:"id"`, `db:"user_id"`, `validate:"required"`
- [ ] Anonymous structs; `struct{}` as a zero-byte signal (`map[string]struct{}`, `chan struct{}`)
- [ ] Comparable structs and `==`

## C.2 Encapsulation — the package is the boundary 🔴
- [ ] `Capitalized` = exported · `lowercase` = package-private
- [ ] **The package, not the type, is the unit of encapsulation** — this changes how you design
- [ ] `internal/` packages — enforced by the compiler
- [ ] No `Get` prefix: `b.Title()` / `b.SetTitle()`

## C.3 Construction — no constructors 🔴
- [ ] Zero value first: design types that are useful without initialization (`sync.Mutex`, `bytes.Buffer`)
- [ ] `func NewBook(title string) (*Book, error)` — the convention
- [ ] **Functional options** ⭐ — `NewServer(addr, WithTimeout(5*time.Second), WithTLS(cfg))` — Go's answer to overloaded constructors and optional params
- [ ] Builder pattern when options get complex

## C.4 Interfaces — implicit satisfaction 🔴
- [ ] `type Reader interface { Read(p []byte) (int, error) }` — **no `implements` keyword**
- [ ] Structural typing: if it has the methods, it satisfies the interface. Decoupling by default
- [ ] **Keep interfaces small** — `io.Reader` has one method. "The bigger the interface, the weaker the abstraction"
- [ ] ⭐ **Define interfaces in the consumer package**, not the producer — the inverse of Java/C#
- [ ] **Accept interfaces, return structs**
- [ ] `interface{}` / `any`; type assertion `v, ok := i.(string)`; **type switch**
- [ ] Standard interfaces: `error`, `fmt.Stringer`, `io.Reader/Writer/Closer`, `sort.Interface`, `json.Marshaler`, `driver.Valuer`
- [ ] ⚠️ **Nil interface vs interface holding a nil pointer** — a non-nil interface can hold nil and `!= nil`. Classic bug
- [ ] How it works underneath: interface value = `(type, value)` pair, dispatch via itab
- [ ] Compile-time check: `var _ Storage = (*PostgresStore)(nil)`

## C.5 Embedding — Go's composition 🔴
- [ ] **Struct embedding:** `type EBook struct { Book; FileSize int }`
- [ ] **Promoted fields and methods** — looks like inheritance, is *not*
- [ ] Why it isn't inheritance: no subtype relationship, no virtual dispatch, no `super`
- [ ] Name shadowing and ambiguity (ambiguous promotion = compile error)
- [ ] **Interface embedding:** `type ReadWriter interface { Reader; Writer }`
- [ ] Embedding an interface in a struct (partial implementation / decorator trick)
- [ ] Explicit delegation when embedding is too magic

## C.6 Polymorphism, Go-style 🔴
- [ ] Interfaces are the *only* runtime polymorphism
- [ ] No overloading → distinct names (`Print`, `Printf`, `Println`) or variadic/option structs
- [ ] No operator overloading → methods (`a.Add(b)`)
- [ ] **Generics** (1.18+): type params, type sets, constraints, `comparable`, `any`, `constraints.Ordered`
- [ ] Generics vs interfaces — when each is the right tool

## C.7 Errors, not exceptions 🔴
- [ ] `error` is just an interface: `type error interface { Error() string }`
- [ ] Sentinel errors (`var ErrNotFound = errors.New(...)`)
- [ ] Custom error types + `errors.As`
- [ ] Wrapping: `fmt.Errorf("load user: %w", err)`, `errors.Is`, `errors.Unwrap`
- [ ] `panic` / `recover` — for programmer errors and library boundaries only
- [ ] `defer` for cleanup; ⚠️ deferred-argument evaluation and loop-variable gotchas

## C.8 Go Design Philosophy 🟡
- [ ] "A little copying is better than a little dependency"
- [ ] "Clear is better than clever"
- [ ] "Don't communicate by sharing memory; share memory by communicating"
- [ ] Composition over inheritance — enforced, not suggested
- [ ] Read **Effective Go** and the **Go Proverbs**

---
---

# PART D — Translation Table
*Same idea, four spellings. This is your quick-reference page.*

| Concept | C++ | Python | TypeScript | Go |
|---|---|---|---|---|
| Type definition | `class` / `struct` | `class` | `class` | `struct` + methods |
| Receiver | implicit `this` | explicit `self` | `this` (⚠️ rebinding) | named receiver `(b *Book)` |
| Constructor | `Foo()` ctor | `__init__` | `constructor()` | `NewFoo()` convention |
| Optional ctor args | overloads / defaults | keyword defaults | optional params | **functional options** |
| Destructor | `~Foo()` / RAII | `__del__` / `with` | `dispose()` / `using` | `defer` / `Close()` |
| Privacy | `private` (real) | `_`/`__` (convention) | `private` (compile-time), `#x` (real) | lowercase (real, package-scoped) |
| Encapsulation unit | class | module (loosely) | module/class | **package** |
| Inheritance | ✅ multiple | ✅ multiple | ✅ single only | ❌ → embedding |
| Call parent | `Base::m()` | `super().m()` | `super.m()` | `e.Book.M()` |
| Abstract class | pure virtual | `ABC` | `abstract class` | ❌ |
| Interface | all-pure-virtual class | `ABC` / `Protocol` | `interface` | `interface` |
| Interface typing | nominal | nominal (`Protocol`=structural) | **structural** | **structural** |
| Declares implementation | inherit base | inherit `ABC` | `implements` (optional) | **nothing — implicit** |
| Method overloading | ✅ | ❌ | signatures only | ❌ |
| Operator overloading | ✅ | ✅ dunders | ❌ | ❌ |
| Generics | templates | `Generic[T]` | `<T>` | type params (1.18+) |
| Runtime dispatch | vtable (`virtual`) | MRO dict lookup | prototype chain | interface itab |
| Static members | `static` | class attr / `@classmethod` | `static` | package-level |
| Enum | `enum class` | `enum.Enum` | `enum` / `as const` | `const` + `iota` |
| Errors | exceptions | exceptions | exceptions | **error values** |
| Async | threads / `std::future` | `asyncio` / threads | Promise / `async-await` | **goroutines + channels** |
| Memory | manual / smart ptrs | refcount + GC | GC | GC + escape analysis |
| Deep copy | copy ctor | `deepcopy()` | `structuredClone()` | manual |
| Nullability | `nullptr` / `optional` | `None` | `null \| undefined`, strictNullChecks | `nil` (⚠️ typed-nil trap) |
| Reflection | RTTI (limited) | full | limited (`typeof`) | `reflect` |

---
---

# PART E — Design Principles
*Language-independent. These are the actual "software engineering" part.*

## E.1 SOLID 🔴
- [ ] **S — Single Responsibility** — one reason to change. Don't mix HTTP + business logic + SQL
- [ ] **O — Open/Closed** — extend by adding code, not editing existing code
- [ ] **L — Liskov Substitution** — a subtype must work anywhere the supertype does
  - Violations: `Square extends Rectangle`, strengthened preconditions, `throw NotImplemented`
- [ ] **I — Interface Segregation** — many small interfaces > one fat one (Go's `io.Reader` is the gold standard; TS: split fat interfaces)
- [ ] **D — Dependency Inversion** — depend on abstractions, not concretions
- [ ] Write a **violating** and a **fixed** version of each. That's the whole exercise

## E.2 GRASP — who owns this responsibility? 🟡
- [ ] Information Expert · Creator · Controller
- [ ] Low Coupling · High Cohesion
- [ ] Polymorphism · Pure Fabrication · Indirection · Protected Variations

## E.3 DRY 🔴
- [ ] One authoritative representation per piece of **knowledge** (not per piece of text)
- [ ] **Rule of Three** — abstract on the third duplication, not the first
- [ ] ⚠️ Over-DRY is real: **AHA — Avoid Hasty Abstractions**. A wrong abstraction costs more than duplication
- [ ] Opposite: WET

## E.4 KISS 🔴
- [ ] Simplicity is the goal. Essential vs **accidental** complexity
- [ ] Go's entire design is KISS; TypeScript's type system is where people forget it

## E.5 YAGNI 🔴
- [ ] Build for today. Premature abstraction *is* technical debt
- [ ] "Premature optimization is the root of all evil" — measure first

## E.6 Law of Demeter 🟡
- [ ] ❌ `order.customer().address().city()` → ✅ `order.shippingCity()`

## E.7 Separation of Concerns 🔴
- [ ] Handler / Service / Repository — SRP at the architecture level

## E.8 Inversion of Control & Dependency Injection 🔴
- [ ] IoC — the framework calls you ("Hollywood Principle")
- [ ] ⚠️ **DI is the *how*, DIP is the *why*** — they are not the same thing
- [ ] Constructor injection (preferred) · setter injection · interface injection
- [ ] **Composition root** — wire everything in exactly one place (`main()`, `AppModule`)
- [ ] Service Locator — and why it's an anti-pattern
- [ ] Tooling: Wire/Fx (Go) · InversifyJS, NestJS DI, tsyringe (TS) · `dependency-injector` (Python) · manual (C++)
- [ ] **DI is what makes code testable** — mocks enter through the interface

## E.9 Design by Contract 🟡
- [ ] Preconditions · postconditions · invariants
- [ ] `assert` (Python/C++), `static_assert` (C++), explicit checks (Go), type-level contracts (TS)

## E.10 Command–Query Separation 🟡
- [ ] A method either **does** (command, returns nothing) or **answers** (query, no side effects) — never both

## E.11 Program to an Interface, Not an Implementation 🔴
- [ ] Encapsulate what varies
- [ ] Go: accept interfaces, return structs · TS: type params by interface · C++: depend on abstract bases or concepts

## E.12 Tell, Don't Ask 🟡
- [ ] ❌ `if acct.balance() >= amt { acct.setBalance(acct.balance() - amt) }` → ✅ `acct.withdraw(amt)`
- [ ] Prevents **anemic domain models**

## E.13 POLA — Principle of Least Astonishment 🟡
- [ ] `delete()` deletes. It doesn't archive

## E.14 Fail Fast 🟡
- [ ] Validate at the boundary; never let bad data flow inward silently

## E.15 Also worth knowing 🟢
- [ ] **SLAP** — Single Level of Abstraction per function
- [ ] Single Source of Truth · Principle of Least Privilege
- [ ] **Boy Scout Rule** · Orthogonality · Convention over Configuration
- [ ] Postel's Law (and its critics) · **Idempotency** (critical for APIs & retries)
- [ ] Stable Dependencies · **Acyclic Dependencies** (Go enforces this; Python and TS don't)

---
---

# PART F — Design Patterns

## F.1 Creational 🟡
| Pattern | Idea | Language notes |
|---|---|---|
| **Singleton** | Exactly one instance | Go: `sync.Once` · Python: module-level · TS: module singleton · C++: Meyers singleton. ⚠️ Often an anti-pattern — global state, untestable |
| **Factory Method** | Subclass/function decides what to build | Go: plain `NewX()` funcs |
| **Abstract Factory** | Families of related objects | |
| **Builder** | Step-by-step construction | TS: fluent chaining · Go: often replaced by options |
| **Prototype** | Clone instead of construct | Python `deepcopy` · TS `structuredClone` |
| **Object Pool** | Reuse expensive objects | Go `sync.Pool`, DB conn pools |
| **Functional Options** ⭐ | Optional named params | **Go idiom** — TS/Python just use an options object |

## F.2 Structural 🟡
| Pattern | Idea | Language notes |
|---|---|---|
| **Adapter** | Bridge incompatible interfaces | Go: trivially easy via implicit interfaces |
| **Bridge** | Split abstraction from implementation | |
| **Composite** | Trees and leaves treated alike | Filesystems, React component trees |
| **Decorator** ⭐ | Wrap to add behaviour | **Go HTTP middleware** · **TS/Python decorators (`@x`)** are a *different* thing — know both meanings |
| **Facade** | Simple door to a complex subsystem | |
| **Flyweight** | Share immutable state | |
| **Proxy** | Control access (lazy, cache, remote, auth) | TS: `Proxy` object is built-in |

## F.3 Behavioral 🟡
| Pattern | Idea | Language notes |
|---|---|---|
| **Strategy** ⭐ | Interchangeable algorithms | Go/TS: often just a function type — no class needed |
| **Observer / Pub-Sub** | Notify dependents on change | TS: EventEmitter/RxJS · Go: channels |
| **Chain of Responsibility** ⭐ | Pass through handlers | Middleware everywhere |
| **Command** | Request as an object | Undo, job queues |
| **State** | Behaviour changes with state | State machines |
| **Template Method** | Skeleton + overridable steps | Go: function fields instead of inheritance |
| **Iterator** | Sequential access | Python generators · TS `Symbol.iterator` · C++ iterators · Go 1.23 range-over-func |
| **Mediator** | Centralize many-to-many talk | |
| **Memento** | Capture/restore state | |
| **Visitor** | Add ops without editing the hierarchy | TS: discriminated unions are usually better |
| **Null Object** | Do-nothing impl instead of null checks | |
| **Interpreter** ⚪ | | |

## F.4 Anti-Patterns 🟡
- [ ] **God Object / God Function**
- [ ] **Anemic Domain Model** — all data, no behaviour (epidemic in TS/Go backends)
- [ ] Singleton abuse / global mutable state
- [ ] Spaghetti code · **Big Ball of Mud**
- [ ] Golden Hammer ("everything is a microservice")
- [ ] Premature abstraction / speculative generality
- [ ] **Circular dependencies** — Go: compile error · Python & TS: silent runtime pain
- [ ] Magic numbers & magic strings
- [ ] **Primitive obsession** — `string userID` everywhere instead of a `UserID` type (TS: branded types; Go: `type UserID string`)
- [ ] Callback hell / promise nesting (TS)
- [ ] `any` everywhere (TS) · bare `except:` (Python) · `interface{}` everywhere (Go)
- [ ] Yo-yo problem · Lava flow · Poltergeist

---
---

# PART G — Concurrency Patterns 🔴
*Especially critical for Go backends, but every language has these problems.*

- [ ] **Concurrency vs parallelism** — not the same thing
- [ ] Models compared:
  - **Go:** goroutines + channels (CSP) — cheap, preemptive
  - **TypeScript:** single-threaded event loop, Promises/`async-await`, Web Workers; **no shared-memory races** by default
  - **Python:** the **GIL** — threads for I/O, `multiprocessing` for CPU, `asyncio` for concurrency
  - **C++:** `std::thread`, `std::async`, `std::atomic`, `std::mutex`
- [ ] **Producer–Consumer** (queues/channels)
- [ ] **Worker Pool** — N workers draining one job queue
- [ ] **Fan-out / Fan-in**
- [ ] **Pipeline** — stages connected by channels/streams
- [ ] **Semaphore / rate limiting** — buffered channels, `x/time/rate`, `p-limit` (TS)
- [ ] **Future / Promise** — `std::future` · `asyncio.Future` · `Promise` · channel-returning func
- [ ] **Cancellation & timeouts** — `context.Context` (Go) · `AbortController` (TS) · `asyncio.CancelledError`
- [ ] **Mutex / RWMutex**, `sync.Once`, `WaitGroup`, `errgroup` · `Promise.all/allSettled/race` (TS)
- [ ] **Actor model** · double-checked locking ⚪
- [ ] Thread safety in OOP: immutability, confinement, copy-on-write
- [ ] Race conditions, deadlock, livelock, starvation — `go test -race`, TSan (C++)

---
---

# PART H — Architecture

## H.1 Layering 🔴
- [ ] Layered / N-tier: Handler → Service → Repository → DB
- [ ] **Hexagonal / Ports & Adapters**
- [ ] **Clean / Onion Architecture** — dependencies point inward only
- [ ] The dependency rule: domain must never import infrastructure

## H.2 Structural patterns 🟡
- [ ] **MVC** · MVP · **MVVM** (Angular/Vue) · component architecture (React)
- [ ] **Repository** — data access behind an interface
- [ ] Unit of Work · Service Layer
- [ ] **DTO** (transport) vs **DAO** (access) vs **Entity** (domain)
- [ ] Active Record (Django ORM, TypeORM) vs Data Mapper (SQLAlchemy, Prisma)
- [ ] Gateway/Adapter for third-party APIs · Anti-Corruption Layer

## H.3 Domain-Driven Design basics 🟢
- [ ] Ubiquitous language
- [ ] **Entity** (identity) vs **Value Object** (equality by value, immutable)
- [ ] **Aggregate** + aggregate root — the consistency boundary
- [ ] Repository · Domain Service · Domain Event
- [ ] Bounded Context · Context Map

## H.4 Modules, Packages & Dependencies 🔴
- [ ] Public API surface design · **semantic versioning** · backward compatibility
- [ ] **Go:** packages, `internal/`, no import cycles (compiler-enforced)
- [ ] **Python:** modules/packages, `__init__.py`, circular imports
- [ ] **TypeScript:** ES modules, barrel files (`index.ts`) and their pitfalls, path aliases, `tsconfig` project references
- [ ] **C++:** headers vs impl, include guards, **PIMPL** for compile-time decoupling, C++20 modules
- [ ] The dependency graph *is* your architecture

## H.5 Distributed / Backend patterns 🟢
- [ ] **CQRS** · **Event Sourcing** · Event-Driven Architecture
- [ ] **Saga** · **Outbox** pattern
- [ ] **Circuit Breaker** · **Retry with exponential backoff + jitter** · Bulkhead · Timeout
- [ ] **Idempotency keys** · Sidecar · API Gateway · BFF
- [ ] Monolith → modular monolith → microservices, and when *not* to split

---
---

# PART I — Errors, Testing & Code Quality

## I.1 Error Handling as Design 🔴
- [ ] **Exceptions vs error values** — the fundamental fork
- [ ] **C++:** `try/catch/throw`, exception-safety guarantees (basic/strong/nothrow), `noexcept`, RAII cleanup, `std::expected` (C++23)
- [ ] **Python:** exception hierarchy, custom exception classes, `try/except/else/finally`, `raise ... from`, context managers
- [ ] **TypeScript:** ⚠️ `throw` is untyped — anything can be thrown; `catch (e: unknown)`; custom `Error` subclasses + `instanceof`; **Result/Either types** as a typed alternative; ⚠️ unhandled promise rejections
- [ ] **Go:** `error` interface, sentinels, wrapping with `%w`, `errors.Is` / `errors.As`, `panic`/`recover`
- [ ] Error wrapping & context; error boundaries; recoverable vs programmer errors
- [ ] Never swallow errors silently

## I.2 Testing & Testability 🔴
- [ ] Unit / integration / end-to-end; the test pyramid
- [ ] **Test doubles:** dummy · stub · fake · spy · **mock** — know the differences
- [ ] Mocking through interfaces — the payoff of DIP + DI
- [ ] Per language:
  - **Go:** `testing`, **table-driven tests** ⭐, `httptest`, `gomock`/`testify`, `-race`, `-fuzz`, benchmarks
  - **Python:** `pytest`, fixtures, `parametrize`, `unittest.mock`, `monkeypatch`, Hypothesis
  - **TypeScript:** Jest/Vitest, `jest.mock`, spies, MSW for HTTP, Playwright for E2E
  - **C++:** GoogleTest/Catch2, GMock
- [ ] **TDD** red → green → refactor · BDD given/when/then
- [ ] Coverage and its limits; mutation testing ⚪
- [ ] **Designing for testability:** no globals, inject clock/random/IO, pure functions at the core

## I.3 Coupling & Cohesion 🔴
- [ ] Tight vs loose coupling; afferent/efferent coupling
- [ ] High vs low cohesion
- [ ] Connascence (name, type, position, algorithm, timing) ⚪

## I.4 Code Smells 🟡
- [ ] Long function · long parameter list · large class
- [ ] Deep nesting → early returns / guard clauses
- [ ] Feature envy · shotgun surgery · divergent change
- [ ] Data clumps · primitive obsession · magic numbers
- [ ] Dead code · speculative generality · comments that excuse bad code

## I.5 Refactoring Techniques 🟡
- [ ] Extract Function / Class / Interface
- [ ] Introduce Parameter Object · Replace Magic Number with Constant
- [ ] **Replace Conditional with Polymorphism** ⭐
- [ ] Replace Inheritance with Delegation
- [ ] Encapsulate Field / Collection · Decompose Conditional
- [ ] **Strangler Fig** — incremental rewrite

## I.6 Professional Practice 🟢
- [ ] Naming as design; intention-revealing names
- [ ] Comments explain **why**, not what. Doc tooling: `godoc` · docstrings · TSDoc · Doxygen
- [ ] Code review · style guides · linters: `golangci-lint` · `ruff`+`mypy` · `eslint`+`tsc --strict` · `clang-tidy`
- [ ] Logging, metrics, tracing (observability)
- [ ] Technical debt — deliberate vs accidental

---
---

# PART J — Per-Language Deep Dives

## J.1 C++ 🟡
- [ ] `virtual` / `override` / `final`, vtable layout, **object slicing**
- [ ] **Virtual destructors**; abstract base classes
- [ ] `friend`, `explicit`, `const` correctness, `mutable`, `static`
- [ ] **RAII + smart pointers** (`unique_ptr`, `shared_ptr`, `weak_ptr`) — ownership as design
- [ ] Rule of 0/3/5; move semantics; copy elision
- [ ] Templates, specialization, **CRTP**, **concepts** (C++20)
- [ ] Multiple & virtual inheritance; the diamond
- [ ] `dynamic_cast` / RTTI / `typeid`
- [ ] **PIMPL** idiom
- [ ] Operator overloading rules; `<=>` spaceship operator
- [ ] Exception safety guarantees

## J.2 Python 🟡
- [ ] Everything is an object; `type` and `object`
- [ ] **Dunders:** `__init__`, `__new__`, `__str__`, `__repr__`, `__eq__`, `__hash__`, `__len__`, `__call__`, `__iter__`, `__getitem__`, `__enter__`/`__exit__`
- [ ] ⚠️ `__eq__` without `__hash__` breaks dict/set usage
- [ ] `@property`, **descriptors** (`__get__`/`__set__`/`__set_name__`)
- [ ] `@staticmethod` vs `@classmethod` vs instance method
- [ ] **MRO / C3 linearization**, cooperative `super()`
- [ ] Mixins in practice
- [ ] `abc.ABC` + `@abstractmethod`; **`typing.Protocol`** (structural)
- [ ] `@dataclass`, `NamedTuple`, `attrs`, **Pydantic**
- [ ] `__slots__` — memory + attribute locking
- [ ] **Decorators** (function & class) as composition
- [ ] Metaclasses, `__init_subclass__` ⚪
- [ ] Context managers, generators, iterators
- [ ] Duck typing; **EAFP vs LBYL**
- [ ] Type hints, `mypy`, `TypeVar`, generics
- [ ] The **GIL** and what it means for your design

## J.3 TypeScript 🟡
- [ ] It's **JavaScript + types**: prototypes underneath, classes are sugar
- [ ] ⚠️ **`this` binding** — arrow-function class fields vs `bind`; the #1 TS OOP bug
- [ ] `private`/`protected`/`readonly` are **compile-time only**; `#field` is real
- [ ] Parameter properties: `constructor(private readonly repo: Repo) {}`
- [ ] `interface` vs `type` vs `abstract class` — when to use each
- [ ] **Structural typing** — `implements` is optional and only a check
- [ ] `implements` multiple interfaces; **mixin factory functions** for multiple inheritance
- [ ] Generics: constraints (`extends`), `keyof`, conditional types, mapped types, `infer`
- [ ] Utility types: `Partial`, `Required`, `Readonly`, `Pick`, `Omit`, `Record`, `Returntype`
- [ ] **Discriminated unions + exhaustive `switch` with `never`** — often replaces polymorphism entirely
- [ ] Type guards, `is` predicates, `satisfies`
- [ ] Decorators (`@Injectable`, `@Component`) — NestJS/Angular style; `reflect-metadata`
- [ ] `strict` mode, `strictNullChecks` — non-negotiable
- [ ] ⚠️ `any` vs `unknown` vs `never`
- [ ] **Branded/nominal types** to fight primitive obsession
- [ ] Modules, `async/await`, the event loop
- [ ] Getters/setters, `static` blocks, `enum` vs `as const`

## J.4 Go 🔴
*See Part C — the whole of Part C is the Go deep dive.* Additionally:
- [ ] Slices, maps, and their aliasing behaviour in structs
- [ ] `defer` evaluation order and gotchas
- [ ] Generics vs interfaces — choosing correctly
- [ ] `reflect` and struct tags (how `encoding/json` actually works)
- [ ] `go vet`, `golangci-lint`, race detector
- [ ] Effective Go · Go Proverbs · *100 Go Mistakes*

---
---

# PART K — Study Plan

| Phase | Topics | Priority | Est. |
|---|---|---|---|
| 1 | **Part A** — universal concepts, object model, lifecycle, type systems | 🔴 | 4 d |
| 2 | **Part B.1–B.4** — classes, ctors, dtors, encapsulation in C++/Py/TS | 🔴 | 1 wk |
| 3 | **Part B.5–B.7** — inheritance, abstraction, polymorphism | 🔴 | 1 wk |
| 4 | **Part C entirely** — Go's model, in one focused pass | 🔴 | 1 wk |
| 5 | **Part D** — build the translation table yourself from memory | 🔴 | 1 d |
| 6 | **E.1 SOLID** — a violating + fixed example for each, in 2 languages | 🔴 | 1 wk |
| 7 | **E.3–E.7** — DRY, KISS, YAGNI, Demeter, SoC | 🔴 | 3 d |
| 8 | **E.8 DI/IoC** + **I.2 testing & mocks** — learn these together | 🔴 | 1 wk |
| 9 | **I.1** — error handling as design (all 4 models) | 🔴 | 3 d |
| 10 | **F.1–F.3 core patterns** — Strategy, Factory, Decorator, Observer, Adapter, Repository, Functional Options | 🟡 | 2 wk |
| 11 | **Part G** — concurrency patterns (heaviest for Go) | 🔴 | 1 wk |
| 12 | **H.1–H.2** — layering, Clean/Hexagonal, Repository, DTO | 🟡 | 1 wk |
| 13 | **F.4 anti-patterns** + **I.4–I.5** smells & refactoring | 🟡 | 1 wk |
| 14 | Remaining GoF patterns | 🟡 | 2 wk |
| 15 | **E.2 GRASP, E.9–E.15** misc principles | 🟢 | 4 d |
| 16 | **H.3 DDD** + **H.5** distributed patterns | 🟢 | 2 wk |
| 17 | **Part J** — deep dive whichever language you ship in | 🟡 | ongoing |

## How to actually learn this (not just read it)
1. **Four-file rule** — for each concept in Parts A–C, write the same tiny program in C++, Python, TypeScript and Go. Note what each language *can't* do and how it compensates. That gap is where the real understanding is.
2. **Violate then fix** — for every principle in Part E, write the broken version first. You don't understand SRP until you've felt the mess it prevents.
3. **One project that forces all of it** — a REST API with: handler/service/repository layers · interfaces for storage · DI wired in `main()` · table-driven tests with mocks · a worker pool for background jobs · a middleware chain · wrapped errors · graceful shutdown with context. That single project exercises A, B/C, E.1, E.8, F, G, H.1, I.1, I.2.
4. **Then build it twice** — once in Go, once in TypeScript. The differences will teach you more than any book.

## 60-second interview answers you should own
Composition vs inheritance · vtable & dynamic dispatch · diamond problem & MRO · an LSP violation · interface vs abstract class · DI vs DIP · Go's nil-interface trap · value vs pointer receiver · shallow vs deep copy · why Singleton is usually a mistake · structural vs nominal typing · Python's GIL · TypeScript's `this` binding · exceptions vs error values · idempotency

## Reading list
- *Clean Code* & *Clean Architecture* — Robert C. Martin
- *Refactoring* (2nd ed., JS examples — perfect for the TS track) — Martin Fowler
- *Design Patterns* — GoF (reference, not cover-to-cover)
- *Head First Design Patterns* — the readable intro
- *The Pragmatic Programmer* — Hunt & Thomas
- **Go:** *Effective Go* · Go Proverbs · *100 Go Mistakes and How to Avoid Them* · *Learning Go* (Bodner)
- **C++:** *Effective Modern C++* — Meyers
- **Python:** *Fluent Python* — Ramalho
- **TypeScript:** *Effective TypeScript* — Dan Vanderkam · *Programming TypeScript* — Boris Cherny
- *Domain-Driven Design Distilled* — Vaughn Vernon
- **refactoring.guru** — best free pattern catalogue, code in all four languages
