//go:build integration

package odoo_test

import (
	"context"
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	tc "github.com/testcontainers/testcontainers-go/modules/compose"
	"github.com/xiatechs/XFuze/testutil"

	"github.com/xiatechs/go-odoo/v2"
)

func setup(t *testing.T) *odoo.Client {
	compose, err := tc.NewDockerCompose("testdata/docker-compose.yml")
	assert.NoError(t, err, "NewDockerComposeAPI()")

	t.Cleanup(func() {
		assert.NoError(t, compose.Down(
			context.Background(),
			tc.RemoveOrphans(true),
			tc.RemoveImagesLocal,
		), "compose.Down()")
	})

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	assert.NoError(t, compose.Up(ctx, tc.Wait(true)), "compose.Up()")

	// The Odoo instance can take a while to boot up, keep trying every 5 seconds for 30 minutes.
	var client *odoo.Client
	assert.NoError(t, testutil.WaitUntil(5*time.Second, 1800*time.Second, func() (bool, error) {
		client, err = odoo.Connect(
			"http://localhost:8091",
			"odoodb",
			"admin",
			"admin",
		)
		if err != nil {
			t.Log(err.Error())
			return false, nil
		}
		t.Log("successfully connected to Odoo")
		return true, nil
	}))

	return client
}

func Test_Odoo_Integration(t *testing.T) {
	client := setup(t)
	var createdProductID int

	t.Run("Create product within Odoo", func(t *testing.T) {
		// NOTE: Name, Type, Category ID, Tracking, Unit of Measure ID & Unit of Measure Product ID are mandatory params.
		id, err := client.CreateProduct(
			odoo.ProductTemplate{
				Name:          uuid.New().String(),
				Type:          "product",
				ResponsibleId: true,
				CategId:       8,
				UomId:         1,
				UomPoId:       1,
				Tracking:      "none",
			})
		createdProductID = id
		assert.NoError(t, err)
	})

	t.Run("Fetch product from Odoo", func(t *testing.T) {
		productBytes, err := client.GetProduct(createdProductID)
		assert.NoError(t, err)

		var product []odoo.Product
		assert.NoError(t, json.Unmarshal(productBytes, &product))
		assert.Len(t, product, 1)
		assert.Equal(t, createdProductID, product[0].ID)
	})

	t.Run("Fetch sales order from Odoo", func(t *testing.T) {
		salesOrderBytes, err := client.GetSalesOrder(0)
		assert.NoError(t, err)

		var salesOrders []odoo.SaleOrder
		assert.NoError(t, json.Unmarshal(salesOrderBytes, &salesOrders))
		assert.Len(t, salesOrders, 1)
		assert.Equal(t, 0, salesOrders[0].ID)
	})

	t.Run("Create sales order from JSON", func(t *testing.T) {
		orderJSON, err := os.ReadFile("testdata/order.json")
		assert.NoError(t, err)

		_, err = client.CreateSalesOrder(orderJSON)
		assert.NoError(t, err)
	})

	t.Run("Create customer", func(t *testing.T) {
		customer := odoo.Customer{Name: "Nick Pocock"}

		_, err := client.CreateCustomer(customer)
		assert.NoError(t, err)
	})

	t.Run("Fetch customer", func(t *testing.T) {
		customer := odoo.Customer{Name: "Nick Pocock"}

		id, err := client.CreateCustomer(customer)
		assert.NoError(t, err)

		_, err = client.GetCustomer(id)
		assert.NoError(t, err)
	})

	t.Run("Fetch customer from email", func(t *testing.T) {
		customer := odoo.Customer{Name: "Nick Pocock", Email: "test@test.com", Active: true}

		_, err := client.CreateCustomer(customer)
		assert.NoError(t, err)

		cust, err := client.GetCustomerFromEmail("test@test.com")
		require.NoError(t, err)
		assert.Equal(t, "Nick Pocock", cust.Name)
	})
}
