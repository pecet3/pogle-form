package login

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/pecet3/logger"
)

const (
	EXPIRY_DURATION  = time.Second * 60 * 60 * 2
	BLOCK_DURATION   = time.Second * 60 * 60
	CLEANUP_DURATION = time.Second * 60 * 60
)

type LoginSession struct {
	Cookie string
	Expiry time.Time
}

type sessions = map[string]*LoginSession

func newSession() *LoginSession {
	expiresAt := time.Now().Add(EXPIRY_DURATION)
	cookie := uuid.NewString()
	ea := &LoginSession{
		Expiry: expiresAt,
		Cookie: cookie,
	}
	return ea
}

func (l *Login) Add() *LoginSession {
	session := newSession()
	l.sMu.Lock()
	l.sessions[session.Cookie] = session
	l.sMu.Unlock()
	return session
}

func (l *Login) Get(cookie string) (*LoginSession, bool) {
	l.sMu.RLock()
	session, ok := l.sessions[cookie]
	l.sMu.RUnlock()
	return session, ok
}

func (l *Login) Delete(cookie string) {
	l.sMu.Lock()
	delete(l.sessions, cookie)

}

func (ss *Login) cleanUpExpiredSessions() {
	for {
		time.Sleep(CLEANUP_DURATION)
		cleanedSessions := 0
		ss.sMu.Lock()
		for token, session := range ss.sessions {
			if time.Now().After(session.Expiry) {
				delete(ss.sessions, token)
				cleanedSessions++
			}
		}
		ss.sMu.Unlock()
		logger.Info(fmt.Sprintf(`Cleaned Expired Login Sessions: %d`, cleanedSessions))
	}
}
