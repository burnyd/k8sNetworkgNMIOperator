package k8sgnmi

import (
	"context"
	"fmt"

	"github.com/aristanetworks/goarista/gnmi"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

func (c *GNMI_CFG) Set(paths, json, replace string) {
	var cfg = &gnmi.Config{
		Addr:     c.Host,
		Username: c.Username,
		Password: c.Password,
	}

	var setOps []*gnmi.Operation

	path := []string{paths}
	var origin = c.Origin
	ctx := gnmi.NewContext(context.Background(), cfg)
	client, err := gnmi.Dial(cfg)
	if err != nil {
		fmt.Println(err)
	}

	ops := &gnmi.Operation{
		Origin: origin,
		Path:   gnmi.SplitPath(path[0]),
		Type:   "replace",
		Val:    json,
	}

	setOps = append(setOps, ops)
	err = gnmi.Set(ctx, client, setOps)
	if err != nil {
		fmt.Println(err)
	} else {
		log.Log.Info("Set the entry")

	}
}
