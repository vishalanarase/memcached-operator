package controller

import (
	"context"
	"errors"
	"fmt"

	cachev1 "github.com/vishalanarase/memcached-operator/api/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// RemoveFinalizer removes the finalizer from the given custom resource.
func (r *MemcachedReconciler) RemoveFinalizer(ctx context.Context, memcached *cachev1.Memcached) (ctrl.Result, error) {
	log := log.FromContext(ctx)
	meta.SetStatusCondition(&memcached.Status.Conditions, metav1.Condition{
		Type:               cachev1.ConditionDowngraded,
		Status:             metav1.ConditionUnknown,
		Reason:             "Finalizing",
		Message:            fmt.Sprintf("Performing finalizer operations before delete"),
		ObservedGeneration: memcached.Generation,
	})
	if err := r.Status().Update(ctx, memcached); err != nil {
		log.Error(err, "Failed to update Memcached status")
		return ctrl.Result{}, err
	}

	if err := r.finalizeMemcached(ctx, memcached); err != nil {
		log.Error(err, "Failed to finalize Memcached")
		return ctrl.Result{}, err
	}

	if err := r.Get(ctx, types.NamespacedName{Name: memcached.Name, Namespace: memcached.Namespace}, memcached); err != nil {
		log.Error(err, "Failed to re-fetch memcached")
		return ctrl.Result{}, err
	}

	meta.SetStatusCondition(&memcached.Status.Conditions, metav1.Condition{
		Type:               cachev1.ConditionDegraded,
		Status:             metav1.ConditionTrue,
		Reason:             "Finalized",
		Message:            fmt.Sprintf("Finalizer operations complete"),
		ObservedGeneration: memcached.Generation,
	})
	if err := r.Status().Update(ctx, memcached); err != nil {
		log.Error(err, "Failed to update Memcached status")
		return ctrl.Result{}, err
	}

	log.Info("Removing finalizer")
	if ok := controllerutil.RemoveFinalizer(memcached, cachev1.MemcachedFinalizer); !ok {
		log.Error(errors.New("RemoveFinalizer failed"), "Failed to remove finalizer for Memcached, requeuing...")
		return ctrl.Result{Requeue: true}, nil
	}

	if err := r.Update(ctx, memcached); err != nil {
		log.Error(err, "Failed to remove finalizer for Memcached")
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// finalizeMemcached will perform the required operations before delete from the cluster.
func (r *MemcachedReconciler) finalizeMemcached(ctx context.Context, cr *cachev1.Memcached) error {
	log := log.FromContext(ctx)

	r.Recorder.Event(cr, "Warning", "Deleting",
		fmt.Sprintf("Custom Resource %s is being deleted from the namespace %s",
			cr.Name,
			cr.Namespace))
	log.Info("Successfully finalized memcached")
	return nil
}
