// Copyright © 2017 The Qaclana Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fwserver

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"gitlab.com/qaclana/qaclana/pkg/fwserver/server"
)

// NewStartFirewallServerCommand initializes a command that can be used to start the firewall server
func NewStartFirewallServerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Starts a new firewall server",
		Long:  "Starts a new firewall server",
		Run: func(cmd *cobra.Command, args []string) {
			start(cmd, args)
		},
	}

	cmd.Flags().IntP("port", "p", 8000, "The port to bind the main server to")
	cmd.Flags().IntP("healthcheck-port", "", 9000, "The port to bind the healthcheck server to")
	cmd.Flags().IntP("grpc-port", "", 10000, "The port to bind the gRPC interface to")
	viper.BindPFlag("port", cmd.Flags().Lookup("port"))
	viper.BindPFlag("healthcheck-port", cmd.Flags().Lookup("healthcheck-port"))
	viper.BindPFlag("grpc-port", cmd.Flags().Lookup("grpc-port"))

	cmd.Flags().StringP("backend-hostname", "u", "qaclana-backend.qaclana-infra.svc.cluster.local", "The hostname to the Qaclana Backend")
	cmd.Flags().IntP("backend-grpc-port", "", 10000, "The gRPC port for the Qaclana Backend")
	viper.BindPFlag("backend-hostname", cmd.Flags().Lookup("backend-hostname"))
	viper.BindPFlag("backend-grpc-port", cmd.Flags().Lookup("backend-grpc-port"))

	return cmd
}

func start(cmd *cobra.Command, args []string) {
	server.Start()
}