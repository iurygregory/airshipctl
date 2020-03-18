// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package remote

import (
	"github.com/spf13/cobra"

	"opendev.org/airship/airshipctl/pkg/environment"
	"opendev.org/airship/airshipctl/pkg/errors"
)

// NewRemoteCommand creates a new command that provides functionality to control remote entities.
func NewRemoteCommand(rootSettings *environment.AirshipCTLSettings) *cobra.Command {
	remoteRootCmd := &cobra.Command{
		Use:   "remote",
		Short: "Control remote entities, i.e. hosts.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.ErrNotImplemented{}
		},
	}

	return remoteRootCmd
}