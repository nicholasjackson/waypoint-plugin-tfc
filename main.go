package main

import (
	"github.com/nicholasjackson/waypoint-plugin-tfc/builder"
	"github.com/nicholasjackson/waypoint-plugin-tfc/platform"
	"github.com/nicholasjackson/waypoint-plugin-tfc/registry"
	"github.com/nicholasjackson/waypoint-plugin-tfc/release"
	sdk "github.com/hashicorp/waypoint-plugin-sdk"
)

func main() {
	// sdk.Main allows you to register the components which should
	// be included in your plugin
	// Main sets up all the go-plugin requirements

	sdk.Main(sdk.WithComponents(
		// Comment out any components which are not
		// required for your plugin
		&builder.Builder{},
		&registry.Registry{},
		&platform.Platform{},
		&release.ReleaseManager{},
	))
}
