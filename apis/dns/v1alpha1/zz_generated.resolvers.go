/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Record.
func (mg *Record) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Domain),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.DomainRef,
		Selector:     mg.Spec.ForProvider.DomainSelector,
		To: reference.To{
			List:    &DomainList{},
			Managed: &Domain{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Domain")
	}
	mg.Spec.ForProvider.Domain = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DomainRef = rsp.ResolvedReference

	return nil
}
