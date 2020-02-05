/*
Copyright (c) 2019 the Octant contributors. All Rights Reserved.
SPDX-License-Identifier: Apache-2.0
*/

package printer

import (
	"context"

	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"

	"github.com/vmware-tanzu/octant/pkg/view/component"
)

// PersistentVolumeListHandler is a printFunc that prints persistentvolumes
func PersistentVolumeListHandler(ctx context.Context, list *corev1.PersistentVolumeList, options Options) (component.Component, error) {
	if list == nil {
		return nil, errors.New("nil list")
	}

	cols := component.NewTableCols("Name", "Status", "Volume", "Capacity", "Access Modes", "Storage Class", "Age")
	tbl := component.NewTable("Persistent Volume s",
		"We couldn't find any persistent volume s!", cols)

	return tbl, nil
}

// PersistentVolumeHandler is a printFunc that prints a PersistentVolume
func PersistentVolumeHandler(ctx context.Context, persistentVolume *corev1.PersistentVolume, options Options) (component.Component, error) {
	o := NewObject(persistentVolume)
	o.EnableEvents()

	return o.ToComponent(ctx, options)
}

// PersistentVolumeConfiguration generates a persistenvolume configuration
type PersistentVolumeConfiguration struct {
	persistentVolume *corev1.PersistentVolume
}

// NewPersistentVolumeConfiguration creates an instance of PersistentVolumeConfiguration
func NewPersistentVolumeConfiguration(pv *corev1.PersistentVolume) *PersistentVolumeConfiguration {
	return &PersistentVolumeConfiguration{
		persistentVolume: pv,
	}
}
