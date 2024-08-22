// Code generated by ent, DO NOT EDIT.

package providerordertoken

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/paycrest/protocol/ent/predicate"
	"github.com/shopspring/decimal"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldEQ(FieldUpdatedAt, v))
}

// Symbol applies equality check predicate on the "symbol" field. It's identical to SymbolEQ.
func Symbol(v string) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldEQ(FieldSymbol, v))
}

// FixedConversionRate applies equality check predicate on the "fixed_conversion_rate" field. It's identical to FixedConversionRateEQ.
func FixedConversionRate(v decimal.Decimal) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldEQ(FieldFixedConversionRate, v))
}

// FloatingConversionRate applies equality check predicate on the "floating_conversion_rate" field. It's identical to FloatingConversionRateEQ.
func FloatingConversionRate(v decimal.Decimal) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldEQ(FieldFloatingConversionRate, v))
}

// MaxOrderAmount applies equality check predicate on the "max_order_amount" field. It's identical to MaxOrderAmountEQ.
func MaxOrderAmount(v decimal.Decimal) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldEQ(FieldMaxOrderAmount, v))
}

// MinOrderAmount applies equality check predicate on the "min_order_amount" field. It's identical to MinOrderAmountEQ.
func MinOrderAmount(v decimal.Decimal) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldEQ(FieldMinOrderAmount, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldLTE(FieldUpdatedAt, v))
}

// SymbolEQ applies the EQ predicate on the "symbol" field.
func SymbolEQ(v string) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldEQ(FieldSymbol, v))
}

// SymbolNEQ applies the NEQ predicate on the "symbol" field.
func SymbolNEQ(v string) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldNEQ(FieldSymbol, v))
}

// SymbolIn applies the In predicate on the "symbol" field.
func SymbolIn(vs ...string) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldIn(FieldSymbol, vs...))
}

// SymbolNotIn applies the NotIn predicate on the "symbol" field.
func SymbolNotIn(vs ...string) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldNotIn(FieldSymbol, vs...))
}

// SymbolGT applies the GT predicate on the "symbol" field.
func SymbolGT(v string) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldGT(FieldSymbol, v))
}

// SymbolGTE applies the GTE predicate on the "symbol" field.
func SymbolGTE(v string) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldGTE(FieldSymbol, v))
}

// SymbolLT applies the LT predicate on the "symbol" field.
func SymbolLT(v string) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldLT(FieldSymbol, v))
}

// SymbolLTE applies the LTE predicate on the "symbol" field.
func SymbolLTE(v string) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldLTE(FieldSymbol, v))
}

// SymbolContains applies the Contains predicate on the "symbol" field.
func SymbolContains(v string) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldContains(FieldSymbol, v))
}

// SymbolHasPrefix applies the HasPrefix predicate on the "symbol" field.
func SymbolHasPrefix(v string) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldHasPrefix(FieldSymbol, v))
}

// SymbolHasSuffix applies the HasSuffix predicate on the "symbol" field.
func SymbolHasSuffix(v string) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldHasSuffix(FieldSymbol, v))
}

// SymbolEqualFold applies the EqualFold predicate on the "symbol" field.
func SymbolEqualFold(v string) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldEqualFold(FieldSymbol, v))
}

// SymbolContainsFold applies the ContainsFold predicate on the "symbol" field.
func SymbolContainsFold(v string) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldContainsFold(FieldSymbol, v))
}

// FixedConversionRateEQ applies the EQ predicate on the "fixed_conversion_rate" field.
func FixedConversionRateEQ(v decimal.Decimal) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldEQ(FieldFixedConversionRate, v))
}

// FixedConversionRateNEQ applies the NEQ predicate on the "fixed_conversion_rate" field.
func FixedConversionRateNEQ(v decimal.Decimal) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldNEQ(FieldFixedConversionRate, v))
}

// FixedConversionRateIn applies the In predicate on the "fixed_conversion_rate" field.
func FixedConversionRateIn(vs ...decimal.Decimal) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldIn(FieldFixedConversionRate, vs...))
}

// FixedConversionRateNotIn applies the NotIn predicate on the "fixed_conversion_rate" field.
func FixedConversionRateNotIn(vs ...decimal.Decimal) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldNotIn(FieldFixedConversionRate, vs...))
}

// FixedConversionRateGT applies the GT predicate on the "fixed_conversion_rate" field.
func FixedConversionRateGT(v decimal.Decimal) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldGT(FieldFixedConversionRate, v))
}

// FixedConversionRateGTE applies the GTE predicate on the "fixed_conversion_rate" field.
func FixedConversionRateGTE(v decimal.Decimal) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldGTE(FieldFixedConversionRate, v))
}

// FixedConversionRateLT applies the LT predicate on the "fixed_conversion_rate" field.
func FixedConversionRateLT(v decimal.Decimal) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldLT(FieldFixedConversionRate, v))
}

// FixedConversionRateLTE applies the LTE predicate on the "fixed_conversion_rate" field.
func FixedConversionRateLTE(v decimal.Decimal) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldLTE(FieldFixedConversionRate, v))
}

// FloatingConversionRateEQ applies the EQ predicate on the "floating_conversion_rate" field.
func FloatingConversionRateEQ(v decimal.Decimal) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldEQ(FieldFloatingConversionRate, v))
}

// FloatingConversionRateNEQ applies the NEQ predicate on the "floating_conversion_rate" field.
func FloatingConversionRateNEQ(v decimal.Decimal) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldNEQ(FieldFloatingConversionRate, v))
}

// FloatingConversionRateIn applies the In predicate on the "floating_conversion_rate" field.
func FloatingConversionRateIn(vs ...decimal.Decimal) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldIn(FieldFloatingConversionRate, vs...))
}

