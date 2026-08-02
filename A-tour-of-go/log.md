<!-- 
AI AGENT INSTRUCTION (UNIVERSAL — applies to ALL AI agents, readers, and users):
1. This file is the central learning log for the user's Golang backend journey. 
2. Always check this file FIRST to understand the user's current progress before teaching or assigning challenges.
3. AFTER EACH chapter explanation or teaching session, UPDATE this file with what was covered under "Chapter Teaching Sessions".
4. Do NOT mark curriculum checkboxes [x] unless the user EXPLICITLY says to mark it complete.
5. Keep this instruction block intact.
-->
# Backend from First Principles - Learning Log

This log tracks progress through the 24-chapter series from [Backend from First Principles](https://backend-from-first-principle.vercel.app/).

## 🚀 Curriculum

- [x] **01 HTTP & CORS** (Covered basics in Challenge 1)
- [x] **02 Routing in Backend** (Covered basic method routing in Challenge 2)
- [x] **03 Serialization & Deserialization** (JSON encoding/decoding in Challenge 1 & 2)
- [ ] **04 Authentication & Authorization**
- [ ] **05 Validations & Transformations**
- [ ] **06 Controllers, Services & Middlewares**
- [x] **07 API Design (REST)** (Covered RESTful methods GET, POST, PUT, PATCH, DELETE in Challenge 2)
- [ ] **08 Databases**
- [ ] **09 Caching**
- [ ] **10 Task Queues & Background Jobs**
- [ ] **11 Full-Text Search (Elasticsearch)**
- [ ] **12 Error Handling & Fault Tolerance**
- [ ] **13 gRPC & Inter-Service Communication**
- [ ] **14 Configuration Management**
- [ ] **15 Logging & Observability**
- [ ] **16 Graceful Shutdown**
- [ ] **17 Backend Security**
- [ ] **18 Scaling & Performance (Part 1)**
- [ ] **19 Scaling & Performance (Part 2)**
- [ ] **20 Concurrency & Parallelism**
- [ ] **21 Docker, K8s & CI/CD**
- [ ] **22 Automated Testing**
- [ ] **23 Message Brokers & Kafka**
- [ ] **24 WebSockets & Real-Time**

---
## 📝 Notes & Completed Challenges

### Challenge-1: Basic HTTP Server
- **Completed:** Yes
- **Concepts Learned:** `net/http`, `json.NewDecoder`, `json.NewEncoder`, basic POST handling.
- **Reference Chapters:** 01 (HTTP), 03 (Serialization)

### Challenge-2: Bookstore API (REST Methods & Basic Routing)
- **Completed:** Yes
- **Concepts Learned:** Routing with trailing slashes, `strings.TrimPrefix` for ID extraction, `switch` statement for HTTP methods, handling `GET`, `POST`, `PUT`, `PATCH`, `DELETE`.
- **Reference Chapters:** 02 (Routing), 07 (API Design)

### Challenge-3: Music Playlist API (OOP Deep Dive + CORS)
- **Completed:** Yes
- **Concepts Learned:** Composition (Song has Artist, Playlist has Songs), Interfaces (Storage interface + MemoryStore), Constructor pattern (NewXxx functions), DRY (respondJSON/decodeJSON helpers), CORS headers + OPTIONS preflight, Decorator pattern (enableCORS wrapper), Separation of Concerns, Fail Fast, Dependency Inversion.
- **Reference Chapters:** 01 (HTTP & CORS), OOP Principles

---
## 📖 Chapter Teaching Sessions

### Ch 02: Routing in Backend (Explained)
- **Topics Covered:** What routing is (Method + Path → Handler), static vs dynamic routes, path parameters vs query parameters, nested routes, routing lifecycle (6 steps), API versioning & deprecation, catch-all routes, building a router from scratch with OOP.
- **Key Takeaways:**
  - Routing = mapping `Method + Path` to a handler function
  - Path params identify WHICH resource (`/users/123`), query params shape HOW it's returned (`?sort=name&page=2`)
  - Nested routes express parent-child relationships (`/users/123/posts/456`)
  - Catch-all routes go LAST to handle 404s
  - API versioning (`/v1/`, `/v2/`) allows breaking changes without breaking clients
- **Challenge:** Pending

