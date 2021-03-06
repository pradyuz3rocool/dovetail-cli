/*
* Copyright © 2018. TIBCO Software Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
 */

// Package commands is the one containing all the cli commands
package commands

import (
	"fmt"

	"github.com/TIBCOSoftware/dovetail-cli/version"
	"github.com/spf13/cobra"
)

var (
	// VersionCmd prints out the current cli version
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the app version",
		Run:   printVersion,
	}
)

// GetVersion return version of CLI/node and commit hash
func GetVersion() string {
	v := version.Version
	if version.GitCommit != "" {
		v = v + "-" + version.GitCommit
	}
	return v
}

// printVersion prints the version
func printVersion(cmd *cobra.Command, args []string) {
	v := GetVersion()
	fmt.Println(v)
}
