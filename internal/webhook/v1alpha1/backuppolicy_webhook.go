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
var backuppolicylog = logf.Log.WithName("backuppolicy-resource")

// SetupBackupPolicyWebhookWithManager registers the webhook for BackupPolicy in the manager.
func SetupBackupPolicyWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).For(&backupv1alpha1.BackupPolicy{}).
		WithValidator(&BackupPolicyCustomValidator{}).
		WithDefaulter(&BackupPolicyCustomDefaulter{}).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// +kubebuilder:webhook:path=/mutate-backup-pbs-k8s-io-v1alpha1-backuppolicy,mutating=true,failurePolicy=fail,sideEffects=None,groups=backup.pbs.k8s.io,resources=backuppolicies,verbs=create;update,versions=v1alpha1,name=mbackuppolicy-v1alpha1.kb.io,admissionReviewVersions=v1

// BackupPolicyCustomDefaulter struct is responsible for setting default values on the custom resource of the
// Kind BackupPolicy when those are created or updated.
//
// NOTE: The +kubebuilder:object:generate=false marker prevents controller-gen from generating DeepCopy methods,
// as it is used only for temporary operations and does not need to be deeply copied.
type BackupPolicyCustomDefaulter struct {
	// TODO(user): Add more fields as needed for defaulting
}

var _ webhook.CustomDefaulter = &BackupPolicyCustomDefaulter{}

// Default implements webhook.CustomDefaulter so a webhook will be registered for the Kind BackupPolicy.
func (d *BackupPolicyCustomDefaulter) Default(_ context.Context, obj runtime.Object) error {
	backuppolicy, ok := obj.(*backupv1alpha1.BackupPolicy)

	if !ok {
		return fmt.Errorf("expected an BackupPolicy object but got %T", obj)
	}
	backuppolicylog.Info("Defaulting for BackupPolicy", "name", backuppolicy.GetName())

	// TODO(user): fill in your defaulting logic.

	return nil
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// NOTE: The 'path' attribute must follow a specific pattern and should not be modified directly here.
// Modifying the path for an invalid path can cause API server errors; failing to locate the webhook.
// +kubebuilder:webhook:path=/validate-backup-pbs-k8s-io-v1alpha1-backuppolicy,mutating=false,failurePolicy=fail,sideEffects=None,groups=backup.pbs.k8s.io,resources=backuppolicies,verbs=create;update,versions=v1alpha1,name=vbackuppolicy-v1alpha1.kb.io,admissionReviewVersions=v1

// BackupPolicyCustomValidator struct is responsible for validating the BackupPolicy resource
// when it is created, updated, or deleted.
//
// NOTE: The +kubebuilder:object:generate=false marker prevents controller-gen from generating DeepCopy methods,
// as this struct is used only for temporary operations and does not need to be deeply copied.
type BackupPolicyCustomValidator struct {
	// TODO(user): Add more fields as needed for validation
}

var _ webhook.CustomValidator = &BackupPolicyCustomValidator{}

// ValidateCreate implements webhook.CustomValidator so a webhook will be registered for the type BackupPolicy.
func (v *BackupPolicyCustomValidator) ValidateCreate(_ context.Context, obj runtime.Object) (admission.Warnings, error) {
	backuppolicy, ok := obj.(*backupv1alpha1.BackupPolicy)
	if !ok {
		return nil, fmt.Errorf("expected a BackupPolicy object but got %T", obj)
	}
	backuppolicylog.Info("Validation for BackupPolicy upon creation", "name", backuppolicy.GetName())

	// TODO(user): fill in your validation logic upon object creation.

	return nil, nil
}

// ValidateUpdate implements webhook.CustomValidator so a webhook will be registered for the type BackupPolicy.
func (v *BackupPolicyCustomValidator) ValidateUpdate(_ context.Context, oldObj, newObj runtime.Object) (admission.Warnings, error) {
	backuppolicy, ok := newObj.(*backupv1alpha1.BackupPolicy)
	if !ok {
		return nil, fmt.Errorf("expected a BackupPolicy object for the newObj but got %T", newObj)
	}
	backuppolicylog.Info("Validation for BackupPolicy upon update", "name", backuppolicy.GetName())

	// TODO(user): fill in your validation logic upon object update.

	return nil, nil
}

// ValidateDelete implements webhook.CustomValidator so a webhook will be registered for the type BackupPolicy.
func (v *BackupPolicyCustomValidator) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	backuppolicy, ok := obj.(*backupv1alpha1.BackupPolicy)
	if !ok {
		return nil, fmt.Errorf("expected a BackupPolicy object but got %T", obj)
	}
	backuppolicylog.Info("Validation for BackupPolicy upon deletion", "name", backuppolicy.GetName())

	// TODO(user): fill in your validation logic upon object deletion.

	return nil, nil
}
