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
			"--geometry", "{coordinates: [2.3522, 48.8566], type: Point}",
			"--time", "1",
			"--format", "format",
			"--mode", "auto",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(routingIsochrone)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"routing", "isochrone",
			"--geometry.coordinates", "[2.3522, 48.8566]",
			"--geometry.type", "Point",
			"--time", "1",
			"--format", "format",
			"--mode", "auto",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"geometry:\n" +
			"  coordinates:\n" +
			"    - 2.3522\n" +
			"    - 48.8566\n" +
			"  type: Point\n" +
			"time:\n" +
			"  - 1\n" +
			"mode: auto\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"routing", "isochrone",
			"--format", "format",
		)
	})
}

func TestRoutingMatrix(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"routing", "matrix",
			"--destination", "{coordinates: [2.2945, 48.8584], type: Point}",
			"--origin", "{coordinates: [2.3522, 48.8566], type: Point}",
			"--origin", "{coordinates: [2.3376, 48.8606], type: Point}",
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
			"--destination.coordinates", "[2.2945, 48.8584]",
			"--destination.type", "Point",
			"--origin.coordinates", "[2.3522, 48.8566]",
			"--origin.type", "Point",
			"--origin.coordinates", "[2.3376, 48.8606]",
			"--origin.type", "Point",
			"--annotations", "annotations",
			"--fallback-speed", "1",
			"--mode", "auto",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"destinations:\n" +
			"  - coordinates:\n" +
			"      - 2.2945\n" +
			"      - 48.8584\n" +
			"    type: Point\n" +
			"origins:\n" +
			"  - coordinates:\n" +
			"      - 2.3522\n" +
			"      - 48.8566\n" +
			"    type: Point\n" +
			"  - coordinates:\n" +
			"      - 2.3376\n" +
			"      - 48.8606\n" +
			"    type: Point\n" +
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
			"--geometry", "{coordinates: [2.3522, 48.8566], type: Point}",
			"--radius", "1",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(routingNearest)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"routing", "nearest",
			"--geometry.coordinates", "[2.3522, 48.8566]",
			"--geometry.type", "Point",
			"--radius", "1",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"geometry:\n" +
			"  coordinates:\n" +
			"    - 2.3522\n" +
			"    - 48.8566\n" +
			"  type: Point\n" +
			"radius: 1\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"routing", "nearest",
		)
	})
}

func TestRoutingRoute(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"routing", "route",
			"--destination", "{coordinates: [2.2945, 48.8584], type: Point}",
			"--origin", "{coordinates: [2.3522, 48.8566], type: Point}",
			"--format", "format",
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
			"--waypoint", "[{coordinates: [2.3522, 48.8566], type: Point}]",
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
			"--destination.coordinates", "[2.2945, 48.8584]",
			"--destination.type", "Point",
			"--origin.coordinates", "[2.3522, 48.8566]",
			"--origin.type", "Point",
			"--format", "format",
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
			"--waypoint.coordinates", "[2.3522, 48.8566]",
			"--waypoint.type", "Point",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"destination:\n" +
			"  coordinates:\n" +
			"    - 2.2945\n" +
			"    - 48.8584\n" +
			"  type: Point\n" +
			"origin:\n" +
			"  coordinates:\n" +
			"    - 2.3522\n" +
			"    - 48.8566\n" +
			"  type: Point\n" +
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
			"  - coordinates:\n" +
			"      - 2.3522\n" +
			"      - 48.8566\n" +
			"    type: Point\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"routing", "route",
			"--format", "format",
		)
	})
}
