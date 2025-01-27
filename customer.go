package odoo

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/xiatechs/XFuze/pkg/importer/conversion/converter"
)

// CustomerNotFound is returned if we don't find the customer within Odoo.
var CustomerNotFound = fmt.Errorf("customer not found")

// Customer represents a customer within Odoo.
type Customer struct {
	Id                            int         `json:"id"`
	SameVatPartnerId              interface{} `json:"same_vat_partner_id"`
	PartnerGid                    interface{} `json:"partner_gid"`
	AdditionalInfo                interface{} `json:"additional_info"`
	SaleOrderCount                interface{} `json:"sale_order_count"`
	CurrencyId                    interface{} `json:"currency_id"`
	TotalInvoiced                 float64     `json:"total_invoiced"`
	PaymentTokenCount             int         `json:"payment_token_count"`
	Image1920                     interface{} `json:"image_1920"`
	LastUpdate                    interface{} `json:"__last_update"`
	IsCompany                     bool        `json:"is_company"`
	CommercialPartnerId           interface{} `json:"commercial_partner_id"`
	Active                        bool        `json:"active"`
	CompanyType                   interface{} `json:"company_type"`
	Name                          interface{} `json:"name"`
	ParentId                      interface{} `json:"parent_id"`
	CompanyName                   interface{} `json:"company_name"`
	Type                          interface{} `json:"type"`
	Street                        interface{} `json:"street"`
	Street2                       interface{} `json:"street2"`
	City                          interface{} `json:"city"`
	StateId                       interface{} `json:"state_id"`
	Zip                           interface{} `json:"zip"`
	CountryId                     interface{} `json:"country_id"`
	Vat                           interface{} `json:"vat"`
	Function                      interface{} `json:"function"`
	PhoneBlacklisted              bool        `json:"phone_blacklisted"`
	MobileBlacklisted             bool        `json:"mobile_blacklisted"`
	Phone                         interface{} `json:"phone"`
	Mobile                        interface{} `json:"mobile"`
	PhoneSanitized                interface{} `json:"phone_sanitized"`
	UserIds                       interface{} `json:"user_ids"`
	IsBlacklisted                 bool        `json:"is_blacklisted"`
	Email                         string      `json:"email"`
	Website                       interface{} `json:"website"`
	Title                         interface{} `json:"title"`
	ActiveLangCount               int         `json:"active_lang_count"`
	Lang                          interface{} `json:"lang"`
	CategoryId                    interface{} `json:"category_id"`
	ChildIds                      interface{} `json:"child_ids"`
	UserId                        interface{} `json:"user_id"`
	TeamId                        interface{} `json:"team_id"`
	PropertyPaymentTermId         interface{} `json:"property_payment_term_id"`
	PropertySupplierPaymentTermId interface{} `json:"property_supplier_payment_term_id"`
	PropertyPaymentMethodId       interface{} `json:"property_payment_method_id"`
	PropertyAccountPositionId     interface{} `json:"property_account_position_id"`
	Ref                           interface{} `json:"ref"`
	CompanyId                     interface{} `json:"company_id"`
	IndustryId                    interface{} `json:"industry_id"`
	PropertyStockCustomer         interface{} `json:"property_stock_customer"`
	PropertyStockSupplier         interface{} `json:"property_stock_supplier"`
	BankIds                       interface{} `json:"bank_ids"`
	PropertyAccountReceivableId   interface{} `json:"property_account_receivable_id"`
	PropertyAccountPayableId      interface{} `json:"property_account_payable_id"`
	Comment                       interface{} `json:"comment"`
	SaleWarn                      interface{} `json:"sale_warn"`
	SaleWarnMsg                   interface{} `json:"sale_warn_msg"`
	InvoiceWarn                   interface{} `json:"invoice_warn"`
	InvoiceWarnMsg                interface{} `json:"invoice_warn_msg"`
	PickingWarn                   interface{} `json:"picking_warn"`
	PickingWarnMsg                interface{} `json:"picking_warn_msg"`
	MessageFollowerIds            interface{} `json:"message_follower_ids"`
	ActivityIds                   interface{} `json:"activity_ids"`
	MessageIds                    interface{} `json:"message_ids"`
	DisplayName                   string      `json:"display_name"`
}

const CustomerModelID = "res.partner"

// GetCustomer gets a customer from Odoo.
func (c *Client) GetCustomer(id int) ([]byte, error) {
	cust, err := c.Read(CustomerModelID, []int{id}, nil)
	if err != nil {
		return nil, err
	}

	return json.Marshal(cust)
}

// GetCustomerFromEmail gets a customer via an email address. If the customer isn't found we return a
// NotFound error.
func (c *Client) GetCustomerFromEmail(email string) (*Customer, error) {
	log.Printf("looking for customer with email %s", email)

	resp, err := c.SearchRead("res.partner", List{List{"email", "ilike", email}}, nil)
	if err != nil {
		return nil, err
	}

	if len(resp) == 0 {
		return nil, CustomerNotFound
	}

	bytes, err := json.Marshal(resp[0])
	if err != nil {
		return nil, err
	}

	var customer Customer
	err = json.Unmarshal(bytes, &customer)
	if err != nil {
		return nil, err
	}

	return &customer, nil
}

// CreateCustomer creates a new customer within Odoo.
func (c *Client) CreateCustomer(customer Customer) (int, error) {
	bytes, err := json.Marshal(customer)
	if err != nil {
		return 0, err
	}

	conv := converter.SelectFormat("json")

	mapConversion, err := conv.BytesToMap(bytes)
	if err != nil {
		return 0, nil
	}

	custID, err := c.Create(CustomerModelID, mapConversion[0])
	if err != nil {
		return 0, err
	}

	return custID, nil
}
