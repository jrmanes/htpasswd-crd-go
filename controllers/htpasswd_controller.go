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
	"os"

	"github.com/go-logr/logr"
	securityv1 "github.com/jrmanes/htpasswd-crd-go/api/v1"
	"k8s.io/apimachinery/pkg/api/errors"
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
		logWithValues.Info("unable to fetch htpasswd", "name", req.Namespace)
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	var htpasswdList securityv1.HtpasswdList
	if err := r.List(ctx, &htpasswdList, client.InNamespace(req.Namespace)); err != nil {
		logWithValues.Info("unable to list child Jobs", "name", req.Namespace)
		return ctrl.Result{}, err
	}

	// create the file for user and password
	user := htpasswd.Spec.User
	pass := htpasswd.Spec.Password
	// create the htpasswd format
	r.GenerateHtpasswd(user, pass)

	// format the new secret
	secret := r.CreateSecret(htpasswd)
	// call the API in order to creat a new secret and check if there is any error
	err := r.Create(context.TODO(), secret)
	switch {
	case err == nil:
		logWithValues.Info("Created Htpasswd resource")
	case errors.IsAlreadyExists(err):
		logWithValues.Info("Htpasswd resource already exists")
	default:
		logWithValues.Error(err, "Failed to create Htpasswd resource")
		os.Exit(1)
	}

	//if htpasswd.Status.Status == "" {
	//	htpasswd.Status.Status = "HOWDY"
	//}
	//if err := r.Status().Update(ctx, &htpasswd); err != nil {
	//	logWithValues.Info("ERROR updating the STATUS", "name", req.Namespace)
	//	logWithValues.Error(err, "ERROR updating the STATUS", "name", req.Namespace)
	//	return ctrl.Result{}, client.IgnoreNotFound(err)
	//}
	return ctrl.Result{}, nil
}

func (r *HtpasswdReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&securityv1.Htpasswd{}).
		Complete(r)
}

// TODO
// GenerateHtpasswd will encrypt the user and password as htpasswd does
func (r *HtpasswdReconciler) GenerateHtpasswd(user, pass string) {
	log.Printf("///// GenerateHtpasswd ///////////")
	log.Printf("user %s and pass %s =>\n", user, pass)
}

// NewSecret prepare the new secret using our htpasswd content
//func (r *HtpasswdReconciler) CreateSecret(ht securityv1.Htpasswd) *corev1.Secret {
func (r *HtpasswdReconciler) CreateSecret(ht securityv1.Htpasswd) *securityv1.Htpasswd {
	log.Println("Lets create a new secret...", ht.Name)
	labels := map[string]string{
		"app": ht.Name,
	}

	return &securityv1.Htpasswd{
		TypeMeta: metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{
			Name:      ht.Name,
			Namespace: ht.Namespace,
			Labels:    labels,
		},
		Spec:   securityv1.HtpasswdSpec{},
		Status: securityv1.HtpasswdStatus{},
	}
}
