package sessions

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/context"
)

const (
	DefaultSessionName = "gcalert-sessions-default"
	DefaultCookieName  = "acalertsession"
)

func NewSession(store *Store, cookieName string) *Session {
	return &Session{
		cookieName: cookieName,
		store:      store,
		Values:     map[string]interface{}{},
	}
}

type Session struct {
	cookieName string
	ID         string
	store      *Store
	request    *http.Request
	writer     http.ResponseWriter
	Values     map[string]interface{}
}

func StartSession(sessionName, cookieName string, store *Store) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var session *Session
		var err error
		session, err = store.Get(ctx.Request, cookieName)
		if err != nil {
			session, err = store.New(ctx.Request, cookieName)
			if err != nil {
				println("Abort: " + err.Error())
				ctx.Abort()
			}
		}
		session.writer = ctx.Writer
		ctx.Set(sessionName, session)
		defer context.Clear(ctx.Request)
		ctx.Next()
	}
}

func StartDefaultSession(store *Store) gin.HandlerFunc {
	return StartSession(DefaultSessionName, DefaultCookieName, store)
}

func GetSession(c *gin.Context, sessionName string) *Session {
	return c.MustGet(sessionName).(*Session)
}

func GetDefaultSession(c *gin.Context) *Session {
	return GetSession(c, DefaultSessionName)
}

// This returns the same result as s.session.Name()
func (s *Session) Name() string {
	return s.cookieName
}

func (s *Session) Get(key string) (interface{}, bool) {
	ret, exists := s.Values[key]
	return ret, exists
}

func (s *Session) Set(key string, val interface{}) {
	s.Values[key] = val
}

func (s *Session) Delete(key string) {
	delete(s.Values, key)
}

func (s *Session) Save() error {
	return s.store.Save(s.request, s.writer, s)
}

func (s *Session) Terminate() {
	s.store.Delete(s.ID)
}
