// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/TheOguzhan/Drone-Mobile-App-Backend/ent/address"
	"github.com/TheOguzhan/Drone-Mobile-App-Backend/ent/drone"
	"github.com/TheOguzhan/Drone-Mobile-App-Backend/ent/order"
	"github.com/TheOguzhan/Drone-Mobile-App-Backend/ent/product"
	"github.com/TheOguzhan/Drone-Mobile-App-Backend/ent/user"
	"github.com/TheOguzhan/Drone-Mobile-App-Backend/ent/warehouse"
	"github.com/google/uuid"
)

// Order is the model entity for the Order schema.
type Order struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// QrCode holds the value of the "qr_code" field.
	QrCode string `json:"qr_code,omitempty"`
	// DateOfTheOrder holds the value of the "date_of_the_order" field.
	DateOfTheOrder time.Time `json:"date_of_the_order,omitempty"`
	// Completed holds the value of the "completed" field.
	Completed bool `json:"completed,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrderQuery when eager-loading is set.
	Edges                     OrderEdges `json:"edges"`
	address_address_order     *uuid.UUID
	drone_current_order       *uuid.UUID
	product_product_order     *uuid.UUID
	user_user_orders          *uuid.UUID
	warehouse_warehouse_order *uuid.UUID
}

// OrderEdges holds the relations/edges for other nodes in the graph.
type OrderEdges struct {
	// CarrierDrone holds the value of the carrier_drone edge.
	CarrierDrone *Drone `json:"carrier_drone,omitempty"`
	// UserOrder holds the value of the user_order edge.
	UserOrder *User `json:"user_order,omitempty"`
	// OrderWarehouse holds the value of the order_warehouse edge.
	OrderWarehouse *Warehouse `json:"order_warehouse,omitempty"`
	// OrderAddress holds the value of the order_address edge.
	OrderAddress *Address `json:"order_address,omitempty"`
	// OrderProduct holds the value of the order_product edge.
	OrderProduct *Product `json:"order_product,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
	// totalCount holds the count of the edges above.
	totalCount [5]map[string]int
}

// CarrierDroneOrErr returns the CarrierDrone value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderEdges) CarrierDroneOrErr() (*Drone, error) {
	if e.loadedTypes[0] {
		if e.CarrierDrone == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: drone.Label}
		}
		return e.CarrierDrone, nil
	}
	return nil, &NotLoadedError{edge: "carrier_drone"}
}

// UserOrderOrErr returns the UserOrder value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderEdges) UserOrderOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.UserOrder == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.UserOrder, nil
	}
	return nil, &NotLoadedError{edge: "user_order"}
}

// OrderWarehouseOrErr returns the OrderWarehouse value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderEdges) OrderWarehouseOrErr() (*Warehouse, error) {
	if e.loadedTypes[2] {
		if e.OrderWarehouse == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: warehouse.Label}
		}
		return e.OrderWarehouse, nil
	}
	return nil, &NotLoadedError{edge: "order_warehouse"}
}

// OrderAddressOrErr returns the OrderAddress value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderEdges) OrderAddressOrErr() (*Address, error) {
	if e.loadedTypes[3] {
		if e.OrderAddress == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: address.Label}
		}
		return e.OrderAddress, nil
	}
	return nil, &NotLoadedError{edge: "order_address"}
}

