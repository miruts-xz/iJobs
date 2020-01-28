package repository

import (
	"github.com/miruts/iJobs/entity"
)

type SessionMockRepository struct {
}

func NewSessionMockRepository() *SessionMockRepository {
	return &SessionMockRepository{}
}

func (s SessionMockRepository) Session(sessionID string) (*entity.Session, []error) {
	panic("implement me")
}

func (s SessionMockRepository) StoreSession(session *entity.Session) (*entity.Session, []error) {
	panic("implement me")
}

func (s SessionMockRepository) DeleteSession(sessionID string) (*entity.Session, []error) {
	panic("implement me")
}
