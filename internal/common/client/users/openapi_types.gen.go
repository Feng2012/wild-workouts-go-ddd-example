// Package users provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package users

// User defines model for User.
type User struct {
	Balance     int    `json:"balance"`
	DisplayName string `json:"displayName"`
	Role        string `json:"role"`
}
