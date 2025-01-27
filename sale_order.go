package odoo

import (
	"encoding/json"

	"github.com/xiatechs/XFuze/pkg/importer/conversion/converter"
)

// SaleOrderModelID is the model identifier we use to fetch a SaleOrder from Odoo.
const SaleOrderModelID = "sale.order"
const SaleOrderLineModelID = "sale.order.line"

// SaleOrder represents a sales order within Odoo.
// Required fields when creating: PartnerID & DateOrder. PartnerID is the customer this order is for, they must
// already be stored within Odoo.
type SaleOrder struct {
	ID                       int             `json:"id"`
	AuthorizedTransactionIds []interface{}   `json:"authorized_transaction_ids"`
	State                    string          `json:"state"`
	PartnerCreditWarning     string          `json:"partner_credit_warning,omitempty"`
	InvoiceCount             int             `json:"invoice_count"`
	Name                     string          `json:"name"`
	PartnerId                interface{}     `json:"partner_id"`
	SaleOrderTemplateId      interface{}     `json:"sale_order_template_id"`
	ValidityDate             bool            `json:"validity_date"`
	DateOrder                string          `json:"date_order"`
	ShowUpdatePricelist      bool            `json:"show_update_pricelist"`
	PricelistId              interface{}     `json:"pricelist_id"`
	CompanyId                []interface{}   `json:"company_id"`
	CurrencyId               interface{}     `json:"currency_id"`
	TaxCountryId             interface{}     `json:"tax_country_id"`
	PaymentTermId            interface{}     `json:"payment_term_id"`
	OrderLine                [][]interface{} `json:"order_line"`
	Note                     string          `json:"note"`
	TaxTotals                struct {
		AmountUntaxed          float64 `json:"amount_untaxed"`
		AmountTotal            float64 `json:"amount_total"`
		FormattedAmountTotal   string  `json:"formatted_amount_total"`
		FormattedAmountUntaxed string  `json:"formatted_amount_untaxed"`
		GroupsBySubtotal       struct {
		} `json:"groups_by_subtotal"`
		Subtotals      []interface{} `json:"subtotals"`
		SubtotalsOrder []interface{} `json:"subtotals_order"`
		DisplayTaxBase bool          `json:"display_tax_base"`
	} `json:"tax_totals"`
	SaleOrderOptionIds []interface{} `json:"sale_order_option_ids"`
	UserId             interface{}   `json:"user_id"`
	TeamId             interface{}   `json:"team_id"`
	RequireSignature   bool          `json:"require_signature"`
	RequirePayment     bool          `json:"require_payment"`
	Reference          bool          `json:"reference"`
	ClientOrderRef     interface{}   `json:"client_order_ref"`
	TagIds             []interface{} `json:"tag_ids"`
	ShowUpdateFpos     bool          `json:"show_update_fpos"`
	FiscalPositionId   interface{}   `json:"fiscal_position_id"`
	PartnerInvoiceId   interface{}   `json:"partner_invoice_id"`
	InvoiceStatus      string        `json:"invoice_status"`
	CommitmentDate     interface{}   `json:"commitment_date"`
	ExpectedDate       interface{}   `json:"expected_date"`
	Origin             bool          `json:"origin"`
	CampaignId         interface{}   `json:"campaign_id"`
	MediumId           interface{}   `json:"medium_id"`
	SourceId           interface{}   `json:"source_id"`
	DisplayName        string        `json:"display_name"`
}

// GetSalesOrder fetches a sales order from Odoo.
func (c *Client) GetSalesOrder(id int) ([]byte, error) {
	saleOrder, err := c.Read(SaleOrderModelID, []int{id}, nil)
	if err != nil {
		return nil, err
	}

	return json.Marshal(saleOrder)
}

// CreateSalesOrder creates a new sales order within Odoo
func (c *Client) CreateSalesOrder(order []byte) (int, error) {
	conv := converter.SelectFormat("json")

	mapConversion, err := conv.BytesToMap(order)
	if err != nil {
		return 0, nil
	}

	saleOrderID, err := c.Create(SaleOrderModelID, mapConversion[0])
	if err != nil {
		return 0, err
	}

	return saleOrderID, nil
}

// CreateSalesOrderLineItem
func (c *Client) CreateSalesOrderLineItem(orderLine []byte) (int, error) {
	conv := converter.SelectFormat("json")

	mapConversion, err := conv.BytesToMap(orderLine)
	if err != nil {
		return 0, nil
	}

	saleOrderID, err := c.Create(SaleOrderLineModelID, mapConversion[0])
	if err != nil {
		return 0, err
	}

	return saleOrderID, nil
}

// OrderLineItem is a product that gets added to an order.
// Required fields: ProductId, ProductUom
type OrderLineItem struct {
	OrderID                           int             `json:"order_id"`
	Sequence                          int             `json:"sequence"`
	DisplayType                       bool            `json:"display_type"`
	ProductId                         int             `json:"product_id"`
	ProductTemplateId                 int             `json:"product_template_id"`
	ProductCustomAttributeValueIds    []interface{}   `json:"product_custom_attribute_value_ids"`
	ProductNoVariantAttributeValueIds [][]interface{} `json:"product_no_variant_attribute_value_ids"`
	Name                              string          `json:"name"`
	ProductUomQty                     int             `json:"product_uom_qty"`
	QtyDelivered                      int             `json:"qty_delivered"`
	ProductUom                        interface{}     `json:"product_uom"`
	CustomerLead                      int             `json:"customer_lead"`
	PriceUnit                         int             `json:"price_unit"`
	TaxId                             [][]interface{} `json:"tax_id"`
	IsDownpayment                     bool            `json:"is_downpayment"`
}
