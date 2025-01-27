package odoo

import (
	"encoding/json"
	"log"
	"time"

	"github.com/xiatechs/XFuze/pkg/importer/conversion/converter"
)

// ProductModelID is the model identifier for a product within Odoo.
const ProductModelID = "product.product"

// Product represents a product within Odoo.
type Product struct {
	LastUpdate                           interface{}   `json:"__last_update"`
	AccountTagIds                        []interface{} `json:"account_tag_ids"`
	Active                               bool          `json:"active"`
	ActivityDateDeadline                 interface{}   `json:"activity_date_deadline"`
	ActivityExceptionDecoration          interface{}   `json:"activity_exception_decoration"`
	ActivityExceptionIcon                interface{}   `json:"activity_exception_icon"`
	ActivityIds                          []interface{} `json:"activity_ids"`
	ActivityState                        interface{}   `json:"activity_state"`
	ActivitySummary                      interface{}   `json:"activity_summary"`
	ActivityTypeIcon                     interface{}   `json:"activity_type_icon"`
	ActivityTypeId                       interface{}   `json:"activity_type_id"`
	ActivityUserId                       interface{}   `json:"activity_user_id"`
	AdditionalProductTagIds              []interface{} `json:"additional_product_tag_ids"`
	AllProductTagIds                     []interface{} `json:"all_product_tag_ids"`
	AttributeLineIds                     []interface{} `json:"attribute_line_ids"`
	Barcode                              interface{}   `json:"barcode"`
	CanImage1024BeZoomed                 interface{}   `json:"can_image_1024_be_zoomed"`
	CanImageVariant1024BeZoomed          interface{}   `json:"can_image_variant_1024_be_zoomed"`
	CategId                              []interface{} `json:"categ_id"`
	Code                                 interface{}   `json:"code"`
	Color                                int           `json:"color"`
	CombinationIndices                   interface{}   `json:"combination_indices"`
	CompanyId                            interface{}   `json:"company_id"`
	CostCurrencyId                       []interface{} `json:"cost_currency_id"`
	CreateDate                           interface{}   `json:"create_date"`
	CreateUid                            []interface{} `json:"create_uid"`
	CurrencyId                           []interface{} `json:"currency_id"`
	DefaultCode                          interface{}   `json:"default_code"`
	Description                          interface{}   `json:"description"`
	DescriptionPurchase                  interface{}   `json:"description_purchase"`
	DescriptionSale                      interface{}   `json:"description_sale"`
	DetailedType                         interface{}   `json:"detailed_type"`
	DisplayName                          interface{}   `json:"display_name"`
	ExpensePolicy                        interface{}   `json:"expense_policy"`
	HasConfigurableAttributes            interface{}   `json:"has_configurable_attributes"`
	HasMessage                           interface{}   `json:"has_message"`
	ID                                   int           `json:"id"`
	Image1024                            interface{}   `json:"image_1024"`
	Image128                             interface{}   `json:"image_128"`
	Image1920                            interface{}   `json:"image_1920"`
	Image256                             interface{}   `json:"image_256"`
	Image512                             interface{}   `json:"image_512"`
	ImageVariant1024                     interface{}   `json:"image_variant_1024"`
	ImageVariant128                      interface{}   `json:"image_variant_128"`
	ImageVariant1920                     interface{}   `json:"image_variant_1920"`
	ImageVariant256                      interface{}   `json:"image_variant_256"`
	ImageVariant512                      interface{}   `json:"image_variant_512"`
	InvoicePolicy                        interface{}   `json:"invoice_policy"`
	IsProductVariant                     interface{}   `json:"is_product_variant"`
	ListPrice                            int           `json:"list_price"`
	LstPrice                             int           `json:"lst_price"`
	MessageAttachmentCount               int           `json:"message_attachment_count"`
	MessageFollowerIds                   []int         `json:"message_follower_ids"`
	MessageHasError                      bool          `json:"message_has_error"`
	MessageHasErrorCounter               int           `json:"message_has_error_counter"`
	MessageHasSmsError                   bool          `json:"message_has_sms_error"`
	MessageIds                           []int         `json:"message_ids"`
	MessageIsFollower                    bool          `json:"message_is_follower"`
	MessageMainAttachmentId              interface{}   `json:"message_main_attachment_id"`
	MessageNeedaction                    bool          `json:"message_needaction"`
	MessageNeedactionCounter             int           `json:"message_needaction_counter"`
	MessagePartnerIds                    []interface{} `json:"message_partner_ids"`
	MyActivityDateDeadline               interface{}   `json:"my_activity_date_deadline"`
	Name                                 interface{}   `json:"name"`
	OptionalProductIds                   []interface{} `json:"optional_product_ids"`
	PackagingIds                         []interface{} `json:"packaging_ids"`
	PartnerRef                           interface{}   `json:"partner_ref"`
	PriceExtra                           int           `json:"price_extra"`
	PricelistItemCount                   int           `json:"pricelist_item_count"`
	Priority                             interface{}   `json:"priority"`
	ProductTagIds                        []interface{} `json:"product_tag_ids"`
	ProductTemplateAttributeValueIds     []interface{} `json:"product_template_attribute_value_ids"`
	ProductTemplateVariantValueIds       []interface{} `json:"product_template_variant_value_ids"`
	ProductTmplId                        []interface{} `json:"product_tmpl_id"`
	ProductTooltip                       interface{}   `json:"product_tooltip"`
	ProductVariantCount                  int           `json:"product_variant_count"`
	ProductVariantId                     []interface{} `json:"product_variant_id"`
	ProductVariantIds                    []int         `json:"product_variant_ids"`
	PropertyAccountExpenseId             interface{}   `json:"property_account_expense_id"`
	PropertyAccountIncomeId              interface{}   `json:"property_account_income_id"`
	PurchaseOk                           bool          `json:"purchase_ok"`
	SaleLineWarn                         interface{}   `json:"sale_line_warn"`
	SaleLineWarnMsg                      interface{}   `json:"sale_line_warn_msg"`
	SaleOk                               bool          `json:"sale_ok"`
	SalesCount                           int           `json:"sales_count"`
	SellerIds                            []interface{} `json:"seller_ids"`
	Sequence                             int           `json:"sequence"`
	ServiceType                          interface{}   `json:"service_type"`
	StandardPrice                        int           `json:"standard_price"`
	SupplierTaxesId                      []interface{} `json:"supplier_taxes_id"`
	TaxString                            interface{}   `json:"tax_string"`
	TaxesId                              []interface{} `json:"taxes_id"`
	Type                                 interface{}   `json:"type"`
	UomId                                []interface{} `json:"uom_id"`
	UomName                              interface{}   `json:"uom_name"`
	UomPoId                              []interface{} `json:"uom_po_id"`
	ValidProductTemplateAttributeLineIds []interface{} `json:"valid_product_template_attribute_line_ids"`
	VariantSellerIds                     []interface{} `json:"variant_seller_ids"`
	VisibleExpensePolicy                 interface{}   `json:"visible_expense_policy"`
	VisibleQtyConfigurator               interface{}   `json:"visible_qty_configurator"`
	Volume                               int           `json:"volume"`
	VolumeUomName                        interface{}   `json:"volume_uom_name"`
	WebsiteMessageIds                    []interface{} `json:"website_message_ids"`
	Weight                               float64       `json:"weight"`
	WeightUomName                        interface{}   `json:"weight_uom_name"`
	WriteDate                            interface{}   `json:"write_date"`
	WriteUid                             []interface{} `json:"write_uid"`
}

