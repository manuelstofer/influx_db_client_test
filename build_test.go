package influx_db_client_test

import (
	"testing"

	"github.com/influxdata/influxdb/client/v2"
)

func TestCanCompile(t *testing.T) {
	conf := client.HTTPConfig{}
	client.NewHTTPClient(conf)
}
