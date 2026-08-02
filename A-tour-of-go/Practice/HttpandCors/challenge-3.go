
/*
╔══════════════════════════════════════════════════════════════════╗
║              CHALLENGE 3: Music Playlist API                     ║
║              OOP Concepts Deep Dive + CORS                       ║
╚══════════════════════════════════════════════════════════════════╝

GOAL:
Build a Music Playlist API where users can manage Artists, Songs,
and Playlists. This challenge forces you to use MORE OOP concepts
than Challenge 2.

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

NEW CONCEPTS YOU WILL LEARN:
1. COMPOSITION (struct embedding — "Has-a" relationship)
2. INTERFACES (define your own interface)
3. CONSTRUCTOR PATTERN (NewXxx functions — Go's replacement for constructors)
4. DRY (extract repeated code into helper functions)
5. CORS (add proper CORS headers — we ARE in the HTTP & CORS chapter!)
6. SEPARATION OF CONCERNS (split storage logic from HTTP logic)

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

STEP 1: Define these structs
────────────────────────────

type Artist struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}

type Song struct {
    ID       string  `json:"id"`
    Title    string  `json:"title"`
    Artist   Artist  `json:"artist"`    ← COMPOSITION! Song "has an" Artist
    Duration int     `json:"duration"`  ← in seconds
    Genre    string  `json:"genre"`
}

type Playlist struct {
    ID          string `json:"id"`
    Name        string `json:"name"`
    Description string `json:"description"`
    Songs       []Song `json:"songs"`   ← COMPOSITION! Playlist "has" Songs
    CreatedAt   string `json:"created_at"`
}

Notice how Song doesn't "inherit" from Artist.
Instead, Song HAS an Artist inside it. That's composition.
And Playlist HAS a slice of Songs. More composition.

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

STEP 2: Constructor Pattern (NewXxx functions)
──────────────────────────────────────────────

Go has no constructors. Instead, you write factory functions
that start with "New":

func NewSong(title string, artist Artist, duration int, genre string) Song {
    // assign an auto-incremented ID
    // return the fully initialized Song
}

func NewPlaylist(name string, description string) Playlist {
    // assign an auto-incremented ID
    // set CreatedAt to current time (use time.Now().Format(time.RFC3339))
    // initialize Songs as an empty slice
    // return the Playlist
}

WHY? This ensures every Song/Playlist is created with valid data.
This is ENCAPSULATION — you control how objects are created.

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

STEP 3: Define your OWN Interface
─────────────────────────────────

This is the big one. Define a Storage interface:

type Storage interface {
    GetAll() []Playlist
    GetByID(id string) (Playlist, bool)
    Create(p Playlist) Playlist
    Update(id string, p Playlist) (Playlist, bool)
    Delete(id string) bool
}

Then create a struct that IMPLEMENTS this interface:

type MemoryStore struct {
    playlists map[string]Playlist
}

Implement all 5 methods on MemoryStore.
Because MemoryStore has all the methods that Storage requires,
it AUTOMATICALLY satisfies the Storage interface.
This is POLYMORPHISM — you could later swap MemoryStore
for a PostgresStore without changing your HTTP handler code!

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

STEP 4: DRY — Helper Functions
──────────────────────────────

In Challenge 2, you wrote the same JSON decode + error handling
code 3 times. This time, create helper functions:

func respondJSON(w http.ResponseWriter, status int, data interface{}) {
    // set content-type header
    // set status code
    // encode data as JSON
}

func decodeJSON(r *http.Request, target interface{}) error {
    // decode request body into target
    // return error if it fails
}

func respondError(w http.ResponseWriter, status int, message string) {
    // respond with {"error": "message"} as JSON
}

Now your handler code becomes super clean:
    respondJSON(w, http.StatusOK, playlists)
Instead of:
    w.Header().Set("content-type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(playlists)

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

STEP 5: CORS Headers
─────────────────────

Add a CORS middleware/wrapper. Before handling any request,
your server should set these headers:

    Access-Control-Allow-Origin: *
    Access-Control-Allow-Methods: GET, POST, PUT, DELETE, OPTIONS
    Access-Control-Allow-Headers: Content-Type

Also handle the OPTIONS method (preflight request):
    If r.Method == "OPTIONS" {
        w.WriteHeader(http.StatusOK)
        return
    }

WHY? Without CORS headers, a frontend running on localhost:3000
cannot call your API running on localhost:8080. The browser blocks it.

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

STEP 6: API Endpoints
─────────────────────

Implement these routes:

PLAYLISTS:
  GET    /playlists/          → Get all playlists
  GET    /playlists/{id}      → Get one playlist by ID
  POST   /playlists/          → Create a new playlist
  PUT    /playlists/{id}      → Update a playlist (name, description)
  DELETE /playlists/{id}      → Delete a playlist

SONGS IN A PLAYLIST:
  POST   /playlists/{id}/songs/   → Add a song to a playlist
  DELETE /playlists/{id}/songs/{songID}  → Remove a song from a playlist

The song endpoints are a bonus challenge — they require you to
parse nested URL paths (two IDs from one URL).

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

TEST WITH CURL:
───────────────

# Create a playlist
curl -X POST http://localhost:8080/playlists/ \
  -H "Content-Type: application/json" \
  -d '{"name": "Road Trip", "description": "Songs for a long drive"}'

# Get all playlists
curl http://localhost:8080/playlists/

# Add a song to playlist 1
curl -X POST http://localhost:8080/playlists/1/songs/ \
  -H "Content-Type: application/json" \
  -d '{"title": "Bohemian Rhapsody", "artist": {"name": "Queen"}, "duration": 354, "genre": "Rock"}'

# Test CORS preflight
curl -X OPTIONS http://localhost:8080/playlists/ -v

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

OOP CONCEPTS SCORECARD (check these off when done):
───────────────────────────────────────────────────

[ ] ENCAPSULATION      — lowercase fields/variables are private
[ ] ABSTRACTION        — Storage interface hides implementation details
[ ] POLYMORPHISM       — MemoryStore satisfies Storage interface implicitly
[ ] COMPOSITION        — Song has Artist, Playlist has Songs
[ ] CONSTRUCTOR        — NewSong(), NewPlaylist() factory functions
[ ] DRY                — respondJSON(), decodeJSON() helper functions
[ ] FAIL FAST          — return error immediately on bad input
[ ] SRP                — handler logic separated from storage logic
[ ] CORS               — proper cross-origin headers on all responses

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

HINTS:
- You will need these imports: "encoding/json", "fmt", "log",
  "net/http", "strconv", "strings", "time"
- For the song endpoints, you'll need to parse paths like
  "/playlists/1/songs/3" — use strings.Split() or strings.TrimPrefix()
- Remember: your handler function signature should accept Storage
  as a parameter (or use a global variable for now)

START CODING BELOW THIS LINE ↓
*/


