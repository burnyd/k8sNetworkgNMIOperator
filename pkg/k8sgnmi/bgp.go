package k8sgnmi

import (
	"encoding/json"
	"fmt"

	"github.com/burnyd/k8sNetworkgNMIOperator/pkg/bgp"
	"github.com/openconfig/ygot/ygot"
)

type BgpData struct {
	OpenconfigNetworkInstanceNeighbor []struct {
		AfiSafis struct {
			AfiSafi []struct {
				AfiSafiName string `json:"afi-safi-name"`
				Config      struct {
					AfiSafiName string `json:"afi-safi-name"`
				} `json:"config"`
				State struct {
					AfiSafiName string `json:"afi-safi-name"`
					Prefixes    struct {
						AristaBgpAugmentsBestEcmpPaths int `json:"arista-bgp-augments:best-ecmp-paths"`
						AristaBgpAugmentsBestPaths     int `json:"arista-bgp-augments:best-paths"`
						Installed                      int `json:"installed"`
						Received                       int `json:"received"`
						Sent                           int `json:"sent"`
					} `json:"prefixes"`
				} `json:"state"`
			} `json:"afi-safi"`
		} `json:"afi-safis"`
		Config struct {
			NeighborAddress string `json:"neighbor-address"`
			PeerAs          uint32 `json:"peer-as"`
			SendCommunity   string `json:"send-community"`
		} `json:"config"`
		EbgpMultihop struct {
			Config struct {
				MultihopTTL int `json:"multihop-ttl"`
			} `json:"config"`
			State struct {
				MultihopTTL int `json:"multihop-ttl"`
			} `json:"state"`
		} `json:"ebgp-multihop"`
		NeighborAddress string `json:"neighbor-address"`
		State           struct {
			EstablishedTransitions string `json:"established-transitions"`
			LastEstablished        string `json:"last-established"`
			Messages               struct {
				Received struct {
					Update string `json:"UPDATE"`
				} `json:"received"`
				Sent struct {
					Update string `json:"UPDATE"`
				} `json:"sent"`
			} `json:"messages"`
			NeighborAddress string `json:"neighbor-address"`
			PeerAs          uint32 `json:"peer-as"`
			SendCommunity   string `json:"send-community"`
			SessionState    string `json:"session-state"`
		} `json:"state"`
		Transport struct {
			State struct {
				RemoteAddress string `json:"remote-address"`
				RemotePort    int    `json:"remote-port"`
			} `json:"state"`
		} `json:"transport"`
	} `json:"openconfig-network-instance:neighbor"`
}

func (c *GNMI_CFG) ReturnBgp() map[string]uint32 {
	bgp := BgpData{}

	json.Unmarshal(c.JsonData, &bgp)
	bgpmap := map[string]uint32{}
	for _, i := range bgp.OpenconfigNetworkInstanceNeighbor {
		bgpmap[i.Config.NeighborAddress] = i.Config.PeerAs
	}
	return bgpmap

}

func (c *GNMI_CFG) SetBgpYgot(as uint32, bgpmap map[string]uint32) {
	b := &bgp.Bgp{
		Global: &bgp.Bgp_Global{
			As: ygot.Uint32(as),
		},
	}

	for k, v := range bgpmap {
		b.NewNeighbor(k)

		b.Neighbor[k] = &bgp.Bgp_Neighbor{
			PeerAs:          ygot.Uint32(v),
			NeighborAddress: ygot.String(k),
			Enabled:         ygot.Bool(true),
		}
	}
	json, err := ygot.EmitJSON(b, &ygot.EmitJSONConfig{
		Format: ygot.RFC7951,
		Indent: "  ",
		RFC7951Config: &ygot.RFC7951JSONConfig{
			AppendModuleName: true,
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	//return json
	BgPPath := "/network-instances/network-instance[name=default]/protocols/protocol[name=BGP]/bgp/"
	c.Set(BgPPath, json, "replace")
}
