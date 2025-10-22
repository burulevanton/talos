// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package vm

import (
	"github.com/siderolabs/talos/pkg/provision"
)

func (p *Provisioner) startDNSd(_ *State, _ provision.ClusterRequest) error {
	return nil
}

func (p *Provisioner) stopDNSd(_ *State) error {
	return nil
}