package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
// STEP 1: STRUCTS — with COMPOSITION ("Has-a" relationship)
// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
//
// In C++ or Java, you might make Song "extend" or "inherit" from Artist.
// In Go, there is NO inheritance. Instead we use COMPOSITION.
//
// Composition means: one struct CONTAINS another struct as a field.
// It's a "Has-a" relationship:
//   - A Song HAS an Artist (not "Song IS an Artist")
//   - A Playlist HAS Songs (not "Playlist IS a Song")

// Artist is the simplest struct. It has its own identity.
type Artist struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Song COMPOSES Artist inside it.
// When you create a Song, you MUST also provide an Artist.
// This is composition — Song doesn't inherit from Artist,
// it literally contains one as a field.
//
// In JSON, this looks like:
// {
//   "id": "1",
//   "title": "Bohemian Rhapsody",
//   "artist": { "name": "Queen" },     ← nested object!
//   "duration": 354,
//   "genre": "Rock"
// }
type Song struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Artist   Artist `json:"artist"`   // COMPOSITION: Song "has an" Artist
	Duration int    `json:"duration"` // in seconds
	Genre    string `json:"genre"`
}

// Playlist COMPOSES a slice of Songs inside it.
// This is composition again — a Playlist "has" many Songs.
// If you delete a Playlist, its Songs list goes away too.
// (This is the strict form of composition — child can't exist without parent)
type Playlist struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Songs       []Song `json:"songs"`      // COMPOSITION: Playlist "has" Songs
	CreatedAt   string `json:"created_at"`
}

// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
// STEP 2: CONSTRUCTOR PATTERN — NewXxx() functions
// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
//
// Go has NO constructors (no __init__, no constructor(), no MyClass()).
// Instead, Go uses a CONVENTION: create a function named NewXxx()
// that returns a fully initialized struct.
//
// WHY use this pattern?
// - ENCAPSULATION: You control exactly how objects are created.
// - No one can create a Song with a missing ID or an empty Artist.
// - If creation rules change later, you only change ONE function.
//
// This is also the FACTORY pattern from design patterns.

// nextSongID and nextPlaylistID are lowercase = ENCAPSULATED.
// No other package can touch these. Only our NewXxx functions use them.
var nextSongID = 1
var nextPlaylistID = 1

// NewSong is the constructor for Song.
// It auto-generates an ID and returns a fully valid Song.
// The caller doesn't need to know HOW IDs are generated — that's ABSTRACTION.
func NewSong(title string, artist Artist, duration int, genre string) Song {
	song := Song{
		ID:       strconv.Itoa(nextSongID),
		Title:    title,
		Artist:   artist,   // Composition in action — we embed the Artist
		Duration: duration,
		Genre:    genre,
	}
	nextSongID++
	return song
}

// NewPlaylist is the constructor for Playlist.
// It auto-generates an ID, sets CreatedAt to current time,
// and initializes Songs as an empty slice (not nil).
//
// Why []Song{} instead of nil?
// Because when we convert to JSON:
//   nil   → "songs": null       ← ugly, frontend has to check for null
//   []Song{} → "songs": []      ← clean, frontend gets an empty array
func NewPlaylist(name string, description string) Playlist {
	playlist := Playlist{
		ID:          strconv.Itoa(nextPlaylistID),
		Name:        name,
		Description: description,
		Songs:       []Song{},                          // empty, not nil
		CreatedAt:   time.Now().Format(time.RFC3339),   // "2026-08-02T12:30:00Z"
	}
	nextPlaylistID++
	return playlist
}

// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
// STEP 3: INTERFACE — Defining your own abstraction
// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
//
// An interface defines WHAT something can do, not HOW it does it.
// It's pure ABSTRACTION.
//
// Storage defines the contract: "anything that can store playlists
// must have these 5 methods."
//
// WHY is this powerful?
// Right now we implement MemoryStore (stores data in a map).
// Later, you could create PostgresStore, MongoStore, RedisStore.
// As long as they have the same 5 methods, your HTTP handler
// code doesn't change AT ALL. That's the power of interfaces.
//
// This is also the DEPENDENCY INVERSION PRINCIPLE from SOLID:
// Your handler depends on the Storage interface (abstraction),
// not on MemoryStore directly (implementation).

type Storage interface {
	GetAll() []Playlist
	GetByID(id string) (Playlist, bool)
	Create(p Playlist) Playlist
	Update(id string, p Playlist) (Playlist, bool)
	Delete(id string) bool
}

// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
// STEP 3b: IMPLEMENTING the Interface — MemoryStore
// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
//
// MemoryStore is a struct that stores playlists in a map (in RAM).
// It implements ALL 5 methods of the Storage interface.
//
// POLYMORPHISM: Because MemoryStore has all the methods that
// Storage requires, it AUTOMATICALLY satisfies the interface.
// There is no "implements" keyword in Go. This is called
// "implicit interface satisfaction" or "duck typing":
//   "If it walks like a duck and quacks like a duck, it's a duck."
//
// In C++: class MemoryStore : public Storage { ... };
// In Python: class MemoryStore(Storage): ...
// In Go: just implement the methods. Done. No keyword needed.

type MemoryStore struct {
	playlists map[string]Playlist // lowercase = ENCAPSULATED, private
}

// NewMemoryStore is the constructor for MemoryStore.
// It initializes the map so it's ready to use.
// Without this, the map would be nil and crash on first write.
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		playlists: make(map[string]Playlist),
	}
}

// GetAll returns every playlist in the store as a slice.
// This is a METHOD on MemoryStore. The (s *MemoryStore) part
// is called a "receiver" — it's Go's version of "this" in C++
// or "self" in Python.
func (s *MemoryStore) GetAll() []Playlist {
	all := []Playlist{} // empty slice, not nil (clean JSON output)
	for _, p := range s.playlists {
		all = append(all, p)
	}
	return all
}

// GetByID returns a playlist and true if found, or empty + false if not.
// The (Playlist, bool) return pattern is idiomatic Go —
// it's like the "comma ok" pattern you see in map lookups.
func (s *MemoryStore) GetByID(id string) (Playlist, bool) {
	p, exists := s.playlists[id]
	return p, exists
}

// Create stores a new playlist. We pass it the already-constructed
// playlist (created via NewPlaylist), so this method just saves it.
// SEPARATION OF CONCERNS: this method only handles storage,
// not ID generation or timestamp — that's NewPlaylist's job.
func (s *MemoryStore) Create(p Playlist) Playlist {
	s.playlists[p.ID] = p
	return p
}

// Update replaces an existing playlist's name and description,
// but keeps the original ID, Songs, and CreatedAt.
// Returns the updated playlist and true, or empty + false if not found.
func (s *MemoryStore) Update(id string, p Playlist) (Playlist, bool) {
	existing, exists := s.playlists[id]
	if !exists {
		return Playlist{}, false // FAIL FAST: not found, return immediately
	}
	// Only update the fields that make sense to change
	existing.Name = p.Name
	existing.Description = p.Description
	s.playlists[id] = existing
	return existing, true
}

// Delete removes a playlist. Returns true if it existed, false if not.
func (s *MemoryStore) Delete(id string) bool {
	_, exists := s.playlists[id]
	if !exists {
		return false
	}
	delete(s.playlists, id)
	return true
}

// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
// STEP 4: DRY — Helper functions to avoid repeating yourself
// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
//
// In Challenge 2, you wrote this pattern over and over:
//    w.Header().Set("content-type", "application/json")
//    w.WriteHeader(status)
//    json.NewEncoder(w).Encode(data)
//
// And this error pattern:
//    http.Error(w, "some message", http.StatusBadRequest)
//
// DRY says: if you write the same code more than once, extract it
// into a function. Now you write it ONCE and call it everywhere.
//
// Notice: the "data" parameter is type "interface{}" (also written as "any").
// This means it accepts ANY type — a Playlist, a []Playlist, a Song,
// an error message struct, anything. This is POLYMORPHISM through
// Go's empty interface.

// respondJSON sends a JSON response with the given status code.
// Used for ALL successful responses.
func respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// decodeJSON reads the request body and decodes JSON into target.
// "target" is interface{} so it works with ANY struct — Song, Playlist, etc.
// Returns an error if the JSON is invalid.
//
// ABSTRACTION: the caller doesn't need to know about json.NewDecoder,
// r.Body, or io.Reader. They just call decodeJSON(r, &myStruct).
func decodeJSON(r *http.Request, target interface{}) error {
	return json.NewDecoder(r.Body).Decode(target)
}

