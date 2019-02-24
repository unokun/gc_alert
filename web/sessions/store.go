package sessions

import (
	"encoding/base64"

	"errors"
	"net/http"

	"github.com/google/uuid"
)

type Store struct {
	database map[string]interface{}
}

var kvs Store

func init() {
	kvs.database = map[string]interface{}{}
}

func NewStore() *Store {
	return &kvs
}

func (s *Store) NewSessionID() string {
	return longSecureRandomBase64()
}

func (s *Store) Exists(sessionID string) bool {
	_, r := s.database[sessionID]
	return r
}

func (s *Store) Flush() {
	s.database = map[string]interface{}{}
}

func (s *Store) Get(r *http.Request, cookieName string) (*Session, error) {
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		// No cookies in the request.
		return nil, err
	}

	sessionID := cookie.Value
	// restore session
	buffer, exists := s.database[sessionID]
	if !exists {
		return nil, errors.New("Invalid sessionID")
	}

	session := buffer.(*Session)
	session.request = r
	return session, nil
}

func (s *Store) New(r *http.Request, cookieName string) (*Session, error) {
	cookie, err := r.Cookie(cookieName)
	if err == nil && s.Exists(cookie.Value) {
		return nil, errors.New("sessionID already exists")
	}

	session := NewSession(s, cookieName)
	session.ID = s.NewSessionID()
	session.request = r

	return session, nil
}

func (s *Store) Save(r *http.Request, w http.ResponseWriter, session *Session) error {
	s.database[session.ID] = session

	c := &http.Cookie{
		Name:  session.Name(),
		Value: session.ID,
		Path:  "/",
	}

	http.SetCookie(session.writer, c)
	return nil
}

func (s *Store) Delete(sessionID string) {
	delete(s.database, sessionID)
}

func longSecureRandomBase64() string {
	return secureRandomBase64() + secureRandomBase64()
}
func secureRandomBase64() string {
	return base64.StdEncoding.EncodeToString(uuid.New().NodeID())
}
