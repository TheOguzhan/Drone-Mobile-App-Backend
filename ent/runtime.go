// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/TheOguzhan/Drone-Mobile-App-Backend/ent/schema"
	"github.com/TheOguzhan/Drone-Mobile-App-Backend/ent/users"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	usersFields := schema.Users{}.Fields()
	_ = usersFields
	// usersDescEmail is the schema descriptor for email field.
	usersDescEmail := usersFields[2].Descriptor()
	// users.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	users.EmailValidator = usersDescEmail.Validators[0].(func(string) error)
	// usersDescIsUserConfirmed is the schema descriptor for is_user_confirmed field.
	usersDescIsUserConfirmed := usersFields[5].Descriptor()
	// users.DefaultIsUserConfirmed holds the default value on creation for the is_user_confirmed field.
	users.DefaultIsUserConfirmed = usersDescIsUserConfirmed.Default.(bool)
}