// hack - for demo
var (
	tempTimePlaceholder time.Time
	firstTime = true
)

// GetProductsWithinWindow gets all product.product records that have been created within the window value.
// For example if the window is 5 minutes and the time is 12:00, we fetch all products created from 11:55-12
// We return encoded JSON within this method which makes it easy to use within XFuze.
func (c *Client) GetProductsWithinWindow(_ time.Time) ([]byte, error) {
	now := time.Now()

	dtLayout := "2006-01-02 15:04:05"

	var startTime string

	if firstTime {
		firstTime = false
		startTime = now.Format(dtLayout)
	} else {
		startTime = tempTimePlaceholder.Format(dtLayout)
	}

	log.Printf("looking for products created between %s & %s", startTime, now.Format(dtLayout))

	product, err := c.SearchRead(ProductModelID, List{List{"create_date", ">=", startTime}}, nil)
	
	if err != nil {
		return nil, err
	}

	tempTimePlaceholder = now

	return json.Marshal(product)
}

// GetProduct returns a JSON encoded byte slice of an Odoo product.
func (c *Client) GetProduct(id int) ([]byte, error) {
	product, err := c.Read(ProductModelID, []int{id}, nil)
	if err != nil {
		return nil, err
	}

	return json.Marshal(product)
}

// ProductTemplateID is the model identifier for a product template within Odoo.
const ProductTemplateID = "product.template"

// ProductTemplate is the struct used to create a product within Odoo.
type ProductTemplate struct {
	AttributeLineIds        []interface{}   `json:"attribute_line_ids"`
	Image1920               bool            `json:"image_1920"`
	LastUpdate              bool            `json:"__last_update"`
	Name                    string          `json:"name"`
	SaleOk                  bool            `json:"sale_ok"`
	PurchaseOk              bool            `json:"purchase_ok"`
	Active                  bool            `json:"active"`
	Type                    string          `json:"type"`
	CategId                 int             `json:"categ_id"`
	DefaultCode             bool            `json:"default_code"`
	Barcode                 bool            `json:"barcode"`
	ListPrice               int             `json:"list_price"`
	StandardPrice           int             `json:"standard_price"`
	CompanyId               bool            `json:"company_id"`
	UomId                   int             `json:"uom_id"`
	UomPoId                 int             `json:"uom_po_id"`
	Description             bool            `json:"description"`
	DescriptionSale         bool            `json:"description_sale"`
	RouteIds                [][]interface{} `json:"route_ids"`
	ResponsibleId           bool            `json:"responsible_id"`
	Weight                  int             `json:"weight"`
	Volume                  int             `json:"volume"`
	SaleDelay               int             `json:"sale_delay"`
	Tracking                string          `json:"tracking"`
	PropertyStockProduction int             `json:"property_stock_production"`
	PropertyStockInventory  int             `json:"property_stock_inventory"`
	PackagingIds            []interface{}   `json:"packaging_ids"`
	DescriptionPickingout   bool            `json:"description_pickingout"`
	DescriptionPickingin    bool            `json:"description_pickingin"`
	DescriptionPicking      bool            `json:"description_picking"`
	MessageFollowerIds      []interface{}   `json:"message_follower_ids"`
	ActivityIds             []interface{}   `json:"activity_ids"`
	MessageIds              []interface{}   `json:"message_ids"`
}

// CreateProduct creates a new product within Odoo
func (c *Client) CreateProduct(template ProductTemplate) (int, error) {
	bytes, err := json.Marshal(template)
	if err != nil {
		return 0, err
	}

	conv := converter.SelectFormat("json")

	mapConversion, err := conv.BytesToMap(bytes)
	if err != nil {
		return 0, nil
	}

	productID, err := c.Create(ProductTemplateID, mapConversion[0])
	if err != nil {
		return 0, err
	}

	return productID, nil
}
