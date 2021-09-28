// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"myeduate/ent/mstcustomer"
	"myeduate/ent/predicate"
	"sync"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeMstCustomer = "MstCustomer"
)

// MstCustomerMutation represents an operation that mutates the MstCustomer nodes in the graph.
type MstCustomerMutation struct {
	config
	op            Op
	typ           string
	id            *int
	_CustDet      *string
	_CustCode     *string
	_CustAddr     *string
	_CustPlace    *string
	_CustContact  *string
	_CustTel      *string
	_CustEmail    *string
	_CustUrl      *string
	_CustBanner1  *string
	_CustBanner2  *string
	_CustLogoUrl  *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*MstCustomer, error)
	predicates    []predicate.MstCustomer
}

var _ ent.Mutation = (*MstCustomerMutation)(nil)

// mstcustomerOption allows management of the mutation configuration using functional options.
type mstcustomerOption func(*MstCustomerMutation)

// newMstCustomerMutation creates new mutation for the MstCustomer entity.
func newMstCustomerMutation(c config, op Op, opts ...mstcustomerOption) *MstCustomerMutation {
	m := &MstCustomerMutation{
		config:        c,
		op:            op,
		typ:           TypeMstCustomer,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withMstCustomerID sets the ID field of the mutation.
func withMstCustomerID(id int) mstcustomerOption {
	return func(m *MstCustomerMutation) {
		var (
			err   error
			once  sync.Once
			value *MstCustomer
		)
		m.oldValue = func(ctx context.Context) (*MstCustomer, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().MstCustomer.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withMstCustomer sets the old MstCustomer of the mutation.
func withMstCustomer(node *MstCustomer) mstcustomerOption {
	return func(m *MstCustomerMutation) {
		m.oldValue = func(context.Context) (*MstCustomer, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m MstCustomerMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m MstCustomerMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *MstCustomerMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetCustDet sets the "CustDet" field.
func (m *MstCustomerMutation) SetCustDet(s string) {
	m._CustDet = &s
}

// CustDet returns the value of the "CustDet" field in the mutation.
func (m *MstCustomerMutation) CustDet() (r string, exists bool) {
	v := m._CustDet
	if v == nil {
		return
	}
	return *v, true
}

// OldCustDet returns the old "CustDet" field's value of the MstCustomer entity.
// If the MstCustomer object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MstCustomerMutation) OldCustDet(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCustDet is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCustDet requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCustDet: %w", err)
	}
	return oldValue.CustDet, nil
}

// ResetCustDet resets all changes to the "CustDet" field.
func (m *MstCustomerMutation) ResetCustDet() {
	m._CustDet = nil
}

// SetCustCode sets the "CustCode" field.
func (m *MstCustomerMutation) SetCustCode(s string) {
	m._CustCode = &s
}

// CustCode returns the value of the "CustCode" field in the mutation.
func (m *MstCustomerMutation) CustCode() (r string, exists bool) {
	v := m._CustCode
	if v == nil {
		return
	}
	return *v, true
}

// OldCustCode returns the old "CustCode" field's value of the MstCustomer entity.
// If the MstCustomer object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MstCustomerMutation) OldCustCode(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCustCode is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCustCode requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCustCode: %w", err)
	}
	return oldValue.CustCode, nil
}

// ResetCustCode resets all changes to the "CustCode" field.
func (m *MstCustomerMutation) ResetCustCode() {
	m._CustCode = nil
}

// SetCustAddr sets the "CustAddr" field.
func (m *MstCustomerMutation) SetCustAddr(s string) {
	m._CustAddr = &s
}

// CustAddr returns the value of the "CustAddr" field in the mutation.
func (m *MstCustomerMutation) CustAddr() (r string, exists bool) {
	v := m._CustAddr
	if v == nil {
		return
	}
	return *v, true
}

// OldCustAddr returns the old "CustAddr" field's value of the MstCustomer entity.
// If the MstCustomer object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MstCustomerMutation) OldCustAddr(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCustAddr is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCustAddr requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCustAddr: %w", err)
	}
	return oldValue.CustAddr, nil
}

// ResetCustAddr resets all changes to the "CustAddr" field.
func (m *MstCustomerMutation) ResetCustAddr() {
	m._CustAddr = nil
}

// SetCustPlace sets the "CustPlace" field.
func (m *MstCustomerMutation) SetCustPlace(s string) {
	m._CustPlace = &s
}

// CustPlace returns the value of the "CustPlace" field in the mutation.
func (m *MstCustomerMutation) CustPlace() (r string, exists bool) {
	v := m._CustPlace
	if v == nil {
		return
	}
	return *v, true
}

// OldCustPlace returns the old "CustPlace" field's value of the MstCustomer entity.
// If the MstCustomer object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MstCustomerMutation) OldCustPlace(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCustPlace is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCustPlace requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCustPlace: %w", err)
	}
	return oldValue.CustPlace, nil
}