// FloatingConversionRateNotIn applies the NotIn predicate on the "floating_conversion_rate" field.
func FloatingConversionRateNotIn(vs ...decimal.Decimal) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldNotIn(FieldFloatingConversionRate, vs...))
}

// FloatingConversionRateGT applies the GT predicate on the "floating_conversion_rate" field.
func FloatingConversionRateGT(v decimal.Decimal) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldGT(FieldFloatingConversionRate, v))
}

// FloatingConversionRateGTE applies the GTE predicate on the "floating_conversion_rate" field.
func FloatingConversionRateGTE(v decimal.Decimal) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldGTE(FieldFloatingConversionRate, v))
}

// FloatingConversionRateLT applies the LT predicate on the "floating_conversion_rate" field.
func FloatingConversionRateLT(v decimal.Decimal) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldLT(FieldFloatingConversionRate, v))
}

// FloatingConversionRateLTE applies the LTE predicate on the "floating_conversion_rate" field.
func FloatingConversionRateLTE(v decimal.Decimal) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldLTE(FieldFloatingConversionRate, v))
}

// ConversionRateTypeEQ applies the EQ predicate on the "conversion_rate_type" field.
func ConversionRateTypeEQ(v ConversionRateType) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldEQ(FieldConversionRateType, v))
}

// ConversionRateTypeNEQ applies the NEQ predicate on the "conversion_rate_type" field.
func ConversionRateTypeNEQ(v ConversionRateType) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldNEQ(FieldConversionRateType, v))
}

// ConversionRateTypeIn applies the In predicate on the "conversion_rate_type" field.
func ConversionRateTypeIn(vs ...ConversionRateType) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldIn(FieldConversionRateType, vs...))
}

// ConversionRateTypeNotIn applies the NotIn predicate on the "conversion_rate_type" field.
func ConversionRateTypeNotIn(vs ...ConversionRateType) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldNotIn(FieldConversionRateType, vs...))
}

// MaxOrderAmountEQ applies the EQ predicate on the "max_order_amount" field.
func MaxOrderAmountEQ(v decimal.Decimal) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldEQ(FieldMaxOrderAmount, v))
}

// MaxOrderAmountNEQ applies the NEQ predicate on the "max_order_amount" field.
func MaxOrderAmountNEQ(v decimal.Decimal) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldNEQ(FieldMaxOrderAmount, v))
}

// MaxOrderAmountIn applies the In predicate on the "max_order_amount" field.
func MaxOrderAmountIn(vs ...decimal.Decimal) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldIn(FieldMaxOrderAmount, vs...))
}

// MaxOrderAmountNotIn applies the NotIn predicate on the "max_order_amount" field.
func MaxOrderAmountNotIn(vs ...decimal.Decimal) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldNotIn(FieldMaxOrderAmount, vs...))
}

// MaxOrderAmountGT applies the GT predicate on the "max_order_amount" field.
func MaxOrderAmountGT(v decimal.Decimal) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldGT(FieldMaxOrderAmount, v))
}

// MaxOrderAmountGTE applies the GTE predicate on the "max_order_amount" field.
func MaxOrderAmountGTE(v decimal.Decimal) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldGTE(FieldMaxOrderAmount, v))
}

// MaxOrderAmountLT applies the LT predicate on the "max_order_amount" field.
func MaxOrderAmountLT(v decimal.Decimal) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldLT(FieldMaxOrderAmount, v))
}

// MaxOrderAmountLTE applies the LTE predicate on the "max_order_amount" field.
func MaxOrderAmountLTE(v decimal.Decimal) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldLTE(FieldMaxOrderAmount, v))
}

// MinOrderAmountEQ applies the EQ predicate on the "min_order_amount" field.
func MinOrderAmountEQ(v decimal.Decimal) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldEQ(FieldMinOrderAmount, v))
}

// MinOrderAmountNEQ applies the NEQ predicate on the "min_order_amount" field.
func MinOrderAmountNEQ(v decimal.Decimal) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldNEQ(FieldMinOrderAmount, v))
}

// MinOrderAmountIn applies the In predicate on the "min_order_amount" field.
func MinOrderAmountIn(vs ...decimal.Decimal) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldIn(FieldMinOrderAmount, vs...))
}

// MinOrderAmountNotIn applies the NotIn predicate on the "min_order_amount" field.
func MinOrderAmountNotIn(vs ...decimal.Decimal) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldNotIn(FieldMinOrderAmount, vs...))
}

// MinOrderAmountGT applies the GT predicate on the "min_order_amount" field.
func MinOrderAmountGT(v decimal.Decimal) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldGT(FieldMinOrderAmount, v))
}

// MinOrderAmountGTE applies the GTE predicate on the "min_order_amount" field.
func MinOrderAmountGTE(v decimal.Decimal) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldGTE(FieldMinOrderAmount, v))
}

// MinOrderAmountLT applies the LT predicate on the "min_order_amount" field.
func MinOrderAmountLT(v decimal.Decimal) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldLT(FieldMinOrderAmount, v))
}

// MinOrderAmountLTE applies the LTE predicate on the "min_order_amount" field.
func MinOrderAmountLTE(v decimal.Decimal) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.FieldLTE(FieldMinOrderAmount, v))
}

// HasProvider applies the HasEdge predicate on the "provider" edge.
func HasProvider() predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProviderTable, ProviderColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProviderWith applies the HasEdge predicate on the "provider" edge with a given conditions (other predicates).
func HasProviderWith(preds ...predicate.ProviderProfile) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(func(s *sql.Selector) {
		step := newProviderStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ProviderOrderToken) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ProviderOrderToken) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ProviderOrderToken) predicate.ProviderOrderToken {
	return predicate.ProviderOrderToken(sql.NotPredicates(p))
}
