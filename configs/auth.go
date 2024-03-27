package configs

import "github.com/gorilla/sessions"

var SessionStore = sessions.NewCookieStore([]byte(GetSessionKey()))