// ResetCustPlace resets all changes to the "CustPlace" field.
func (m *MstCustomerMutation) ResetCustPlace() {
	m._CustPlace = nil
}

// SetCustContact sets the "CustContact" field.
func (m *MstCustomerMutation) SetCustContact(s string) {
	m._CustContact = &s
}

// CustContact returns the value of the "CustContact" field in the mutation.
func (m *MstCustomerMutation) CustContact() (r string, exists bool) {
	v := m._CustContact
	if v == nil {
		return
	}
	return *v, true
}

// OldCustContact returns the old "CustContact" field's value of the MstCustomer entity.
// If the MstCustomer object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MstCustomerMutation) OldCustContact(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCustContact is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCustContact requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCustContact: %w", err)
	}
	return oldValue.CustContact, nil
}

// ResetCustContact resets all changes to the "CustContact" field.
func (m *MstCustomerMutation) ResetCustContact() {
	m._CustContact = nil
}

// SetCustTel sets the "CustTel" field.
func (m *MstCustomerMutation) SetCustTel(s string) {
	m._CustTel = &s
}

// CustTel returns the value of the "CustTel" field in the mutation.
func (m *MstCustomerMutation) CustTel() (r string, exists bool) {
	v := m._CustTel
	if v == nil {
		return
	}
	return *v, true
}

// OldCustTel returns the old "CustTel" field's value of the MstCustomer entity.
// If the MstCustomer object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MstCustomerMutation) OldCustTel(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCustTel is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCustTel requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCustTel: %w", err)
	}
	return oldValue.CustTel, nil
}

// ResetCustTel resets all changes to the "CustTel" field.
func (m *MstCustomerMutation) ResetCustTel() {
	m._CustTel = nil
}

// SetCustEmail sets the "CustEmail" field.
func (m *MstCustomerMutation) SetCustEmail(s string) {
	m._CustEmail = &s
}

// CustEmail returns the value of the "CustEmail" field in the mutation.
func (m *MstCustomerMutation) CustEmail() (r string, exists bool) {
	v := m._CustEmail
	if v == nil {
		return
	}
	return *v, true
}

// OldCustEmail returns the old "CustEmail" field's value of the MstCustomer entity.
// If the MstCustomer object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MstCustomerMutation) OldCustEmail(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCustEmail is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCustEmail requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCustEmail: %w", err)
	}
	return oldValue.CustEmail, nil
}

// ResetCustEmail resets all changes to the "CustEmail" field.
func (m *MstCustomerMutation) ResetCustEmail() {
	m._CustEmail = nil
}

// SetCustUrl sets the "CustUrl" field.
func (m *MstCustomerMutation) SetCustUrl(s string) {
	m._CustUrl = &s
}

// CustUrl returns the value of the "CustUrl" field in the mutation.
func (m *MstCustomerMutation) CustUrl() (r string, exists bool) {
	v := m._CustUrl
	if v == nil {
		return
	}
	return *v, true
}

// OldCustUrl returns the old "CustUrl" field's value of the MstCustomer entity.
// If the MstCustomer object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MstCustomerMutation) OldCustUrl(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCustUrl is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCustUrl requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCustUrl: %w", err)
	}
	return oldValue.CustUrl, nil
}

// ResetCustUrl resets all changes to the "CustUrl" field.
func (m *MstCustomerMutation) ResetCustUrl() {
	m._CustUrl = nil
}

// SetCustBanner1 sets the "CustBanner1" field.
func (m *MstCustomerMutation) SetCustBanner1(s string) {
	m._CustBanner1 = &s
}

// CustBanner1 returns the value of the "CustBanner1" field in the mutation.
func (m *MstCustomerMutation) CustBanner1() (r string, exists bool) {
	v := m._CustBanner1
	if v == nil {
		return
	}
	return *v, true
}

