package odoo_test

import (
	"context"
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	tc "github.com/testcontainers/testcontainers-go/modules/compose"

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

	time.Sleep(5 * time.Second)

	client, err := odoo.Connect(
		"http://localhost:8091",
		"xiatech_test",
		"admin",
		"admin",
	)
	assert.NoError(t, err)

	// Get a Product to add to our order
	_, err = client.GetProduct(12)
	assert.NoError(t, err)
}
