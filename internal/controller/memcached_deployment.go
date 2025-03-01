package controller

import (
	"context"
	"fmt"

	cachev1 "github.com/vishalanarase/memcached-operator/api/v1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// CreateDeployment creates a new Deployment for a Memcached resource. It also sets
// the appropriate OwnerReferences on the resource so handleObject can discover
// the Memcached resource that 'owns' it.
func (r *MemcachedReconciler) CreateDeployment(ctx context.Context, memcached *cachev1.Memcached) error {
	log := log.FromContext(ctx)
	deployment := r.deploymentForMemcached(memcached)

	// Set Memcached instance as the owner and controller
	if err := controllerutil.SetControllerReference(memcached, deployment, r.Scheme); err != nil {
		log.Error(err, "Failed to set controller reference for Deployment")
		return err
	}

	log.Info("Creating a new Deployment",
		"Deployment.Namespace", deployment.Namespace, "Deployment.Name", deployment.Name)
	err := r.Create(ctx, deployment)
	if err != nil {
		log.Error(err, "Failed to create new Deployment",
			"Deployment.Namespace", deployment.Namespace, "Deployment.Name", deployment.Name)
		return err
	}

	return nil
}

// deploymentForMemcached returns a memcached Deployment object
func (r *MemcachedReconciler) deploymentForMemcached(m *cachev1.Memcached) *appsv1.Deployment {
	ls := labelsForMemcached(m.Name)
	replicas := m.Spec.Size
	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      m.Name,
			Namespace: m.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: ls,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: ls,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Image:   "memcached:1.4.36-alpine",
						Name:    "memcached",
						Command: []string{"memcached", "-m=64", "-o", "modern", "-v"},
						Ports: []corev1.ContainerPort{{
							ContainerPort: 11211,
							Name:          "memcached",
						}},
					}},
				},
			},
		},
	}
}

func labelsForMemcached(name string) map[string]string {
	return map[string]string{"app": "memcached", "memcached_cr": name}
}

//updateDeployment updates the deployment with the given Memcached object

func (r *MemcachedReconciler) UpdateDeployment(ctx context.Context, memcached *cachev1.Memcached, found *appsv1.Deployment) (ctrl.Result, error) {
	log := log.FromContext(ctx)
	size := memcached.Spec.Size
	found.Spec.Replicas = &size
	if err := r.Update(ctx, found); err != nil {
		log.Error(err, "Failed to update Deployment",
			"Deployment.Namespace", found.Namespace, "Deployment.Name", found.Name)

		// Get cr
		err = r.Get(ctx, types.NamespacedName{Name: memcached.Name, Namespace: memcached.Namespace}, memcached)
		if err != nil {
			log.Error(err, "Failed to get Memcached")
			return ctrl.Result{}, err
		}

		// The following implementation will update the status
		meta.SetStatusCondition(&memcached.Status.Conditions, metav1.Condition{
			Type:    cachev1.ConditionAvailable,
			Status:  metav1.ConditionFalse,
			Reason:  "Resizing",
			Message: fmt.Sprintf("Failed to update the size for the custom resource (%s): (%s)", memcached.Name, err),
		})

		if err := r.Status().Update(ctx, memcached); err != nil {
			log.Error(err, "Failed to update Memcached status")
			return ctrl.Result{}, err
		}

		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}