// OldCustBanner1 returns the old "CustBanner1" field's value of the MstCustomer entity.
// If the MstCustomer object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MstCustomerMutation) OldCustBanner1(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCustBanner1 is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCustBanner1 requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCustBanner1: %w", err)
	}
	return oldValue.CustBanner1, nil
}

// ResetCustBanner1 resets all changes to the "CustBanner1" field.
func (m *MstCustomerMutation) ResetCustBanner1() {
	m._CustBanner1 = nil
}

// SetCustBanner2 sets the "CustBanner2" field.
func (m *MstCustomerMutation) SetCustBanner2(s string) {
	m._CustBanner2 = &s
}

// CustBanner2 returns the value of the "CustBanner2" field in the mutation.
func (m *MstCustomerMutation) CustBanner2() (r string, exists bool) {
	v := m._CustBanner2
	if v == nil {
		return
	}
	return *v, true
}

// OldCustBanner2 returns the old "CustBanner2" field's value of the MstCustomer entity.
// If the MstCustomer object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MstCustomerMutation) OldCustBanner2(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCustBanner2 is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCustBanner2 requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCustBanner2: %w", err)
	}
	return oldValue.CustBanner2, nil
}

// ResetCustBanner2 resets all changes to the "CustBanner2" field.
func (m *MstCustomerMutation) ResetCustBanner2() {
	m._CustBanner2 = nil
}

// SetCustLogoUrl sets the "CustLogoUrl" field.
func (m *MstCustomerMutation) SetCustLogoUrl(s string) {
	m._CustLogoUrl = &s
}

// CustLogoUrl returns the value of the "CustLogoUrl" field in the mutation.
func (m *MstCustomerMutation) CustLogoUrl() (r string, exists bool) {
	v := m._CustLogoUrl
	if v == nil {
		return
	}
	return *v, true
}

// OldCustLogoUrl returns the old "CustLogoUrl" field's value of the MstCustomer entity.
// If the MstCustomer object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MstCustomerMutation) OldCustLogoUrl(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCustLogoUrl is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCustLogoUrl requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCustLogoUrl: %w", err)
	}
	return oldValue.CustLogoUrl, nil
}

// ResetCustLogoUrl resets all changes to the "CustLogoUrl" field.
func (m *MstCustomerMutation) ResetCustLogoUrl() {
	m._CustLogoUrl = nil
}

