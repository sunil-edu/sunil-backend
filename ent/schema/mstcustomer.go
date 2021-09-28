package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// MstCustomer holds the schema definition for the MstCustomer entity.
type MstCustomer struct {
	ent.Schema
}

// Fields of the MstCustomer.
func (MstCustomer) Fields() []ent.Field {
	return []ent.Field{
		field.String("CustDet").
			NotEmpty().
			Annotations(
				entgql.OrderField("CUSTDET"),
			),
		field.String("CustCode").
			NotEmpty().
			Annotations(
				entgql.OrderField("CUSTCODE"),
			),
		field.String("CustAddr").
			NotEmpty().
			Annotations(
				entgql.OrderField("CUSTADDR"),
			),
		field.String("CustPlace").
			NotEmpty().
			Annotations(
				entgql.OrderField("CUSTPLACE"),
			),
		field.String("CustContact").
			NotEmpty().
			Annotations(
				entgql.OrderField("CUSTCONTACT"),
			),
		field.String("CustTel").
			NotEmpty().
			Annotations(
				entgql.OrderField("CUSTTEL"),
			),
		field.String("CustEmail").
			NotEmpty().
			Annotations(
				entgql.OrderField("CUSTEMAIL"),
			),
		field.String("CustUrl").
			NotEmpty().
			Annotations(
				entgql.OrderField("CUSTURL"),
			),
		field.String("CustBanner1").
			NotEmpty().
			Annotations(
				entgql.OrderField("CUSTBANNER1"),
			),
		field.String("CustBanner2").
			NotEmpty().
			Annotations(
				entgql.OrderField("CUSTBANNER2"),
			),
		field.String("CustLogoUrl").
			NotEmpty().
			Annotations(
				entgql.OrderField("CUSTLOGOURL"),
			),
	}
}

// Edges of the MstCustomer.
func (MstCustomer) Edges() []ent.Edge {
	return nil
}
