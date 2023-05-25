/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	contactprofile "github.com/upbound/provider-azure/internal/controller/orbital/contactprofile"
	spacecraft "github.com/upbound/provider-azure/internal/controller/orbital/spacecraft"
)

// Setup_orbital creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_orbital(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		contactprofile.Setup,
		spacecraft.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}