// OrderProductOrErr returns the OrderProduct value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderEdges) OrderProductOrErr() (*Product, error) {
	if e.loadedTypes[4] {
		if e.OrderProduct == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: product.Label}
		}
		return e.OrderProduct, nil
	}
	return nil, &NotLoadedError{edge: "order_product"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Order) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case order.FieldCompleted:
			values[i] = new(sql.NullBool)
		case order.FieldQrCode:
			values[i] = new(sql.NullString)
		case order.FieldDateOfTheOrder:
			values[i] = new(sql.NullTime)
		case order.FieldID:
			values[i] = new(uuid.UUID)
		case order.ForeignKeys[0]: // address_address_order
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case order.ForeignKeys[1]: // drone_current_order
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case order.ForeignKeys[2]: // product_product_order
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case order.ForeignKeys[3]: // user_user_orders
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case order.ForeignKeys[4]: // warehouse_warehouse_order
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Order", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Order fields.
func (o *Order) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case order.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				o.ID = *value
			}
		case order.FieldQrCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field qr_code", values[i])
			} else if value.Valid {
				o.QrCode = value.String
			}
		case order.FieldDateOfTheOrder:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field date_of_the_order", values[i])
			} else if value.Valid {
				o.DateOfTheOrder = value.Time
			}
		case order.FieldCompleted:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field completed", values[i])
			} else if value.Valid {
				o.Completed = value.Bool
			}
		case order.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field address_address_order", values[i])
			} else if value.Valid {
				o.address_address_order = new(uuid.UUID)
				*o.address_address_order = *value.S.(*uuid.UUID)
			}
		case order.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field drone_current_order", values[i])
			} else if value.Valid {
				o.drone_current_order = new(uuid.UUID)
				*o.drone_current_order = *value.S.(*uuid.UUID)
			}
		case order.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field product_product_order", values[i])
			} else if value.Valid {
				o.product_product_order = new(uuid.UUID)
				*o.product_product_order = *value.S.(*uuid.UUID)
			}
		case order.ForeignKeys[3]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field user_user_orders", values[i])
			} else if value.Valid {
				o.user_user_orders = new(uuid.UUID)
				*o.user_user_orders = *value.S.(*uuid.UUID)
			}
		case order.ForeignKeys[4]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field warehouse_warehouse_order", values[i])
			} else if value.Valid {
				o.warehouse_warehouse_order = new(uuid.UUID)
				*o.warehouse_warehouse_order = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryCarrierDrone queries the "carrier_drone" edge of the Order entity.
func (o *Order) QueryCarrierDrone() *DroneQuery {
	return NewOrderClient(o.config).QueryCarrierDrone(o)
}

// QueryUserOrder queries the "user_order" edge of the Order entity.
func (o *Order) QueryUserOrder() *UserQuery {
	return NewOrderClient(o.config).QueryUserOrder(o)
}

// QueryOrderWarehouse queries the "order_warehouse" edge of the Order entity.
func (o *Order) QueryOrderWarehouse() *WarehouseQuery {
	return NewOrderClient(o.config).QueryOrderWarehouse(o)
}

// QueryOrderAddress queries the "order_address" edge of the Order entity.
func (o *Order) QueryOrderAddress() *AddressQuery {
	return NewOrderClient(o.config).QueryOrderAddress(o)
}

// QueryOrderProduct queries the "order_product" edge of the Order entity.
func (o *Order) QueryOrderProduct() *ProductQuery {
	return NewOrderClient(o.config).QueryOrderProduct(o)
}

// Update returns a builder for updating this Order.
// Note that you need to call Order.Unwrap() before calling this method if this Order
// was returned from a transaction, and the transaction was committed or rolled back.
func (o *Order) Update() *OrderUpdateOne {
	return NewOrderClient(o.config).UpdateOne(o)
}

// Unwrap unwraps the Order entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (o *Order) Unwrap() *Order {
	_tx, ok := o.config.driver.(*txDriver)
	if !ok {
		panic("ent: Order is not a transactional entity")
	}
	o.config.driver = _tx.drv
	return o
}

// String implements the fmt.Stringer.
func (o *Order) String() string {
	var builder strings.Builder
	builder.WriteString("Order(")
	builder.WriteString(fmt.Sprintf("id=%v, ", o.ID))
	builder.WriteString("qr_code=")
	builder.WriteString(o.QrCode)
	builder.WriteString(", ")
	builder.WriteString("date_of_the_order=")
	builder.WriteString(o.DateOfTheOrder.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("completed=")
	builder.WriteString(fmt.Sprintf("%v", o.Completed))
	builder.WriteByte(')')
	return builder.String()
}

// Orders is a parsable slice of Order.
type Orders []*Order
