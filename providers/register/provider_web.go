/*
 *  *******************************************************************************
 *  * Copyright (c) 2019 Edgeworx, Inc.
 *  *
 *  * This program and the accompanying materials are made available under the
 *  * terms of the Eclipse Public License v. 2.0 which is available at
 *  * http://www.eclipse.org/legal/epl-2.0
 *  *
 *  * SPDX-License-Identifier: EPL-2.0
 *  *******************************************************************************
 *
 */

package register

import (
	"github.com/iofog/iofog-kubelet/providers"
	"github.com/iofog/iofog-kubelet/providers/web"
)

func init() {
	register("web", initWeb)
}

func initWeb(cfg InitConfig) (providers.Provider, error) {
	return web.NewBrokerProvider(
		cfg.DaemonPort,
		cfg.NodeName,
		cfg.OperatingSystem,
		cfg.ControllerToken,
		cfg.ControllerUrl,
		cfg.NodeId)
}
