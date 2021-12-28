/*
Copyright 2021.

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
	"reflect"
	"time"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	ocv1 "github.com/burnyd/k8sNetworkgNMIOperator/api/v1"
	k8sgnmi "github.com/burnyd/k8sNetworkgNMIOperator/pkg/k8sgnmi"
)

// OcswitchesReconciler reconciles a Ocswitches object
type OcswitchesReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=oc.ocoperator.com,resources=ocswitches,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=oc.ocoperator.com,resources=ocswitches/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=oc.ocoperator.com,resources=ocswitches/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Ocswitches object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.10.0/pkg/reconcile
func (r *OcswitchesReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx, "Openconfig switches", req.Namespace)
	log.Log.Info("Starting to check on the device")
	oC := &ocv1.Ocswitches{}
	err := r.Client.Get(ctx, req.NamespacedName, oC)
	if err != nil {
		if errors.IsNotFound(err) {
			log.Log.Info("HelloApp resource not found. Ignoring since object must be deleted")
			return ctrl.Result{}, nil
		}
		log.Log.Error(err, "Failed to get OcSwitches controller")
		return ctrl.Result{}, err
	}
	// Start the logic if the bgp list has anything within it.
	if len(oC.Spec.Bgp) > 0 {
		log.Log.Info("Bgp enabled is within the Kubernetes spec")
		// start the portion to connect to the bgp neighbor path
		log.Log.Info("Checking on Device " + oC.Spec.Host + ":" + oC.Spec.Port)
		gnmi := k8sgnmi.GNMI_CFG{
			Username: oC.Spec.Username,
			Password: oC.Spec.Password,
			Host:     oC.Spec.Host + ":" + oC.Spec.Port,
			Path:     "network-instances/network-instance[name=default]/protocols/protocol[name=BGP]/bgp/neighbors/",
			Origin:   "openconfig",
			OcModel:  "bgp",
		}
		// Connect to the device.
		gnmi.Connect()
		// Run the specific get neighbors method to return the data the data as a map[string]int
		GetNeighbors := gnmi.ReturnBgp()
		// Create a map that will be for the neighbors that are within the switch
		gNMINeighborMap := make(map[string]uint32)
		log.Log.Info("Listing neighbors on the device")
		// Range through the neihbors and apend to the map
		for key, value := range GetNeighbors {
			log.Log.Info("Neighbor ", key, value)
			gNMINeighborMap[key] = value
		}
		// Create a map for what the k8s spec is.
		k8sNeighbormap := make(map[string]uint32)

		for _, neighbor := range oC.Spec.Bgp {
			k8sNeighbormap[neighbor.NeighborIP] = neighbor.RemoteAs
		}
		// Compare the neighbors with reflect.
		if reflect.DeepEqual(k8sNeighbormap, gNMINeighborMap) == false {
			//fmt.Println("Need to retify here")
			gnmi.SetBgpYgot(oC.Spec.BgpAS, k8sNeighbormap)
		} else {
			log.Log.Info("Neighbors match k8s spec to what is on the switch via gnmi")
		}

	}

	return ctrl.Result{RequeueAfter: time.Second * 5}, nil

	//return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *OcswitchesReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&ocv1.Ocswitches{}).
		Complete(r)
}