// respondError sends a JSON error response like: {"error": "message"}
// Instead of the plain text that http.Error() sends.
// This keeps our API responses consistent — always JSON.
func respondError(w http.ResponseWriter, status int, message string) {
	respondJSON(w, status, map[string]string{"error": message})
}

// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
// STEP 5: CORS — Cross-Origin Resource Sharing
// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
//
// WHAT IS CORS?
// When your frontend (React app on localhost:3000) tries to call
// your backend API (on localhost:8080), the browser BLOCKS the request
// because they are on different "origins" (different ports = different origin).
//
// To allow it, your server must send special headers saying:
// "Hey browser, it's okay, I accept requests from other origins."
//
// HOW IT WORKS:
// 1. Browser first sends a "preflight" request (method = OPTIONS)
//    asking: "Am I allowed to send a POST/PUT/DELETE to this server?"
// 2. Server responds with CORS headers saying "Yes, these methods are allowed"
// 3. Browser then sends the actual request
//
// enableCORS is a function that takes an http.HandlerFunc and returns
// a NEW http.HandlerFunc that adds CORS headers before calling the original.
//
// This is the DECORATOR PATTERN from design patterns:
// You "wrap" the original function with extra behavior (CORS headers)
// without modifying the original function itself.
// It's also the OPEN/CLOSED PRINCIPLE: the original handler is
// CLOSED for modification but OPEN for extension (we add CORS on top).

func enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers on EVERY response
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight request
		// The browser sends OPTIONS before the real request to check permissions.
		// We respond with 200 OK and the CORS headers above, then STOP.
		// No need to process the actual request body.
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// If it's not OPTIONS, call the actual handler
		next(w, r)
	}
}

// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
// STEP 6: HTTP HANDLERS — The actual API logic
// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
//
// SEPARATION OF CONCERNS:
// - The handlers ONLY deal with HTTP (parsing requests, sending responses)
// - The Storage interface deals with data (saving, reading, deleting)
// - The helpers deal with JSON encoding/decoding
// Each layer has ONE job. That's SRP (Single Responsibility Principle).
//
// Notice how the handler uses "store Storage" (the interface),
// NOT "store *MemoryStore" (the concrete type).
// This is DEPENDENCY INVERSION — we depend on the abstraction.

// store is our global storage instance.
// It's typed as Storage (the interface), not *MemoryStore.
// This means we could swap it for a PostgresStore later
// without changing ANY handler code. That's the power of interfaces.
var store Storage