// Where appends a list predicates to the MstCustomerMutation builder.
func (m *MstCustomerMutation) Where(ps ...predicate.MstCustomer) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *MstCustomerMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (MstCustomer).
func (m *MstCustomerMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *MstCustomerMutation) Fields() []string {
	fields := make([]string, 0, 11)
	if m._CustDet != nil {
		fields = append(fields, mstcustomer.FieldCustDet)
	}
	if m._CustCode != nil {
		fields = append(fields, mstcustomer.FieldCustCode)
	}
	if m._CustAddr != nil {
		fields = append(fields, mstcustomer.FieldCustAddr)
	}
	if m._CustPlace != nil {
		fields = append(fields, mstcustomer.FieldCustPlace)
	}
	if m._CustContact != nil {
		fields = append(fields, mstcustomer.FieldCustContact)
	}
	if m._CustTel != nil {
		fields = append(fields, mstcustomer.FieldCustTel)
	}
	if m._CustEmail != nil {
		fields = append(fields, mstcustomer.FieldCustEmail)
	}
	if m._CustUrl != nil {
		fields = append(fields, mstcustomer.FieldCustUrl)
	}
	if m._CustBanner1 != nil {
		fields = append(fields, mstcustomer.FieldCustBanner1)
	}
	if m._CustBanner2 != nil {
		fields = append(fields, mstcustomer.FieldCustBanner2)
	}
	if m._CustLogoUrl != nil {
		fields = append(fields, mstcustomer.FieldCustLogoUrl)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *MstCustomerMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case mstcustomer.FieldCustDet:
		return m.CustDet()
	case mstcustomer.FieldCustCode:
		return m.CustCode()
	case mstcustomer.FieldCustAddr:
		return m.CustAddr()
	case mstcustomer.FieldCustPlace:
		return m.CustPlace()
	case mstcustomer.FieldCustContact:
		return m.CustContact()
	case mstcustomer.FieldCustTel:
		return m.CustTel()
	case mstcustomer.FieldCustEmail:
		return m.CustEmail()
	case mstcustomer.FieldCustUrl:
		return m.CustUrl()
	case mstcustomer.FieldCustBanner1:
		return m.CustBanner1()
	case mstcustomer.FieldCustBanner2:
		return m.CustBanner2()
	case mstcustomer.FieldCustLogoUrl:
		return m.CustLogoUrl()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *MstCustomerMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case mstcustomer.FieldCustDet:
		return m.OldCustDet(ctx)
	case mstcustomer.FieldCustCode:
		return m.OldCustCode(ctx)
	case mstcustomer.FieldCustAddr:
		return m.OldCustAddr(ctx)
	case mstcustomer.FieldCustPlace:
		return m.OldCustPlace(ctx)
	case mstcustomer.FieldCustContact:
		return m.OldCustContact(ctx)
	case mstcustomer.FieldCustTel:
		return m.OldCustTel(ctx)
	case mstcustomer.FieldCustEmail:
		return m.OldCustEmail(ctx)
	case mstcustomer.FieldCustUrl:
		return m.OldCustUrl(ctx)
	case mstcustomer.FieldCustBanner1:
		return m.OldCustBanner1(ctx)
	case mstcustomer.FieldCustBanner2:
		return m.OldCustBanner2(ctx)
	case mstcustomer.FieldCustLogoUrl:
		return m.OldCustLogoUrl(ctx)
	}
	return nil, fmt.Errorf("unknown MstCustomer field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *MstCustomerMutation) SetField(name string, value ent.Value) error {
	switch name {
	case mstcustomer.FieldCustDet:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCustDet(v)
		return nil
	case mstcustomer.FieldCustCode:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCustCode(v)
		return nil
	case mstcustomer.FieldCustAddr:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCustAddr(v)
		return nil
	case mstcustomer.FieldCustPlace:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCustPlace(v)
		return nil
	case mstcustomer.FieldCustContact:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCustContact(v)
		return nil
	case mstcustomer.FieldCustTel:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCustTel(v)
		return nil
	case mstcustomer.FieldCustEmail:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCustEmail(v)
		return nil
	case mstcustomer.FieldCustUrl:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCustUrl(v)
		return nil
	case mstcustomer.FieldCustBanner1:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCustBanner1(v)
		return nil
	case mstcustomer.FieldCustBanner2:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCustBanner2(v)
		return nil
	case mstcustomer.FieldCustLogoUrl:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCustLogoUrl(v)
		return nil
	}
	return fmt.Errorf("unknown MstCustomer field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *MstCustomerMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *MstCustomerMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *MstCustomerMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown MstCustomer numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *MstCustomerMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *MstCustomerMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *MstCustomerMutation) ClearField(name string) error {
	return fmt.Errorf("unknown MstCustomer nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *MstCustomerMutation) ResetField(name string) error {
	switch name {
	case mstcustomer.FieldCustDet:
		m.ResetCustDet()
		return nil
	case mstcustomer.FieldCustCode:
		m.ResetCustCode()
		return nil
	case mstcustomer.FieldCustAddr:
		m.ResetCustAddr()
		return nil
	case mstcustomer.FieldCustPlace:
		m.ResetCustPlace()
		return nil
	case mstcustomer.FieldCustContact:
		m.ResetCustContact()
		return nil
	case mstcustomer.FieldCustTel:
		m.ResetCustTel()
		return nil
	case mstcustomer.FieldCustEmail:
		m.ResetCustEmail()
		return nil
	case mstcustomer.FieldCustUrl:
		m.ResetCustUrl()
		return nil
	case mstcustomer.FieldCustBanner1:
		m.ResetCustBanner1()
		return nil
	case mstcustomer.FieldCustBanner2:
		m.ResetCustBanner2()
		return nil
	case mstcustomer.FieldCustLogoUrl:
		m.ResetCustLogoUrl()
		return nil
	}
	return fmt.Errorf("unknown MstCustomer field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *MstCustomerMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *MstCustomerMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *MstCustomerMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *MstCustomerMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *MstCustomerMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *MstCustomerMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *MstCustomerMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown MstCustomer unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *MstCustomerMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown MstCustomer edge %s", name)
}