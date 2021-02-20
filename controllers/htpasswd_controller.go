/*


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

package controllers

import (
	"context"
	"log"

	"github.com/go-logr/logr"
	securityv1 "github.com/jrmanes/htpasswd-crd-go/api/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// HtpasswdReconciler reconciles a Htpasswd object
type HtpasswdReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=security.htpasswd-crd-go,resources=htpasswds,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=security.htpasswd-crd-go,resources=htpasswds/status,verbs=get;update;patch

// Reconcile is where controller logic lives
func (r *HtpasswdReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	logWithValues := r.Log.WithValues("htpasswd", req.NamespacedName)

	// Retrieve the htpasswd data from manifest
	var htpasswd securityv1.Htpasswd
	if err := r.Get(ctx, req.NamespacedName, &htpasswd); err != nil {
		logWithValues.Error(err, "unable to fetch htpasswd")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	log.Println("manifest htpasswd.Name =>", htpasswd.Name)
	log.Println("manifest htpasswd.Namespace =>", htpasswd.Namespace)
	log.Println("manifest htpasswd.Spec =>", htpasswd.Spec)

	//existingSecrets := &corev1.Secret{}
	//err := r.Get(ctx, req.NamespacedName, existingSecrets)
	//if err != nil {
	//	if errors.IsNotFound(err) {
	//		// Request object not found, could have been deleted after reconcile request.
	//		// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
	//		// Return and don't requeue
	//		log.Println("NOT FOUND ANY SECRETS", existingSecrets)
	//		return reconcile.Result{}, nil
	//	}
	//	// Error reading the object - requeue the request.
	//	return reconcile.Result{}, err
	//}

	//log.Println("###########")
	//log.Println(existingSecrets.Name)
	//log.Println("###########")

	// format the new secret
	secret := r.NewSecret(htpasswd)
	log.Println("secret", secret)

	// call the API in order to creat a new secret
	err := r.Create(context.TODO(), secret)
	if err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *HtpasswdReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&securityv1.Htpasswd{}).
		Complete(r)
}

// NewSecret prepare the new secret using our htpasswd content
func (r *HtpasswdReconciler) NewSecret(ht securityv1.Htpasswd) *corev1.Secret {
	labels := map[string]string{
		"app": ht.Name,
	}

	return &corev1.Secret{
		TypeMeta: metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{
			Name:      ht.Name,
			Namespace: ht.Namespace,
			Labels:    labels,
		},
		Data:       nil,
		StringData: nil,
		Type:       "Opaque",
	}
}
