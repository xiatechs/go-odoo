package odoo

import (
	"fmt"
)

// AccountAnalyticTag represents account.analytic.tag model.
type AccountAnalyticTag struct {
	LastUpdate                 *Time     `xmlrpc:"__last_update,omitempty"`
	Active                     *Bool     `xmlrpc:"active,omitempty"`
	ActiveAnalyticDistribution *Bool     `xmlrpc:"active_analytic_distribution,omitempty"`
	AnalyticDistributionIds    *Relation `xmlrpc:"analytic_distribution_ids,omitempty"`
	Color                      *Int      `xmlrpc:"color,omitempty"`
	CompanyId                  *Many2One `xmlrpc:"company_id,omitempty"`
	CreateDate                 *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid                  *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName                *String   `xmlrpc:"display_name,omitempty"`
	Id                         *Int      `xmlrpc:"id,omitempty"`
	Name                       *String   `xmlrpc:"name,omitempty"`
	WriteDate                  *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid                   *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AccountAnalyticTags represents array of account.analytic.tag model.
type AccountAnalyticTags []AccountAnalyticTag

// AccountAnalyticTagModel is the odoo model name.
const AccountAnalyticTagModel = "account.analytic.tag"

// Many2One convert AccountAnalyticTag to *Many2One.
func (aat *AccountAnalyticTag) Many2One() *Many2One {
	return NewMany2One(aat.Id.Get(), "")
}

// CreateAccountAnalyticTag creates a new account.analytic.tag model and returns its id.
func (c *Client) CreateAccountAnalyticTag(aat *AccountAnalyticTag) (int64, error) {
	return c.Create(AccountAnalyticTagModel, aat)
}

// UpdateAccountAnalyticTag updates an existing account.analytic.tag record.
func (c *Client) UpdateAccountAnalyticTag(aat *AccountAnalyticTag) error {
	return c.UpdateAccountAnalyticTags([]int64{aat.Id.Get()}, aat)
}

// UpdateAccountAnalyticTags updates existing account.analytic.tag records.
// All records (represented by ids) will be updated by aat values.
func (c *Client) UpdateAccountAnalyticTags(ids []int64, aat *AccountAnalyticTag) error {
	return c.Update(AccountAnalyticTagModel, ids, aat)
}

// DeleteAccountAnalyticTag deletes an existing account.analytic.tag record.
func (c *Client) DeleteAccountAnalyticTag(id int64) error {
	return c.DeleteAccountAnalyticTags([]int64{id})
}

// DeleteAccountAnalyticTags deletes existing account.analytic.tag records.
func (c *Client) DeleteAccountAnalyticTags(ids []int64) error {
	return c.Delete(AccountAnalyticTagModel, ids)
}

// GetAccountAnalyticTag gets account.analytic.tag existing record.
func (c *Client) GetAccountAnalyticTag(id int64) (*AccountAnalyticTag, error) {
	aats, err := c.GetAccountAnalyticTags([]int64{id})
	if err != nil {
		return nil, err
	}
	if aats != nil && len(*aats) > 0 {
		return &((*aats)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.analytic.tag not found", id)
}

// GetAccountAnalyticTags gets account.analytic.tag existing records.
func (c *Client) GetAccountAnalyticTags(ids []int64) (*AccountAnalyticTags, error) {
	aats := &AccountAnalyticTags{}
	if err := c.Read(AccountAnalyticTagModel, ids, nil, aats); err != nil {
		return nil, err
	}
	return aats, nil
}

// FindAccountAnalyticTag finds account.analytic.tag record by querying it with criteria.
func (c *Client) FindAccountAnalyticTag(criteria *Criteria) (*AccountAnalyticTag, error) {
	aats := &AccountAnalyticTags{}
	if err := c.SearchRead(AccountAnalyticTagModel, criteria, NewOptions().Limit(1), aats); err != nil {
		return nil, err
	}
	if aats != nil && len(*aats) > 0 {
		return &((*aats)[0]), nil
	}
	return nil, fmt.Errorf("account.analytic.tag was not found")
}

// FindAccountAnalyticTags finds account.analytic.tag records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAnalyticTags(criteria *Criteria, options *Options) (*AccountAnalyticTags, error) {
	aats := &AccountAnalyticTags{}
	if err := c.SearchRead(AccountAnalyticTagModel, criteria, options, aats); err != nil {
		return nil, err
	}
	return aats, nil
}

// FindAccountAnalyticTagIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAnalyticTagIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountAnalyticTagModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountAnalyticTagId finds record id by querying it with criteria.
func (c *Client) FindAccountAnalyticTagId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountAnalyticTagModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.analytic.tag was not found")
}
