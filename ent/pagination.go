// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"myeduate/ent/mstcustomer"
	"strconv"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/arsmn/fastgql/graphql/errcode"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"github.com/vmihailenco/msgpack/v5"
)

// OrderDirection defines the directions in which to order a list of items.
type OrderDirection string

const (
	// OrderDirectionAsc specifies an ascending order.
	OrderDirectionAsc OrderDirection = "ASC"
	// OrderDirectionDesc specifies a descending order.
	OrderDirectionDesc OrderDirection = "DESC"
)

// Validate the order direction value.
func (o OrderDirection) Validate() error {
	if o != OrderDirectionAsc && o != OrderDirectionDesc {
		return fmt.Errorf("%s is not a valid OrderDirection", o)
	}
	return nil
}

// String implements fmt.Stringer interface.
func (o OrderDirection) String() string {
	return string(o)
}

// MarshalGQL implements graphql.Marshaler interface.
func (o OrderDirection) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(o.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (o *OrderDirection) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("order direction %T must be a string", val)
	}
	*o = OrderDirection(str)
	return o.Validate()
}

func (o OrderDirection) reverse() OrderDirection {
	if o == OrderDirectionDesc {
		return OrderDirectionAsc
	}
	return OrderDirectionDesc
}

func (o OrderDirection) orderFunc(field string) OrderFunc {
	if o == OrderDirectionDesc {
		return Desc(field)
	}
	return Asc(field)
}

func cursorsToPredicates(direction OrderDirection, after, before *Cursor, field, idField string) []func(s *sql.Selector) {
	var predicates []func(s *sql.Selector)
	if after != nil {
		if after.Value != nil {
			var predicate func([]string, ...interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.CompositeGT
			} else {
				predicate = sql.CompositeLT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.Columns(field, idField),
					after.Value, after.ID,
				))
			})
		} else {
			var predicate func(string, interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.GT
			} else {
				predicate = sql.LT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.C(idField),
					after.ID,
				))
			})
		}
	}
	if before != nil {
		if before.Value != nil {
			var predicate func([]string, ...interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.CompositeLT
			} else {
				predicate = sql.CompositeGT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.Columns(field, idField),
					before.Value, before.ID,
				))
			})
		} else {
			var predicate func(string, interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.LT
			} else {
				predicate = sql.GT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.C(idField),
					before.ID,
				))
			})
		}
	}
	return predicates
}

// PageInfo of a connection type.
type PageInfo struct {
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *Cursor `json:"startCursor"`
	EndCursor       *Cursor `json:"endCursor"`
}

// Cursor of an edge type.
type Cursor struct {
	ID    int   `msgpack:"i"`
	Value Value `msgpack:"v,omitempty"`
}

