// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/plazafyi/plaza-cli/internal/mocktest"
	"github.com/plazafyi/plaza-cli/internal/requestflag"
)

func TestRoutingIsochrone(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"routing", "isochrone",
			"--lat", "0",
			"--lng", "0",
			"--time", "0",
			"--mode", "mode",
			"--output-fields", "output[fields]",
			"--output-geometry=true",
			"--output-include", "output[include]",
			"--output-precision", "0",
			"--output-simplify", "0",
		)
	})
}

func TestRoutingIsochronePost(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"routing", "isochrone-post",
			"--lat", "0",
			"--lng", "0",
			"--time", "0",
			"--mode", "mode",
			"--output-fields", "output[fields]",
			"--output-geometry=true",
			"--output-include", "output[include]",
			"--output-precision", "0",
			"--output-simplify", "0",
		)
	})
}

func TestRoutingMatrix(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"routing", "matrix",
			"--destination", "{lat: 48.8584, lng: 2.2945}",
			"--origin", "{lat: 48.8566, lng: 2.3522}",
			"--origin", "{lat: 48.8606, lng: 2.3376}",
			"--annotations", "annotations",
			"--fallback-speed", "1",
			"--mode", "auto",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(routingMatrix)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"routing", "matrix",
			"--destination.lat", "48.8584",
			"--destination.lng", "2.2945",
			"--origin.lat", "48.8566",
			"--origin.lng", "2.3522",
			"--origin.lat", "48.8606",
			"--origin.lng", "2.3376",
			"--annotations", "annotations",
			"--fallback-speed", "1",
			"--mode", "auto",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"destinations:\n" +
			"  - lat: 48.8584\n" +
			"    lng: 2.2945\n" +
			"origins:\n" +
			"  - lat: 48.8566\n" +
			"    lng: 2.3522\n" +
			"  - lat: 48.8606\n" +
			"    lng: 2.3376\n" +
			"annotations: annotations\n" +
			"fallback_speed: 1\n" +
			"mode: auto\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"routing", "matrix",
		)
	})
}

func TestRoutingNearest(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"routing", "nearest",
			"--lat", "0",
			"--lng", "0",
			"--output-fields", "output[fields]",
			"--output-include", "output[include]",
			"--output-precision", "0",
			"--radius", "0",
		)
	})
}

func TestRoutingNearestPost(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"routing", "nearest-post",
			"--lat", "0",
			"--lng", "0",
			"--output-fields", "output[fields]",
			"--output-include", "output[include]",
			"--output-precision", "0",
			"--radius", "0",
		)
	})
}

func TestRoutingRoute(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"routing", "route",
			"--destination", "{lat: 48.8584, lng: 2.2945}",
			"--origin", "{lat: 48.8566, lng: 2.3522}",
			"--alternatives", "0",
			"--annotations=true",
			"--depart-at", "'2019-12-27T18:11:19.117Z'",
			"--ev", "{battery_capacity_wh: 75000, connector_types: [string], initial_charge_pct: 0, min_charge_pct: 0, min_power_kw: 0}",
			"--exclude", "exclude",
			"--geometries", "geojson",
			"--mode", "auto",
			"--overview", "full",
			"--steps=true",
			"--traffic-model", "best_guess",
			"--waypoint", "[{lat: 48.8566, lng: 2.3522}]",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(routingRoute)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"routing", "route",
			"--destination.lat", "48.8584",
			"--destination.lng", "2.2945",
			"--origin.lat", "48.8566",
			"--origin.lng", "2.3522",
			"--alternatives", "0",
			"--annotations=true",
			"--depart-at", "'2019-12-27T18:11:19.117Z'",
			"--ev.battery-capacity-wh", "75000",
			"--ev.connector-types", "[string]",
			"--ev.initial-charge-pct", "0",
			"--ev.min-charge-pct", "0",
			"--ev.min-power-kw", "0",
			"--exclude", "exclude",
			"--geometries", "geojson",
			"--mode", "auto",
			"--overview", "full",
			"--steps=true",
			"--traffic-model", "best_guess",
			"--waypoint.lat", "48.8566",
			"--waypoint.lng", "2.3522",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"destination:\n" +
			"  lat: 48.8584\n" +
			"  lng: 2.2945\n" +
			"origin:\n" +
			"  lat: 48.8566\n" +
			"  lng: 2.3522\n" +
			"alternatives: 0\n" +
			"annotations: true\n" +
			"depart_at: '2019-12-27T18:11:19.117Z'\n" +
			"ev:\n" +
			"  battery_capacity_wh: 75000\n" +
			"  connector_types:\n" +
			"    - string\n" +
			"  initial_charge_pct: 0\n" +
			"  min_charge_pct: 0\n" +
			"  min_power_kw: 0\n" +
			"exclude: exclude\n" +
			"geometries: geojson\n" +
			"mode: auto\n" +
			"overview: full\n" +
			"steps: true\n" +
			"traffic_model: best_guess\n" +
			"waypoints:\n" +
			"  - lat: 48.8566\n" +
			"    lng: 2.3522\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"routing", "route",
		)
	})
}
