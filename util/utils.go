package util

import (
	"errors"
	"fmt"
	"github.com/miruts/iJobs/entity"
	"github.com/miruts/iJobs/usecases/company"
	"github.com/miruts/iJobs/usecases/jobseeker"
	"github.com/miruts/iJobs/usecases/session"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

// SaveFile saves multipart file on the given path
func SaveFile(file multipart.File, path string) bool {
	var data []byte
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return false
	}
	err = ioutil.WriteFile(path, data, 0644)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return false
	}
	return true
}

func Authenticate(sessSrv session.SessionService, r *http.Request) (sess entity.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = entity.Session{Uuid: cookie.Value}
		session, err := sessSrv.Check(&sess)
		if err != nil {
			fmt.Println(err)
			return session, err
		}
		return session, nil
	}
	return sess, err
}
func CreateSession(w *http.ResponseWriter, sess *entity.Session) error {
	if sess != nil {
		cookie := http.Cookie{}
		cookie.Name = "_cookie"
		cookie.Value = sess.Uuid
		cookie.HttpOnly = true
		http.SetCookie(*w, &cookie)
		return nil
	}
	return errors.New("invalid session")
}

func ClearExpiredSessions(sess session.SessionService) {
	sessions, err := sess.Sessions()
	if err != nil {
		return
	}
	for _, s := range sessions {
		_, _ = sess.Check(&s)
	}
}
func DetectUser(w *http.ResponseWriter, r *http.Request, sess entity.Session, jsSrv jobseeker.JobseekerService, cmpSrv company.CompanyService) {
	_, err1 := jsSrv.JobSeeker(int(sess.UserID))
	_, err2 := cmpSrv.Company(int(sess.UserID))
	if err2 != nil && err1 == nil {
		// Its jobseeker
		http.Redirect(*w, r, "/jobseeker/home", http.StatusSeeOther)
	} else if err1 != nil && err2 == nil {
		// Its company
		http.Redirect(*w, r, "/company/home", http.StatusSeeOther)
	} else {
		fmt.Printf("sess error")
	}
}
