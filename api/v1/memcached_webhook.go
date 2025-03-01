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

package v1

import (
	"errors"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	rplog "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// log is for logging in this package.
var log = rplog.Log.WithName("memcached-resource.webhook")

// SetupWebhookWithManager will setup the manager to manage the webhooks
func (r *Memcached) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// +kubebuilder:webhook:path=/mutate-cache-devspace-com-v1-memcached,mutating=true,failurePolicy=fail,sideEffects=None,groups=cache.devspace.com,resources=memcacheds,verbs=create;update,versions=v1,name=mmemcached.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &Memcached{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *Memcached) Default() {
	log.Info("default", "name", r.Name, "namespace", r.Namespace)

	if r.Spec.Size == 0 {
		log.Info("cr does not have size, defaulting to 3", "name", r.Name, "namespace", r.Namespace)
		r.Spec.Size = 3
	}
}

// +kubebuilder:webhook:path=/validate-cache-devspace-com-v1-memcached,mutating=false,failurePolicy=fail,sideEffects=None,groups=cache.devspace.com,resources=memcacheds,verbs=create;update,versions=v1,name=vmemcached.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &Memcached{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *Memcached) ValidateCreate() (admission.Warnings, error) {
	log.Info("validate create", "name", r.Name, "namespace", r.Namespace)

	return nil, r.validateOdd(r.Spec.Size)
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *Memcached) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	log.Info("validate update", "name", r.Name, "namespace", r.Namespace)

	return nil, r.validateOdd(r.Spec.Size)
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *Memcached) ValidateDelete() (admission.Warnings, error) {
	log.Info("validate delete", "name", r.Name, "namespace", r.Namespace)

	return nil, nil
}

func (r *Memcached) validateOdd(n int) error {
	if n%2 == 0 {
		log.Error(errors.New("cluster size must be an odd number"), "validation failed", "name", r.Name, "namespace", r.Namespace)
		return errors.New("cluster size must be an odd number")
	}
	return nil
}
