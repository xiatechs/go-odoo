package odoo_test

import (
	"context"
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	tc "github.com/testcontainers/testcontainers-go/modules/compose"
	"github.com/xiatechs/XFuze/testutil"

	"github.com/xiatechs/go-odoo/v2"
)

func TestClient_Unmarshal_Products(t *testing.T) {
	// products.json is what we get after we fetch data from Odoo and serialize it as JSON.
	productsJSON, err := os.ReadFile("testdata/products.json")
	assert.NoError(t, err)

	var products []odoo.Product
	err = json.Unmarshal(productsJSON, &products)
	assert.NoError(t, err)
}

func TestClient_CreateProduct(t *testing.T) {
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

	var client *odoo.Client
	assert.NoError(t, testutil.WaitUntil(5*time.Second, 1000*time.Second, func() (bool, error) {
		client, err = odoo.Connect(
			"http://localhost:8091",
			"xiatech_test",
			"admin",
			"admin",
		)
		if err != nil {
			t.Log(err.Error())
			return false, nil
		}
		return true, nil
	}))

	// NOTE: Name and Type are mandatory params.
	_, err = client.CreateProduct(odoo.ProductTemplate{Name: "test", Type: "product"})
	assert.NoError(t, err)
}
