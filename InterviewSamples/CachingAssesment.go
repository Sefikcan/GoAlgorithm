package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type User struct {
	Id       int
	Username string
}

type CachingServer struct {
	db    map[int]*User
	cache map[int]*User
	dbHit int
}

func NewServer() *CachingServer {
	db := make(map[int]*User)

	for i := 0; i < 100; i++ {
		db[i+1] = &User{
			Id:       i + 1,
			Username: fmt.Sprintf("user_%d", i+1),
		}
	}

	return &CachingServer{
		db:    db,
		cache: make(map[int]*User),
	}
}

func (s *CachingServer) tryCache(id int) (*User, bool) {
	user, ok := s.cache[id]
	return user, ok
}

func (s *CachingServer) handleGetUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	// first try to hit the cache
	user, ok := s.tryCache(id)
	if ok {
		json.NewEncoder(w).Encode(user)
		return
	}

	// hit the database
	user, ok = s.db[id]
	if !ok {
		panic("user not found")
	}
	s.dbHit++

	// insert in cache
	s.cache[id] = user

	json.NewEncoder(w).Encode(user)
}

func main() {

}
