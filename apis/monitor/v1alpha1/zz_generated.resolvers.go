/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	v1alpha1 "github.com/crossplane-contrib/provider-upjet-digitalocean/apis/droplet/v1alpha1"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Alert.
func (mg *Alert) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var mrsp reference.MultiResolutionResponse
	var err error

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.Entities),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.EntitiesRefs,
		Selector:      mg.Spec.ForProvider.EntitiesSelector,
		To: reference.To{
			List:    &v1alpha1.DropletList{},
			Managed: &v1alpha1.Droplet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Entities")
	}
	mg.Spec.ForProvider.Entities = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.EntitiesRefs = mrsp.ResolvedReferences

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.Entities),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.InitProvider.EntitiesRefs,
		Selector:      mg.Spec.InitProvider.EntitiesSelector,
		To: reference.To{
			List:    &v1alpha1.DropletList{},
			Managed: &v1alpha1.Droplet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Entities")
	}
	mg.Spec.InitProvider.Entities = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.EntitiesRefs = mrsp.ResolvedReferences

	return nil
}
