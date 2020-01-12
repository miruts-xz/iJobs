package service

import (
	"errors"
	"fmt"
	"github.com/miruts/iJobs/entity"
	"github.com/miruts/iJobs/usecases/session"
	"time"
)

// SessionServiceImpl implements SessionService
type SessionServiceImpl struct {
	sessRepo session.SessionRepository
}

// NewSessionServiceImpl creates new SessionServiceImpl
func NewSessionServiceImpl(ss session.SessionRepository) *SessionServiceImpl {
	return &SessionServiceImpl{sessRepo: ss}
}

// Sessions retrieves all sessions
func (ss *SessionServiceImpl) Sessions() ([]entity.Session, error) {
	return ss.sessRepo.Sessions()
}

// Session retrieves a session with given id
func (ss *SessionServiceImpl) Session(id int) (entity.Session, error) {
	return ss.sessRepo.Session(id)
}

// DeleteSession deletes a session with given id
func (ss *SessionServiceImpl) DeleteSession(id int) (entity.Session, error) {
	return ss.sessRepo.DeleteSession(id)
}

// StoreSession stores a given session
func (ss *SessionServiceImpl) StoreSession(sess *entity.Session) (*entity.Session, error) {
	return ss.sessRepo.StoreSession(sess)
}

// SessionByValue retrieves a session with given session value
func (ss *SessionServiceImpl) SessionByValue(value string) (entity.Session, error) {
	return ss.sessRepo.SessionByValue(value)
}

// UpdateSession updates a given session
func (ss *SessionServiceImpl) UpdateSession(sess *entity.Session) (*entity.Session, error) {
	return ss.sessRepo.UpdateSession(sess)
}

// Check checks for availability and expiration of a session
func (ss *SessionServiceImpl) Check(sess *entity.Session) (bool, entity.Session, error) {
	session, err := ss.SessionByValue(sess.Uuid)
	if err != nil {
		return false, session, errors.New("session not found")
	} else if time.Now().Sub(session.CreatedAt) < 6*time.Hour {
		session.CreatedAt = time.Now()
		session, err := ss.UpdateSession(&session)
		if err != nil {
			fmt.Println("Storing session error")
			return false, *session, err
		}
		return true, *session, nil
	} else {
		_, err := ss.DeleteSession(int(session.ID))
		if err != nil {
			return false, session, errors.New("invalid session")
		}
		return false, session, nil
	}
}
