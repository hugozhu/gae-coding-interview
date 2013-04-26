package webapp

import (
	"appengine"
	"appengine/channel"
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/interview/create", interview_create)
	http.HandleFunc("/interview/send", interview_send)

}

func handler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://hugozhu.myalert.info", 302)
}

func interview_create(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	tok, err := channel.Create(c, r.FormValue("id"))
	callback := r.FormValue("callback")
	if err != nil {
		http.Error(w, "Couldn't create Channel", http.StatusInternalServerError)
		c.Errorf("channel.Create: %v", err)
		return
	}
	if callback == "" {
		callback = "callback"
	}
	fmt.Fprintf(w, callback+"('%s')", tok)
}

func interview_send(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	token := r.FormValue("id")
	code := r.FormValue("code")
	channel.Send(c, token, code)
}
