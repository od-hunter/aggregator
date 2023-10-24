// Code generated by ent, DO NOT EDIT.

package apikey

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the apikey type in the database.
	Label = "api_key"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSecret holds the string denoting the secret field in the database.
	FieldSecret = "secret"
	// EdgeSenderProfile holds the string denoting the sender_profile edge name in mutations.
	EdgeSenderProfile = "sender_profile"
	// EdgeProviderProfile holds the string denoting the provider_profile edge name in mutations.
	EdgeProviderProfile = "provider_profile"
	// EdgeValidatorProfile holds the string denoting the validator_profile edge name in mutations.
	EdgeValidatorProfile = "validator_profile"
	// EdgePaymentOrders holds the string denoting the payment_orders edge name in mutations.
	EdgePaymentOrders = "payment_orders"
	// Table holds the table name of the apikey in the database.
	Table = "api_keys"
	// SenderProfileTable is the table that holds the sender_profile relation/edge.
	SenderProfileTable = "api_keys"
	// SenderProfileInverseTable is the table name for the SenderProfile entity.
	// It exists in this package in order to avoid circular dependency with the "senderprofile" package.
	SenderProfileInverseTable = "sender_profiles"
	// SenderProfileColumn is the table column denoting the sender_profile relation/edge.
	SenderProfileColumn = "sender_profile_api_key"
	// ProviderProfileTable is the table that holds the provider_profile relation/edge.
	ProviderProfileTable = "api_keys"
	// ProviderProfileInverseTable is the table name for the ProviderProfile entity.
	// It exists in this package in order to avoid circular dependency with the "providerprofile" package.
	ProviderProfileInverseTable = "provider_profiles"
	// ProviderProfileColumn is the table column denoting the provider_profile relation/edge.
	ProviderProfileColumn = "provider_profile_api_key"
	// ValidatorProfileTable is the table that holds the validator_profile relation/edge.
	ValidatorProfileTable = "api_keys"
	// ValidatorProfileInverseTable is the table name for the ValidatorProfile entity.
	// It exists in this package in order to avoid circular dependency with the "validatorprofile" package.
	ValidatorProfileInverseTable = "validator_profiles"
	// ValidatorProfileColumn is the table column denoting the validator_profile relation/edge.
	ValidatorProfileColumn = "validator_profile_api_key"
	// PaymentOrdersTable is the table that holds the payment_orders relation/edge.
	PaymentOrdersTable = "payment_orders"
	// PaymentOrdersInverseTable is the table name for the PaymentOrder entity.
	// It exists in this package in order to avoid circular dependency with the "paymentorder" package.
	PaymentOrdersInverseTable = "payment_orders"
	// PaymentOrdersColumn is the table column denoting the payment_orders relation/edge.
	PaymentOrdersColumn = "api_key_payment_orders"
)

// Columns holds all SQL columns for apikey fields.
var Columns = []string{
	FieldID,
	FieldSecret,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "api_keys"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"provider_profile_api_key",
	"sender_profile_api_key",
	"validator_profile_api_key",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// SecretValidator is a validator for the "secret" field. It is called by the builders before save.
	SecretValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the APIKey queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// BySecret orders the results by the secret field.
func BySecret(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSecret, opts...).ToFunc()
}

// BySenderProfileField orders the results by sender_profile field.
func BySenderProfileField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSenderProfileStep(), sql.OrderByField(field, opts...))
	}
}

// ByProviderProfileField orders the results by provider_profile field.
func ByProviderProfileField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProviderProfileStep(), sql.OrderByField(field, opts...))
	}
}

// ByValidatorProfileField orders the results by validator_profile field.
func ByValidatorProfileField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newValidatorProfileStep(), sql.OrderByField(field, opts...))
	}
}

// ByPaymentOrdersCount orders the results by payment_orders count.
func ByPaymentOrdersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPaymentOrdersStep(), opts...)
	}
}

// ByPaymentOrders orders the results by payment_orders terms.
func ByPaymentOrders(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPaymentOrdersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newSenderProfileStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SenderProfileInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, SenderProfileTable, SenderProfileColumn),
	)
}
func newProviderProfileStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProviderProfileInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, ProviderProfileTable, ProviderProfileColumn),
	)
}
func newValidatorProfileStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ValidatorProfileInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, ValidatorProfileTable, ValidatorProfileColumn),
	)
}
func newPaymentOrdersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PaymentOrdersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PaymentOrdersTable, PaymentOrdersColumn),
	)
}
