// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/paycrest/paycrest-protocol/ent/apikey"
	"github.com/paycrest/paycrest-protocol/ent/user"
	"github.com/paycrest/paycrest-protocol/ent/validatorprofile"
)

// ValidatorProfile is the model entity for the ValidatorProfile schema.
type ValidatorProfile struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// WalletAddress holds the value of the "wallet_address" field.
	WalletAddress string `json:"wallet_address,omitempty"`
	// HostIdentifier holds the value of the "host_identifier" field.
	HostIdentifier string `json:"host_identifier,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ValidatorProfileQuery when eager-loading is set.
	Edges                  ValidatorProfileEdges `json:"edges"`
	user_validator_profile *uuid.UUID
	selectValues           sql.SelectValues
}

// ValidatorProfileEdges holds the relations/edges for other nodes in the graph.
type ValidatorProfileEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// ValidatedFulfillments holds the value of the validated_fulfillments edge.
	ValidatedFulfillments []*LockOrderFulfillment `json:"validated_fulfillments,omitempty"`
	// APIKey holds the value of the api_key edge.
	APIKey *APIKey `json:"api_key,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ValidatorProfileEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// ValidatedFulfillmentsOrErr returns the ValidatedFulfillments value or an error if the edge
// was not loaded in eager-loading.
func (e ValidatorProfileEdges) ValidatedFulfillmentsOrErr() ([]*LockOrderFulfillment, error) {
	if e.loadedTypes[1] {
		return e.ValidatedFulfillments, nil
	}
	return nil, &NotLoadedError{edge: "validated_fulfillments"}
}

// APIKeyOrErr returns the APIKey value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ValidatorProfileEdges) APIKeyOrErr() (*APIKey, error) {
	if e.loadedTypes[2] {
		if e.APIKey == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: apikey.Label}
		}
		return e.APIKey, nil
	}
	return nil, &NotLoadedError{edge: "api_key"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ValidatorProfile) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case validatorprofile.FieldWalletAddress, validatorprofile.FieldHostIdentifier:
			values[i] = new(sql.NullString)
		case validatorprofile.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case validatorprofile.FieldID:
			values[i] = new(uuid.UUID)
		case validatorprofile.ForeignKeys[0]: // user_validator_profile
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ValidatorProfile fields.
func (vp *ValidatorProfile) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case validatorprofile.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				vp.ID = *value
			}
		case validatorprofile.FieldWalletAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field wallet_address", values[i])
			} else if value.Valid {
				vp.WalletAddress = value.String
			}
		case validatorprofile.FieldHostIdentifier:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field host_identifier", values[i])
			} else if value.Valid {
				vp.HostIdentifier = value.String
			}
		case validatorprofile.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				vp.UpdatedAt = value.Time
			}
		case validatorprofile.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field user_validator_profile", values[i])
			} else if value.Valid {
				vp.user_validator_profile = new(uuid.UUID)
				*vp.user_validator_profile = *value.S.(*uuid.UUID)
			}
		default:
			vp.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ValidatorProfile.
// This includes values selected through modifiers, order, etc.
func (vp *ValidatorProfile) Value(name string) (ent.Value, error) {
	return vp.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the ValidatorProfile entity.
func (vp *ValidatorProfile) QueryUser() *UserQuery {
	return NewValidatorProfileClient(vp.config).QueryUser(vp)
}

// QueryValidatedFulfillments queries the "validated_fulfillments" edge of the ValidatorProfile entity.
func (vp *ValidatorProfile) QueryValidatedFulfillments() *LockOrderFulfillmentQuery {
	return NewValidatorProfileClient(vp.config).QueryValidatedFulfillments(vp)
}

// QueryAPIKey queries the "api_key" edge of the ValidatorProfile entity.
func (vp *ValidatorProfile) QueryAPIKey() *APIKeyQuery {
	return NewValidatorProfileClient(vp.config).QueryAPIKey(vp)
}

// Update returns a builder for updating this ValidatorProfile.
// Note that you need to call ValidatorProfile.Unwrap() before calling this method if this ValidatorProfile
// was returned from a transaction, and the transaction was committed or rolled back.
func (vp *ValidatorProfile) Update() *ValidatorProfileUpdateOne {
	return NewValidatorProfileClient(vp.config).UpdateOne(vp)
}

// Unwrap unwraps the ValidatorProfile entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (vp *ValidatorProfile) Unwrap() *ValidatorProfile {
	_tx, ok := vp.config.driver.(*txDriver)
	if !ok {
		panic("ent: ValidatorProfile is not a transactional entity")
	}
	vp.config.driver = _tx.drv
	return vp
}

// String implements the fmt.Stringer.
func (vp *ValidatorProfile) String() string {
	var builder strings.Builder
	builder.WriteString("ValidatorProfile(")
	builder.WriteString(fmt.Sprintf("id=%v, ", vp.ID))
	builder.WriteString("wallet_address=")
	builder.WriteString(vp.WalletAddress)
	builder.WriteString(", ")
	builder.WriteString("host_identifier=")
	builder.WriteString(vp.HostIdentifier)
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(vp.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// ValidatorProfiles is a parsable slice of ValidatorProfile.
type ValidatorProfiles []*ValidatorProfile
