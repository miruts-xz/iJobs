package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/miruts/iJobs/entity"
)

type SessionGormRepositoryImpl struct {
	conn *gorm.DB
}

func NewSessionGormRepositoryImpl(conn *gorm.DB) *SessionGormRepositoryImpl {
	return &SessionGormRepositoryImpl{conn: conn}
}

func (sgr *SessionGormRepositoryImpl) Sessions() ([]entity.Session, error) {
	var sessions []entity.Session
	errs := sgr.conn.Find(&sessions).GetErrors()
	if len(errs) > 0 {
		return sessions, errs[0]
	}
	return sessions, nil
}
func (sgr *SessionGormRepositoryImpl) Session(id int) (entity.Session, error) {
	var session entity.Session
	errs := sgr.conn.First(&session, id).GetErrors()
	if len(errs) > 0 {
		return session, errs[0]
	}
	return session, nil
}
func (sgr *SessionGormRepositoryImpl) DeleteSession(id int) (entity.Session, error) {
	session, err := sgr.Session(id)
	if err != nil {
		return session, err
	}
	errs := sgr.conn.Delete(session, id).GetErrors()
	if len(errs) > 0 {
		return session, errs[0]
	}
	return session, nil
}
func (sgr *SessionGormRepositoryImpl) StoreSession(sess *entity.Session) (*entity.Session, error) {
	session := sess
	errs := sgr.conn.Create(&session).GetErrors()
	if len(errs) > 0 {
		return session, errs[0]
	}
	return session, nil
}
func (sgr *SessionGormRepositoryImpl) UpdateSession(sess *entity.Session) (*entity.Session, error) {
	session := sess
	errs := sgr.conn.Save(&session).GetErrors()
	if len(errs) > 0 {
		return session, errs[0]
	}
	return session, nil
}
func (sgr *SessionGormRepositoryImpl) SessionByValue(value string) (entity.Session, error) {
	var session entity.Session
	errs := sgr.conn.Where("uuid = ?", value).First(&session).GetErrors()
	if len(errs) > 0 {
		return session, errs[0]
	}
	return session, nil
}
