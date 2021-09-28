// Code generated by entc, DO NOT EDIT.

package ent

// CreateMstCustomerInput represents a mutation input for creating mstcustomers.
type CreateMstCustomerInput struct {
	CustDet     string
	CustCode    string
	CustAddr    string
	CustPlace   string
	CustContact string
	CustTel     string
	CustEmail   string
	CustUrl     string
	CustBanner1 string
	CustBanner2 string
	CustLogoUrl string
}

// Mutate applies the CreateMstCustomerInput on the MstCustomerCreate builder.
func (i *CreateMstCustomerInput) Mutate(m *MstCustomerCreate) {
	m.SetCustDet(i.CustDet)
	m.SetCustCode(i.CustCode)
	m.SetCustAddr(i.CustAddr)
	m.SetCustPlace(i.CustPlace)
	m.SetCustContact(i.CustContact)
	m.SetCustTel(i.CustTel)
	m.SetCustEmail(i.CustEmail)
	m.SetCustUrl(i.CustUrl)
	m.SetCustBanner1(i.CustBanner1)
	m.SetCustBanner2(i.CustBanner2)
	m.SetCustLogoUrl(i.CustLogoUrl)
}

// SetInput applies the change-set in the CreateMstCustomerInput on the create builder.
func (c *MstCustomerCreate) SetInput(i CreateMstCustomerInput) *MstCustomerCreate {
	i.Mutate(c)
	return c
}

// UpdateMstCustomerInput represents a mutation input for updating mstcustomers.
type UpdateMstCustomerInput struct {
	CustDet     *string
	CustCode    *string
	CustAddr    *string
	CustPlace   *string
	CustContact *string
	CustTel     *string
	CustEmail   *string
	CustUrl     *string
	CustBanner1 *string
	CustBanner2 *string
	CustLogoUrl *string
}

// Mutate applies the UpdateMstCustomerInput on the MstCustomerMutation.
func (i *UpdateMstCustomerInput) Mutate(m *MstCustomerMutation) {
	if v := i.CustDet; v != nil {
		m.SetCustDet(*v)
	}
	if v := i.CustCode; v != nil {
		m.SetCustCode(*v)
	}
	if v := i.CustAddr; v != nil {
		m.SetCustAddr(*v)
	}
	if v := i.CustPlace; v != nil {
		m.SetCustPlace(*v)
	}
	if v := i.CustContact; v != nil {
		m.SetCustContact(*v)
	}
	if v := i.CustTel; v != nil {
		m.SetCustTel(*v)
	}
	if v := i.CustEmail; v != nil {
		m.SetCustEmail(*v)
	}
	if v := i.CustUrl; v != nil {
		m.SetCustUrl(*v)
	}
	if v := i.CustBanner1; v != nil {
		m.SetCustBanner1(*v)
	}
	if v := i.CustBanner2; v != nil {
		m.SetCustBanner2(*v)
	}
	if v := i.CustLogoUrl; v != nil {
		m.SetCustLogoUrl(*v)
	}
}

// SetInput applies the change-set in the UpdateMstCustomerInput on the update builder.
func (u *MstCustomerUpdate) SetInput(i UpdateMstCustomerInput) *MstCustomerUpdate {
	i.Mutate(u.Mutation())
	return u
}

// SetInput applies the change-set in the UpdateMstCustomerInput on the update-one builder.
func (u *MstCustomerUpdateOne) SetInput(i UpdateMstCustomerInput) *MstCustomerUpdateOne {
	i.Mutate(u.Mutation())
	return u
}
