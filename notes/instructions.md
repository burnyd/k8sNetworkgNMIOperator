## Downgrade to go 1.16
## Scaffold some of the operator. So everything will be operator.com/v1 for example in the api.
operator-sdk init --skip-go-version-check --domain ocoperator.com --repo github.com/burnyd/k8sNetworkgNMIOperator --verbose

## Create the struct data and such.

operator-sdk create api --group oc --version v1 --kind Ocswitches --resource --controller

## Edit the actual structs and what they should be within the api portion
api/v1/oc_switches_types.go

^^ This will convert this data to an api within the crd that will later be applies to the k8s cluster.

## Make install will then turn the go structs into openapiv3

➜  ocopator make installer
/home/burnyd/projects/ocoperator/bin/controller-gen rbac:roleName=manager-role crd webhook paths="./..." output:crd:artifacts:config=config/crd/bases
go: creating new go.mod: module tmp
Downloading sigs.k8s.io/kustomize/kustomize/v3@v3.8.7
go: downloading github.com/emicklei/go-restful v1.1.3
go: downloading github.com/dgrijalva/jwt-go v1.0.2
go get: added sigs.k8s.io/kustomize/kustomize/v3 v3.8.7
/home/burnyd/projects/ocoperator/bin/kustomize build config/crd | kubectl apply -f -
customresourcedefinition.apiextensions.k8s.io/ocswitches.oc.ocoperator.com created
➜  ocoperator

## Checkout the crd
config/crd/bases/oc.operator.com_ocswitches.yaml

## Apply the manifest.
kubectl apply -f config/samples/ceos1.yaml

## Place holder for the ygoy portion.
### Bgp
### The bgp models were taken from the public 4.26.0F repo on arista yang github repo.
./generator -path=models/bgp/4.26.0F -output_file=pkg/bgp.go -package_name=bgp -generate_fakeroot -fakeroot_name=device -compress_paths=true -shorten_enum_leaf_names -typedef_enum_with_defmod -exclude_modules=ietf-interfaces models/bgp/4.26.0F/openconfig-bgp.yang

### Interface models golang was built within the ygot repo.
/generator -path=yang -output_file=pkg/ocdemo/interface.go -package_name=interface -generate_fakeroot -fakeroot_name=device -compress_paths=true -shorten_enum_leaf_names -typedef_enum_with_defmod -exclude_modules=ietf-interfaces yang/openconfig-interfaces.yang



