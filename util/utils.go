package util

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
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

/*// Authenticate authenticates a given request for validity (of user)
func Authenticate(sessSrv session.SessionService, r *http.Request) (ok bool, sess entity.Session) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = entity.Session{Uuid: cookie.Value}
		ok, session, err := sessSrv.Check(&sess)
		if err != nil || !ok {
			fmt.Println(err)
			return ok, session
		}
		return ok, session
	}
	return false, sess
}

// CreateSession creates new Session
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

// DetectUser detects the type of logged in user
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
		fmt.Printf("session error")
	}
}

 DestroySession destroy a session
func DestroySession(w *http.ResponseWriter, r *http.Request) error {
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		return err
	}
	cookie.MaxAge = -1
	http.SetCookie(*w, cookie)
	return nil
}
*/
