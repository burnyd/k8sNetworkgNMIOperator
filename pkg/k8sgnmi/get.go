package k8sgnmi

import (
	"context"
	"fmt"

	"github.com/aristanetworks/goarista/gnmi"
	pb "github.com/openconfig/gnmi/proto/gnmi"
)



type Features struct {
	Bgp bool
}

func Get(ctx context.Context, client pb.GNMIClient, req *pb.GetRequest) ([]string, error) {
	resp, err := client.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	s := []string{}
	for _, notif := range resp.Notification {
		for _, update := range notif.Update {
			s = append(s, gnmi.StrUpdateVal(update))
		}
	}
	return s, nil
}

func (c *GNMI_CFG) Connect() {
	var cfg = &gnmi.Config{
		Addr:     c.Host,
		Username: c.Username,
		Password: c.Password,
	}
	paths := []string{c.Path}
	var origin = c.Origin
	ctx := gnmi.NewContext(context.Background(), cfg)
	client, err := gnmi.Dial(cfg)
	if err != nil {
		fmt.Println(err)
	}
	req, err := gnmi.NewGetRequest(gnmi.SplitPaths(paths), origin)
	if err != nil {
		fmt.Println(err)
	}
	if cfg.Addr != "" {
		if req.Prefix == nil {
			req.Prefix = &pb.Path{}
		}
		req.Prefix.Target = cfg.Addr
	}
	R, err := Get(ctx, client, req)
	if err != nil {
		fmt.Println(err)
	} else {
		Response := R[0]
		Byteresponse := []byte(Response)
		c.JsonData = Byteresponse

	}
}
