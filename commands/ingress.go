package commands

import (
	"fmt"
	"log"

	. "github.com/nicbet/kstub/types"
	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"
)

var (
	ingressCmd = &cobra.Command{
		Use:   "ingress",
		Short: "Generate a ingress manifest",
		Long:  ``,
		Run:   ingress,
	}
)

func ingress(cmd *cobra.Command, args []string) {

	d := Ingress{
		TypeMeta: TypeMeta{
			APIVersion: "extensions/v1beta1",
			Kind:       "Ingress",
		},
		ObjectMeta: ObjectMeta{
			Name: name + "-ingress",
			Annotations: map[string]string{
				"nginx.ingress.kubernetes.io/rewrite-target": "/",
				"kubernetes.io/ingress.class":                "nginx",
			},
		},
		Spec: IngressSpec{
			TLS: []IngressTLS{
				IngressTLS{
					Hosts:      []string{"foo.bar.com"},
					SecretName: name + "-tls-secret",
				},
			},
			Rules: []IngressRule{
				IngressRule{
					Host: "foo.bar.com",
					IngressRuleValue: IngressRuleValue{
						HTTP: &HTTPIngressRuleValue{
							Paths: []HTTPIngressPath{
								HTTPIngressPath{
									Path: "/foo",
									Backend: IngressBackend{
										ServiceName: name + "-service",
										ServicePort: port,
									},
								},
							},
						},
					},
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

func includeIngressFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&name, "name", "n", "name", "The name of the ingress")
	cmd.Flags().Int32VarP(&port, "port", "p", 80, "The main container port")
}

func init() {
	includeIngressFlags(ingressCmd)
}