// MarshalGQL implements graphql.Marshaler interface.
func (c Cursor) MarshalGQL(w io.Writer) {
	quote := []byte{'"'}
	w.Write(quote)
	defer w.Write(quote)
	wc := base64.NewEncoder(base64.RawStdEncoding, w)
	defer wc.Close()
	_ = msgpack.NewEncoder(wc).Encode(c)
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (c *Cursor) UnmarshalGQL(v interface{}) error {
	s, ok := v.(string)
	if !ok {
		return fmt.Errorf("%T is not a string", v)
	}
	if err := msgpack.NewDecoder(
		base64.NewDecoder(
			base64.RawStdEncoding,
			strings.NewReader(s),
		),
	).Decode(c); err != nil {
		return fmt.Errorf("cannot decode cursor: %w", err)
	}
	return nil
}

const errInvalidPagination = "INVALID_PAGINATION"

func validateFirstLast(first, last *int) (err *gqlerror.Error) {
	switch {
	case first != nil && last != nil:
		err = &gqlerror.Error{
			Message: "Passing both `first` and `last` to paginate a connection is not supported.",
		}
	case first != nil && *first < 0:
		err = &gqlerror.Error{
			Message: "`first` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	case last != nil && *last < 0:
		err = &gqlerror.Error{
			Message: "`last` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	}
	return err
}

func getCollectedField(ctx context.Context, path ...string) *graphql.CollectedField {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return nil
	}
	oc := graphql.GetOperationContext(ctx)
	field := fc.Field

walk:
	for _, name := range path {
		for _, f := range graphql.CollectFields(oc, field.Selections, nil) {
			if f.Name == name {
				field = f
				continue walk
			}
		}
		return nil
	}
	return &field
}

func hasCollectedField(ctx context.Context, path ...string) bool {
	if graphql.GetFieldContext(ctx) == nil {
		return true
	}
	return getCollectedField(ctx, path...) != nil
}

const (
	edgesField      = "edges"
	nodeField       = "node"
	pageInfoField   = "pageInfo"
	totalCountField = "totalCount"
)

// MstCustomerEdge is the edge representation of MstCustomer.
type MstCustomerEdge struct {
	Node   *MstCustomer `json:"node"`
	Cursor Cursor       `json:"cursor"`
}

// MstCustomerConnection is the connection containing edges to MstCustomer.
type MstCustomerConnection struct {
	Edges      []*MstCustomerEdge `json:"edges"`
	PageInfo   PageInfo           `json:"pageInfo"`
	TotalCount int                `json:"totalCount"`
}

// MstCustomerPaginateOption enables pagination customization.
type MstCustomerPaginateOption func(*mstCustomerPager) error

// WithMstCustomerOrder configures pagination ordering.
func WithMstCustomerOrder(order *MstCustomerOrder) MstCustomerPaginateOption {
	if order == nil {
		order = DefaultMstCustomerOrder
	}
	o := *order
	return func(pager *mstCustomerPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultMstCustomerOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithMstCustomerFilter configures pagination filter.
func WithMstCustomerFilter(filter func(*MstCustomerQuery) (*MstCustomerQuery, error)) MstCustomerPaginateOption {
	return func(pager *mstCustomerPager) error {
		if filter == nil {
			return errors.New("MstCustomerQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type mstCustomerPager struct {
	order  *MstCustomerOrder
	filter func(*MstCustomerQuery) (*MstCustomerQuery, error)
}

func newMstCustomerPager(opts []MstCustomerPaginateOption) (*mstCustomerPager, error) {
	pager := &mstCustomerPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultMstCustomerOrder
	}
	return pager, nil
}

func (p *mstCustomerPager) applyFilter(query *MstCustomerQuery) (*MstCustomerQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *mstCustomerPager) toCursor(mc *MstCustomer) Cursor {
	return p.order.Field.toCursor(mc)
}

func (p *mstCustomerPager) applyCursors(query *MstCustomerQuery, after, before *Cursor) *MstCustomerQuery {
	for _, predicate := range cursorsToPredicates(
		p.order.Direction, after, before,
		p.order.Field.field, DefaultMstCustomerOrder.Field.field,
	) {
		query = query.Where(predicate)
	}
	return query
}

func (p *mstCustomerPager) applyOrder(query *MstCustomerQuery, reverse bool) *MstCustomerQuery {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	query = query.Order(direction.orderFunc(p.order.Field.field))
	if p.order.Field != DefaultMstCustomerOrder.Field {
		query = query.Order(direction.orderFunc(DefaultMstCustomerOrder.Field.field))
	}
	return query
}

// Paginate executes the query and returns a relay based cursor connection to MstCustomer.
func (mc *MstCustomerQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...MstCustomerPaginateOption,
) (*MstCustomerConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newMstCustomerPager(opts)
	if err != nil {
		return nil, err
	}

	if mc, err = pager.applyFilter(mc); err != nil {
		return nil, err
	}

	conn := &MstCustomerConnection{Edges: []*MstCustomerEdge{}}
	if !hasCollectedField(ctx, edgesField) || first != nil && *first == 0 || last != nil && *last == 0 {
		if hasCollectedField(ctx, totalCountField) ||
			hasCollectedField(ctx, pageInfoField) {
			count, err := mc.Count(ctx)
			if err != nil {
				return nil, err
			}
			conn.TotalCount = count
			conn.PageInfo.HasNextPage = first != nil && count > 0
			conn.PageInfo.HasPreviousPage = last != nil && count > 0
		}
		return conn, nil
	}

	if (after != nil || first != nil || before != nil || last != nil) && hasCollectedField(ctx, totalCountField) {
		count, err := mc.Clone().Count(ctx)
		if err != nil {
			return nil, err
		}
		conn.TotalCount = count
	}

	mc = pager.applyCursors(mc, after, before)
	mc = pager.applyOrder(mc, last != nil)
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	if limit > 0 {
		mc = mc.Limit(limit)
	}

	if field := getCollectedField(ctx, edgesField, nodeField); field != nil {
		mc = mc.collectField(graphql.GetOperationContext(ctx), *field)
	}

	nodes, err := mc.All(ctx)
	if err != nil || len(nodes) == 0 {
		return conn, err
	}

	if len(nodes) == limit {
		conn.PageInfo.HasNextPage = first != nil
		conn.PageInfo.HasPreviousPage = last != nil
		nodes = nodes[:len(nodes)-1]
	}

	var nodeAt func(int) *MstCustomer
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *MstCustomer {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *MstCustomer {
			return nodes[i]
		}
	}

	conn.Edges = make([]*MstCustomerEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		conn.Edges[i] = &MstCustomerEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}

	conn.PageInfo.StartCursor = &conn.Edges[0].Cursor
	conn.PageInfo.EndCursor = &conn.Edges[len(conn.Edges)-1].Cursor
	if conn.TotalCount == 0 {
		conn.TotalCount = len(nodes)
	}

	return conn, nil
}

var (
	// MstCustomerOrderFieldCustDet orders MstCustomer by CustDet.
	MstCustomerOrderFieldCustDet = &MstCustomerOrderField{
		field: mstcustomer.FieldCustDet,
		toCursor: func(mc *MstCustomer) Cursor {
			return Cursor{
				ID:    mc.ID,
				Value: mc.CustDet,
			}
		},
	}
	// MstCustomerOrderFieldCustCode orders MstCustomer by CustCode.
	MstCustomerOrderFieldCustCode = &MstCustomerOrderField{
		field: mstcustomer.FieldCustCode,
		toCursor: func(mc *MstCustomer) Cursor {
			return Cursor{
				ID:    mc.ID,
				Value: mc.CustCode,
			}
		},
	}
	// MstCustomerOrderFieldCustAddr orders MstCustomer by CustAddr.
	MstCustomerOrderFieldCustAddr = &MstCustomerOrderField{
		field: mstcustomer.FieldCustAddr,
		toCursor: func(mc *MstCustomer) Cursor {
			return Cursor{
				ID:    mc.ID,
				Value: mc.CustAddr,
			}
		},
	}
	// MstCustomerOrderFieldCustPlace orders MstCustomer by CustPlace.
	MstCustomerOrderFieldCustPlace = &MstCustomerOrderField{
		field: mstcustomer.FieldCustPlace,
		toCursor: func(mc *MstCustomer) Cursor {
			return Cursor{
				ID:    mc.ID,
				Value: mc.CustPlace,
			}
		},
	}
	// MstCustomerOrderFieldCustContact orders MstCustomer by CustContact.
	MstCustomerOrderFieldCustContact = &MstCustomerOrderField{
		field: mstcustomer.FieldCustContact,
		toCursor: func(mc *MstCustomer) Cursor {
			return Cursor{
				ID:    mc.ID,
				Value: mc.CustContact,
			}
		},
	}
	// MstCustomerOrderFieldCustTel orders MstCustomer by CustTel.
	MstCustomerOrderFieldCustTel = &MstCustomerOrderField{
		field: mstcustomer.FieldCustTel,
		toCursor: func(mc *MstCustomer) Cursor {
			return Cursor{
				ID:    mc.ID,
				Value: mc.CustTel,
			}
		},
	}
	// MstCustomerOrderFieldCustEmail orders MstCustomer by CustEmail.
	MstCustomerOrderFieldCustEmail = &MstCustomerOrderField{
		field: mstcustomer.FieldCustEmail,
		toCursor: func(mc *MstCustomer) Cursor {
			return Cursor{
				ID:    mc.ID,
				Value: mc.CustEmail,
			}
		},
	}
	// MstCustomerOrderFieldCustUrl orders MstCustomer by CustUrl.
	MstCustomerOrderFieldCustUrl = &MstCustomerOrderField{
		field: mstcustomer.FieldCustUrl,
		toCursor: func(mc *MstCustomer) Cursor {
			return Cursor{
				ID:    mc.ID,
				Value: mc.CustUrl,
			}
		},
	}
	// MstCustomerOrderFieldCustBanner1 orders MstCustomer by CustBanner1.
	MstCustomerOrderFieldCustBanner1 = &MstCustomerOrderField{
		field: mstcustomer.FieldCustBanner1,
		toCursor: func(mc *MstCustomer) Cursor {
			return Cursor{
				ID:    mc.ID,
				Value: mc.CustBanner1,
			}
		},
	}
	// MstCustomerOrderFieldCustBanner2 orders MstCustomer by CustBanner2.
	MstCustomerOrderFieldCustBanner2 = &MstCustomerOrderField{
		field: mstcustomer.FieldCustBanner2,
		toCursor: func(mc *MstCustomer) Cursor {
			return Cursor{
				ID:    mc.ID,
				Value: mc.CustBanner2,
			}
		},
	}
	// MstCustomerOrderFieldCustLogoUrl orders MstCustomer by CustLogoUrl.
	MstCustomerOrderFieldCustLogoUrl = &MstCustomerOrderField{
		field: mstcustomer.FieldCustLogoUrl,
		toCursor: func(mc *MstCustomer) Cursor {
			return Cursor{
				ID:    mc.ID,
				Value: mc.CustLogoUrl,
			}
		},
	}
)

// String implement fmt.Stringer interface.
func (f MstCustomerOrderField) String() string {
	var str string
	switch f.field {
	case mstcustomer.FieldCustDet:
		str = "CUSTDET"
	case mstcustomer.FieldCustCode:
		str = "CUSTCODE"
	case mstcustomer.FieldCustAddr:
		str = "CUSTADDR"
	case mstcustomer.FieldCustPlace:
		str = "CUSTPLACE"
	case mstcustomer.FieldCustContact:
		str = "CUSTCONTACT"
	case mstcustomer.FieldCustTel:
		str = "CUSTTEL"
	case mstcustomer.FieldCustEmail:
		str = "CUSTEMAIL"
	case mstcustomer.FieldCustUrl:
		str = "CUSTURL"
	case mstcustomer.FieldCustBanner1:
		str = "CUSTBANNER1"
	case mstcustomer.FieldCustBanner2:
		str = "CUSTBANNER2"
	case mstcustomer.FieldCustLogoUrl:
		str = "CUSTLOGOURL"
	}
	return str
}

// MarshalGQL implements graphql.Marshaler interface.
func (f MstCustomerOrderField) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(f.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (f *MstCustomerOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("MstCustomerOrderField %T must be a string", v)
	}
	switch str {
	case "CUSTDET":
		*f = *MstCustomerOrderFieldCustDet
	case "CUSTCODE":
		*f = *MstCustomerOrderFieldCustCode
	case "CUSTADDR":
		*f = *MstCustomerOrderFieldCustAddr
	case "CUSTPLACE":
		*f = *MstCustomerOrderFieldCustPlace
	case "CUSTCONTACT":
		*f = *MstCustomerOrderFieldCustContact
	case "CUSTTEL":
		*f = *MstCustomerOrderFieldCustTel
	case "CUSTEMAIL":
		*f = *MstCustomerOrderFieldCustEmail
	case "CUSTURL":
		*f = *MstCustomerOrderFieldCustUrl
	case "CUSTBANNER1":
		*f = *MstCustomerOrderFieldCustBanner1
	case "CUSTBANNER2":
		*f = *MstCustomerOrderFieldCustBanner2
	case "CUSTLOGOURL":
		*f = *MstCustomerOrderFieldCustLogoUrl
	default:
		return fmt.Errorf("%s is not a valid MstCustomerOrderField", str)
	}
	return nil
}

// MstCustomerOrderField defines the ordering field of MstCustomer.
type MstCustomerOrderField struct {
	field    string
	toCursor func(*MstCustomer) Cursor
}

// MstCustomerOrder defines the ordering of MstCustomer.
type MstCustomerOrder struct {
	Direction OrderDirection         `json:"direction"`
	Field     *MstCustomerOrderField `json:"field"`
}

// DefaultMstCustomerOrder is the default ordering of MstCustomer.
var DefaultMstCustomerOrder = &MstCustomerOrder{
	Direction: OrderDirectionAsc,
	Field: &MstCustomerOrderField{
		field: mstcustomer.FieldID,
		toCursor: func(mc *MstCustomer) Cursor {
			return Cursor{ID: mc.ID}
		},
	},
}

// ToEdge converts MstCustomer into MstCustomerEdge.
func (mc *MstCustomer) ToEdge(order *MstCustomerOrder) *MstCustomerEdge {
	if order == nil {
		order = DefaultMstCustomerOrder
	}
	return &MstCustomerEdge{
		Node:   mc,
		Cursor: order.Field.toCursor(mc),
	}
}