// handlePlaylists handles all /playlists/ routes
func handlePlaylists(w http.ResponseWriter, r *http.Request) {
	// Parse the path to extract playlist ID and possible sub-routes
	// Examples:
	//   /playlists/        → path = ""        → list all or create
	//   /playlists/1       → path = "1"       → get/update/delete playlist 1
	//   /playlists/1/songs/ → path = "1/songs/" → add song to playlist 1
	path := strings.TrimPrefix(r.URL.Path, "/playlists/")
	parts := strings.Split(path, "/")

	// Check if this is a song sub-route: /playlists/{id}/songs/...
	if len(parts) >= 2 && parts[1] == "songs" {
		handlePlaylistSongs(w, r, parts)
		return
	}

	id := parts[0] // will be "" if path is just /playlists/

	switch r.Method {

	case http.MethodGet:
		if id == "" {
			// GET /playlists/ → return all playlists
			respondJSON(w, http.StatusOK, store.GetAll())
		} else {
			// GET /playlists/{id} → return one playlist
			playlist, found := store.GetByID(id)
			if !found {
				respondError(w, http.StatusNotFound, "playlist not found")
				return // FAIL FAST
			}
			respondJSON(w, http.StatusOK, playlist)
		}

	case http.MethodPost:
		// POST /playlists/ → create a new playlist
		var input Playlist
		if err := decodeJSON(r, &input); err != nil {
			respondError(w, http.StatusBadRequest, "invalid JSON body")
			return // FAIL FAST
		}
		// Use the CONSTRUCTOR to create a proper playlist with auto-ID and timestamp
		newPlaylist := NewPlaylist(input.Name, input.Description)
		created := store.Create(newPlaylist)
		respondJSON(w, http.StatusCreated, created)

	case http.MethodPut:
		// PUT /playlists/{id} → update playlist name and description
		if id == "" {
			respondError(w, http.StatusBadRequest, "playlist ID is required")
			return
		}
		var input Playlist
		if err := decodeJSON(r, &input); err != nil {
			respondError(w, http.StatusBadRequest, "invalid JSON body")
			return
		}
		updated, found := store.Update(id, input)
		if !found {
			respondError(w, http.StatusNotFound, "playlist not found")
			return
		}
		respondJSON(w, http.StatusOK, updated)

	case http.MethodDelete:
		// DELETE /playlists/{id} → remove a playlist
		if id == "" {
			respondError(w, http.StatusBadRequest, "playlist ID is required")
			return
		}
		if !store.Delete(id) {
			respondError(w, http.StatusNotFound, "playlist not found")
			return
		}
		w.WriteHeader(http.StatusNoContent) // 204 = deleted, no body needed

	default:
		respondError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

// handlePlaylistSongs handles song operations within a playlist.
// This is the BONUS challenge — nested routes.
//
// Route: /playlists/{playlistID}/songs/{songID}
// parts[0] = playlistID, parts[1] = "songs", parts[2] = songID (optional)
func handlePlaylistSongs(w http.ResponseWriter, r *http.Request, parts []string) {
	playlistID := parts[0]

	// First, get the playlist — FAIL FAST if it doesn't exist
	playlist, found := store.GetByID(playlistID)
	if !found {
		respondError(w, http.StatusNotFound, "playlist not found")
		return
	}

	switch r.Method {

	case http.MethodPost:
		// POST /playlists/{id}/songs/ → add a song to the playlist
		var input Song
		if err := decodeJSON(r, &input); err != nil {
			respondError(w, http.StatusBadRequest, "invalid JSON body")
			return
		}
		// Use the CONSTRUCTOR to create a song with auto-generated ID
		newSong := NewSong(input.Title, input.Artist, input.Duration, input.Genre)

		// Add the song to the playlist's Songs slice (COMPOSITION in action)
		playlist.Songs = append(playlist.Songs, newSong)
		store.Update(playlistID, playlist)

		respondJSON(w, http.StatusCreated, newSong)

	case http.MethodDelete:
		// DELETE /playlists/{id}/songs/{songID} → remove a song
		if len(parts) < 3 || parts[2] == "" {
			respondError(w, http.StatusBadRequest, "song ID is required")
			return
		}
		songID := parts[2]

		// Find and remove the song from the slice
		found := false
		newSongs := []Song{}
		for _, song := range playlist.Songs {
			if song.ID == songID {
				found = true // skip this song (effectively deleting it)
			} else {
				newSongs = append(newSongs, song)
			}
		}
		if !found {
			respondError(w, http.StatusNotFound, "song not found in playlist")
			return
		}
		playlist.Songs = newSongs
		store.Update(playlistID, playlist)

		w.WriteHeader(http.StatusNoContent)

	default:
		respondError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
// MAIN — Wiring everything together
// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

func main() {
	// Initialize the store using the CONSTRUCTOR pattern.
	// store is typed as Storage (interface), but we assign
	// a *MemoryStore (concrete type). This works because
	// MemoryStore satisfies the Storage interface.
	// This is POLYMORPHISM — the variable's type is the interface,
	// but the actual value is a concrete implementation.
	store = NewMemoryStore()

	// Register routes with CORS enabled.
	// enableCORS wraps handlePlaylists with CORS headers.
	// This is the DECORATOR pattern — adding behavior without
	// modifying the original function.
	http.HandleFunc("/playlists/", enableCORS(handlePlaylists))

	fmt.Println("🎵 Music Playlist API running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
