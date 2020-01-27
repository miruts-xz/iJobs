package session

import "github.com/miruts/iJobs/entity"

// SessionServices defines Session related services
type SessionService interface {
	Sessions() ([]entity.Session, error)
	Session(id int) (entity.Session, error)
	DeleteSession(id int) (entity.Session, error)
	UpdateSession(sess *entity.Session) (*entity.Session, error)
	StoreSession(sess *entity.Session) (*entity.Session, error)
	SessionByValue(value string) (entity.Session, error)
	Check(sess *entity.Session) (bool, entity.Session, error)
}
