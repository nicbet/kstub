package commands

import (
	"fmt"
	"log"

	. "github.com/kuyio/kstub/types"
	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"
)

var (
	serviceCmd = &cobra.Command{
		Use:   "service",
		Short: "Generate a service manifest",
		Long:  ``,
		Run:   service,
	}
)

func service(cmd *cobra.Command, args []string) {

	d := Service{
		TypeMeta: TypeMeta{
			APIVersion: "v1",
			Kind:       "Service",
		},
		ObjectMeta: ObjectMeta{
			Name: name + "-service",
		},
		Spec: ServiceSpec{
			Type:     ServiceType(atype),
			Selector: map[string]string{"app": name},
			Ports: []ServicePort{
				ServicePort{
					Protocol:   "TCP",
					Port:       port,
					TargetPort: port,
				},
			},
		},
	}

	s, err := yaml.Marshal(&d)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Printf("---\n%v\n", string(s))
}

func includeServiceFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&name, "name", "n", "name", "The name of the service")
	cmd.Flags().StringVarP(&atype, "type", "t", "ClusterIP", "The type of the service")
	cmd.Flags().Int32VarP(&port, "port", "p", 80, "The main container port")
}

func init() {
	includeServiceFlags(serviceCmd)
}
