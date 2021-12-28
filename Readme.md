# Openconfig kubernetes operator

This repo covers the demo of generally using a kubernetes operator to reconcile network objects.

![Alt text](/images/overall.jpg?raw=true "overall")

![Alt text](/images/schema.jpg?raw=true "schema")


The thought behind this is that the kubernetes api via a CRD will keep and maintain the state in this instance BGP neighbors and any time there is any sort of drift with what the operator has state of versus what is currently running on the switch the operator will reconcile the neighbors and push the config through gNMI what it should be.

This is obviously not production ready and more or less an ongoing project.  At this point in 2021 I want to return back to my Christmas break but also feel like I have accomplished something in the past 2 weeks :D


# Motivation

The motivation behind this is mainly due to learning about [crossplane](https://crossplane.io/) many weeks ago but now getting around to this during my christmas break of working on this project.  I am a firm believer in using K8s not just as a way to deploy services and schedule pods but to also use the k8s api as a means of keeping the correct state.  I feel like kubernetes should be and needs to be used as not just for pods but for external resources as well.

# Tooling

 - [ygot](https://github.com/openconfig/ygot)
 - [goarista gNMI](https://github.com/aristanetworks/goarista)
 - [gNMI](https://github.com/openconfig/gnmi)
 - [goyang](https://github.com/openconfig/goyang)
 - [Kubernetes operator framework](https://operatorframework.io/)

 # Demo

Create a kind cluster unless you already have access to a kubernetes cluster:
```
kind create cluster
```

Apply the crd:
```
kubectl apply -f config/crd/bases/oc.ocoperator.com_ocswitches.yaml
```

Apply the cr:
```
kubectl apply -f config/samples/ceos1.yaml
```

Run the program:
```
go run main.go
```


![Alt text](/images/running.jpg?raw=true "running")


## Tested versions
- Kubernetes 1.22.2
- cEOS 4.26.1F
- Operatoe framework v3
- go 1.16

## Todos
- Add support for interfaces
- Add support for vlans
- Add support for system things
- Fix the return portion of connect / get where we key off the list.
- Build container rather than just running through go.
- ztp support

