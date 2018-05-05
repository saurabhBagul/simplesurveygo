package main

import (
	"fmt"
	"net/http"
	"simplesurveygo/servicehandlers"
)

func test(messages chan string) {

	http.Handle("/", http.FileServer(http.Dir("./static")))

	messages <- "ping"
	pingHandler := servicehandlers.PingHandler{}
	authHandler := servicehandlers.UserValidationHandler{}
	sessionHandler := servicehandlers.SessionHandler{}
	surveyHandler := servicehandlers.SurveyHandler{}
	userSurveyHandler := servicehandlers.UserSurveyHandler{}
	signupHandler := servicehandlers.SignupHandler{}
	// Serves the API content
	http.Handle("/api/v1/ping/", pingHandler)

	http.Handle("/api/v1/signup/", signupHandler)
	http.Handle("/api/v1/authenticate/", authHandler)
	http.Handle("/api/v1/validate/", sessionHandler)

	http.Handle("/api/v1/survey/{surveyname}", surveyHandler)
	http.Handle("/api/v1/survey/", surveyHandler)

	http.Handle("/api/v1/usersurvey/", userSurveyHandler)

	// Start Server
	http.ListenAndServe(":3000", nil)

}

func main() {
	messages := make(chan string, 10)
	go test(messages)
	fmt.Println(<-messages)
	// Serves the html pages
	// http.Handle("/", http.FileServer(http.Dir("./static")))

	// pingHandler := servicehandlers.PingHandler{}
	// authHandler := servicehandlers.UserValidationHandler{}
	// sessionHandler := servicehandlers.SessionHandler{}
	// surveyHandler := servicehandlers.SurveyHandler{}
	// userSurveyHandler := servicehandlers.UserSurveyHandler{}
	// signupHandler := servicehandlers.SignupHandler{}

	// // Serves the API content
	// http.Handle("/api/v1/ping/", pingHandler)

	// http.Handle("/api/v1/signup/", signupHandler)
	// http.Handle("/api/v1/authenticate/", authHandler)
	// http.Handle("/api/v1/validate/", sessionHandler)

	// http.Handle("/api/v1/survey/{surveyname}", surveyHandler)
	// http.Handle("/api/v1/survey/", surveyHandler)

	// http.Handle("/api/v1/usersurvey/", userSurveyHandler)

	// // Start Server
	// http.ListenAndServe(":3000", nil)
}
