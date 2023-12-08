// Code generated by ent, DO NOT EDIT.

package lockorderfulfillment

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/paycrest/protocol/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(sql.FieldEQ(FieldUpdatedAt, v))
}

// TxID applies equality check predicate on the "tx_id" field. It's identical to TxIDEQ.
func TxID(v string) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(sql.FieldEQ(FieldTxID, v))
}

// ValidationError applies equality check predicate on the "validation_error" field. It's identical to ValidationErrorEQ.
func ValidationError(v string) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(sql.FieldEQ(FieldValidationError, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(sql.FieldLTE(FieldUpdatedAt, v))
}

// TxIDEQ applies the EQ predicate on the "tx_id" field.
func TxIDEQ(v string) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(sql.FieldEQ(FieldTxID, v))
}

// TxIDNEQ applies the NEQ predicate on the "tx_id" field.
func TxIDNEQ(v string) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(sql.FieldNEQ(FieldTxID, v))
}

// TxIDIn applies the In predicate on the "tx_id" field.
func TxIDIn(vs ...string) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(sql.FieldIn(FieldTxID, vs...))
}

// TxIDNotIn applies the NotIn predicate on the "tx_id" field.
func TxIDNotIn(vs ...string) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(sql.FieldNotIn(FieldTxID, vs...))
}

// TxIDGT applies the GT predicate on the "tx_id" field.
func TxIDGT(v string) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(sql.FieldGT(FieldTxID, v))
}

// TxIDGTE applies the GTE predicate on the "tx_id" field.
func TxIDGTE(v string) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(sql.FieldGTE(FieldTxID, v))
}

// TxIDLT applies the LT predicate on the "tx_id" field.
func TxIDLT(v string) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(sql.FieldLT(FieldTxID, v))
}

// TxIDLTE applies the LTE predicate on the "tx_id" field.
func TxIDLTE(v string) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(sql.FieldLTE(FieldTxID, v))
}

// TxIDContains applies the Contains predicate on the "tx_id" field.
func TxIDContains(v string) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(sql.FieldContains(FieldTxID, v))
}

// TxIDHasPrefix applies the HasPrefix predicate on the "tx_id" field.
func TxIDHasPrefix(v string) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(sql.FieldHasPrefix(FieldTxID, v))
}

// TxIDHasSuffix applies the HasSuffix predicate on the "tx_id" field.
func TxIDHasSuffix(v string) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(sql.FieldHasSuffix(FieldTxID, v))
}

// TxIDEqualFold applies the EqualFold predicate on the "tx_id" field.
func TxIDEqualFold(v string) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(sql.FieldEqualFold(FieldTxID, v))
}

// TxIDContainsFold applies the ContainsFold predicate on the "tx_id" field.
func TxIDContainsFold(v string) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(sql.FieldContainsFold(FieldTxID, v))
}

// ValidationStatusEQ applies the EQ predicate on the "validation_status" field.
func ValidationStatusEQ(v ValidationStatus) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(sql.FieldEQ(FieldValidationStatus, v))
}

// ValidationStatusNEQ applies the NEQ predicate on the "validation_status" field.
func ValidationStatusNEQ(v ValidationStatus) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(sql.FieldNEQ(FieldValidationStatus, v))
}

// ValidationStatusIn applies the In predicate on the "validation_status" field.
func ValidationStatusIn(vs ...ValidationStatus) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(sql.FieldIn(FieldValidationStatus, vs...))
}

// ValidationStatusNotIn applies the NotIn predicate on the "validation_status" field.
func ValidationStatusNotIn(vs ...ValidationStatus) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(sql.FieldNotIn(FieldValidationStatus, vs...))
}

// ValidationErrorEQ applies the EQ predicate on the "validation_error" field.
func ValidationErrorEQ(v string) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(sql.FieldEQ(FieldValidationError, v))
}

// ValidationErrorNEQ applies the NEQ predicate on the "validation_error" field.
func ValidationErrorNEQ(v string) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(sql.FieldNEQ(FieldValidationError, v))
}

// ValidationErrorIn applies the In predicate on the "validation_error" field.
func ValidationErrorIn(vs ...string) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(sql.FieldIn(FieldValidationError, vs...))
}

// ValidationErrorNotIn applies the NotIn predicate on the "validation_error" field.
func ValidationErrorNotIn(vs ...string) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(sql.FieldNotIn(FieldValidationError, vs...))
}

// ValidationErrorGT applies the GT predicate on the "validation_error" field.
func ValidationErrorGT(v string) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(sql.FieldGT(FieldValidationError, v))
}

// ValidationErrorGTE applies the GTE predicate on the "validation_error" field.
func ValidationErrorGTE(v string) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(sql.FieldGTE(FieldValidationError, v))
}

// ValidationErrorLT applies the LT predicate on the "validation_error" field.
func ValidationErrorLT(v string) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(sql.FieldLT(FieldValidationError, v))
}

// ValidationErrorLTE applies the LTE predicate on the "validation_error" field.
func ValidationErrorLTE(v string) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(sql.FieldLTE(FieldValidationError, v))
}

// ValidationErrorContains applies the Contains predicate on the "validation_error" field.
func ValidationErrorContains(v string) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(sql.FieldContains(FieldValidationError, v))
}

// ValidationErrorHasPrefix applies the HasPrefix predicate on the "validation_error" field.
func ValidationErrorHasPrefix(v string) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(sql.FieldHasPrefix(FieldValidationError, v))
}

// ValidationErrorHasSuffix applies the HasSuffix predicate on the "validation_error" field.
func ValidationErrorHasSuffix(v string) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(sql.FieldHasSuffix(FieldValidationError, v))
}

// ValidationErrorIsNil applies the IsNil predicate on the "validation_error" field.
func ValidationErrorIsNil() predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(sql.FieldIsNull(FieldValidationError))
}

// ValidationErrorNotNil applies the NotNil predicate on the "validation_error" field.
func ValidationErrorNotNil() predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(sql.FieldNotNull(FieldValidationError))
}

// ValidationErrorEqualFold applies the EqualFold predicate on the "validation_error" field.
func ValidationErrorEqualFold(v string) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(sql.FieldEqualFold(FieldValidationError, v))
}

// ValidationErrorContainsFold applies the ContainsFold predicate on the "validation_error" field.
func ValidationErrorContainsFold(v string) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(sql.FieldContainsFold(FieldValidationError, v))
}

// HasOrder applies the HasEdge predicate on the "order" edge.
func HasOrder() predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, OrderTable, OrderColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrderWith applies the HasEdge predicate on the "order" edge with a given conditions (other predicates).
func HasOrderWith(preds ...predicate.LockPaymentOrder) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(func(s *sql.Selector) {
		step := newOrderStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.LockOrderFulfillment) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.LockOrderFulfillment) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.LockOrderFulfillment) predicate.LockOrderFulfillment {
	return predicate.LockOrderFulfillment(func(s *sql.Selector) {
		p(s.Not())
	})
}
