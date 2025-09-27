/*
Copyright 2025.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	backupv1alpha1 "github.com/0x4272616E646F6E/kubepbs/api/v1alpha1"
)

// nolint:unused
// log is for logging in this package.
var restorerequestlog = logf.Log.WithName("restorerequest-resource")

// SetupRestoreRequestWebhookWithManager registers the webhook for RestoreRequest in the manager.
func SetupRestoreRequestWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).For(&backupv1alpha1.RestoreRequest{}).
		WithValidator(&RestoreRequestCustomValidator{}).
		WithDefaulter(&RestoreRequestCustomDefaulter{}).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// +kubebuilder:webhook:path=/mutate-backup-pbs-k8s-io-v1alpha1-restorerequest,mutating=true,failurePolicy=fail,sideEffects=None,groups=backup.pbs.k8s.io,resources=restorerequests,verbs=create;update,versions=v1alpha1,name=mrestorerequest-v1alpha1.kb.io,admissionReviewVersions=v1

// RestoreRequestCustomDefaulter struct is responsible for setting default values on the custom resource of the
// Kind RestoreRequest when those are created or updated.
//
// NOTE: The +kubebuilder:object:generate=false marker prevents controller-gen from generating DeepCopy methods,
// as it is used only for temporary operations and does not need to be deeply copied.
type RestoreRequestCustomDefaulter struct {
	// TODO(user): Add more fields as needed for defaulting
}

var _ webhook.CustomDefaulter = &RestoreRequestCustomDefaulter{}

// Default implements webhook.CustomDefaulter so a webhook will be registered for the Kind RestoreRequest.
func (d *RestoreRequestCustomDefaulter) Default(_ context.Context, obj runtime.Object) error {
	restorerequest, ok := obj.(*backupv1alpha1.RestoreRequest)

	if !ok {
		return fmt.Errorf("expected an RestoreRequest object but got %T", obj)
	}
	restorerequestlog.Info("Defaulting for RestoreRequest", "name", restorerequest.GetName())

	// TODO(user): fill in your defaulting logic.

	return nil
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// NOTE: The 'path' attribute must follow a specific pattern and should not be modified directly here.
// Modifying the path for an invalid path can cause API server errors; failing to locate the webhook.
// +kubebuilder:webhook:path=/validate-backup-pbs-k8s-io-v1alpha1-restorerequest,mutating=false,failurePolicy=fail,sideEffects=None,groups=backup.pbs.k8s.io,resources=restorerequests,verbs=create;update,versions=v1alpha1,name=vrestorerequest-v1alpha1.kb.io,admissionReviewVersions=v1

// RestoreRequestCustomValidator struct is responsible for validating the RestoreRequest resource
// when it is created, updated, or deleted.
//
// NOTE: The +kubebuilder:object:generate=false marker prevents controller-gen from generating DeepCopy methods,
// as this struct is used only for temporary operations and does not need to be deeply copied.
type RestoreRequestCustomValidator struct {
	// TODO(user): Add more fields as needed for validation
}

var _ webhook.CustomValidator = &RestoreRequestCustomValidator{}

// ValidateCreate implements webhook.CustomValidator so a webhook will be registered for the type RestoreRequest.
func (v *RestoreRequestCustomValidator) ValidateCreate(_ context.Context, obj runtime.Object) (admission.Warnings, error) {
	restorerequest, ok := obj.(*backupv1alpha1.RestoreRequest)
	if !ok {
		return nil, fmt.Errorf("expected a RestoreRequest object but got %T", obj)
	}
	restorerequestlog.Info("Validation for RestoreRequest upon creation", "name", restorerequest.GetName())

	// TODO(user): fill in your validation logic upon object creation.

	return nil, nil
}

// ValidateUpdate implements webhook.CustomValidator so a webhook will be registered for the type RestoreRequest.
func (v *RestoreRequestCustomValidator) ValidateUpdate(_ context.Context, oldObj, newObj runtime.Object) (admission.Warnings, error) {
	restorerequest, ok := newObj.(*backupv1alpha1.RestoreRequest)
	if !ok {
		return nil, fmt.Errorf("expected a RestoreRequest object for the newObj but got %T", newObj)
	}
	restorerequestlog.Info("Validation for RestoreRequest upon update", "name", restorerequest.GetName())

	// TODO(user): fill in your validation logic upon object update.

	return nil, nil
}

// ValidateDelete implements webhook.CustomValidator so a webhook will be registered for the type RestoreRequest.
func (v *RestoreRequestCustomValidator) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	restorerequest, ok := obj.(*backupv1alpha1.RestoreRequest)
	if !ok {
		return nil, fmt.Errorf("expected a RestoreRequest object but got %T", obj)
	}
	restorerequestlog.Info("Validation for RestoreRequest upon deletion", "name", restorerequest.GetName())

	// TODO(user): fill in your validation logic upon object deletion.

	return nil, nil
}
