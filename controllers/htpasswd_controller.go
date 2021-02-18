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
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	htpasswdv1 "github.com/jrmanes/htpasswd-crd-go/api/v1"
)

// HtpasswdReconciler reconciles a Htpasswd object
type HtpasswdReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=htpasswd.htpasswd-crd-go,resources=htpasswds,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=htpasswd.htpasswd-crd-go,resources=htpasswds/status,verbs=get;update;patch

func (r *HtpasswdReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("htpasswd", req.NamespacedName)

	// your logic here
	log.Println("THIS IS THE HTPASSWD")
	log.Println(req.Name)
	log.Println(req.Namespace)
	return ctrl.Result{}, nil
}

func (r *HtpasswdReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&htpasswdv1.Htpasswd{}).
		Complete(r)
}
