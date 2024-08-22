// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/paycrest/protocol/ent/providerprofile"
	"github.com/paycrest/protocol/ent/senderprofile"
	"github.com/paycrest/protocol/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// FirstName holds the value of the "first_name" field.
	FirstName string `json:"first_name,omitempty"`
	// LastName holds the value of the "last_name" field.
	LastName string `json:"last_name,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"-"`
	// Scope holds the value of the "scope" field.
	Scope string `json:"scope,omitempty"`
	// IsEmailVerified holds the value of the "is_email_verified" field.
	IsEmailVerified bool `json:"is_email_verified,omitempty"`
	// HasEarlyAccess holds the value of the "has_early_access" field.
	HasEarlyAccess bool `json:"has_early_access,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges        UserEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// SenderProfile holds the value of the sender_profile edge.
	SenderProfile *SenderProfile `json:"sender_profile,omitempty"`
	// ProviderProfile holds the value of the provider_profile edge.
	ProviderProfile *ProviderProfile `json:"provider_profile,omitempty"`
	// VerificationToken holds the value of the verification_token edge.
	VerificationToken []*VerificationToken `json:"verification_token,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// SenderProfileOrErr returns the SenderProfile value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) SenderProfileOrErr() (*SenderProfile, error) {
	if e.SenderProfile != nil {
		return e.SenderProfile, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: senderprofile.Label}
	}
	return nil, &NotLoadedError{edge: "sender_profile"}
}

// ProviderProfileOrErr returns the ProviderProfile value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) ProviderProfileOrErr() (*ProviderProfile, error) {
	if e.ProviderProfile != nil {
		return e.ProviderProfile, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: providerprofile.Label}
	}
	return nil, &NotLoadedError{edge: "provider_profile"}
}

// VerificationTokenOrErr returns the VerificationToken value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) VerificationTokenOrErr() ([]*VerificationToken, error) {
	if e.loadedTypes[2] {
		return e.VerificationToken, nil
	}
	return nil, &NotLoadedError{edge: "verification_token"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldIsEmailVerified, user.FieldHasEarlyAccess:
			values[i] = new(sql.NullBool)
		case user.FieldFirstName, user.FieldLastName, user.FieldEmail, user.FieldPassword, user.FieldScope:
			values[i] = new(sql.NullString)
		case user.FieldCreatedAt, user.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case user.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				u.ID = *value
			}
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		case user.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				u.UpdatedAt = value.Time
			}
		case user.FieldFirstName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field first_name", values[i])
			} else if value.Valid {
				u.FirstName = value.String
			}
		case user.FieldLastName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field last_name", values[i])
			} else if value.Valid {
				u.LastName = value.String
			}
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = value.String
			}
		case user.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				u.Password = value.String
			}
		case user.FieldScope:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field scope", values[i])
			} else if value.Valid {
				u.Scope = value.String
			}
		case user.FieldIsEmailVerified:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_email_verified", values[i])
			} else if value.Valid {
				u.IsEmailVerified = value.Bool
			}
		case user.FieldHasEarlyAccess:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field has_early_access", values[i])
			} else if value.Valid {
				u.HasEarlyAccess = value.Bool
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QuerySenderProfile queries the "sender_profile" edge of the User entity.
func (u *User) QuerySenderProfile() *SenderProfileQuery {
	return NewUserClient(u.config).QuerySenderProfile(u)
}

// QueryProviderProfile queries the "provider_profile" edge of the User entity.
func (u *User) QueryProviderProfile() *ProviderProfileQuery {
	return NewUserClient(u.config).QueryProviderProfile(u)
}

// QueryVerificationToken queries the "verification_token" edge of the User entity.
func (u *User) QueryVerificationToken() *VerificationTokenQuery {
	return NewUserClient(u.config).QueryVerificationToken(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("first_name=")
	builder.WriteString(u.FirstName)
	builder.WriteString(", ")
	builder.WriteString("last_name=")
	builder.WriteString(u.LastName)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(u.Email)
	builder.WriteString(", ")
	builder.WriteString("password=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("scope=")
	builder.WriteString(u.Scope)
	builder.WriteString(", ")
	builder.WriteString("is_email_verified=")
	builder.WriteString(fmt.Sprintf("%v", u.IsEmailVerified))
	builder.WriteString(", ")
	builder.WriteString("has_early_access=")
	builder.WriteString(fmt.Sprintf("%v", u.HasEarlyAccess))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User
