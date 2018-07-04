package commands

import (
	"fmt"
	"log"

	. "github.com/nicbet/kstub/types"
	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"
)

var (
	// Command
	deploymentCmd = &cobra.Command{
		Use:   "deployment",
		Short: "Generate a deployment manifest",
		Long:  ``,
		Run:   deployment,
	}
)

func deployment(cmd *cobra.Command, args []string) {

	d := Deployment{
		TypeMeta: TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "Deployment",
		},
		ObjectMeta: ObjectMeta{
			Name:   name + "-deployment",
			Labels: map[string]string{"app": name},
		},
		Spec: DeploymentSpec{
			Selector: LabelSelector{
				MatchLabels: map[string]string{"app": name},
			},
			Replicas: replicas,
			Template: PodTemplate{
				ObjectMeta: ObjectMeta{
					Labels: map[string]string{"app": name},
				},
				Spec: PodSpec{
					Containers: []Container{
						Container{
							Name:  name,
							Image: name + ":latest",
							Ports: []ContainerPort{
								ContainerPort{
									ContainerPort: port,
									HostPort:      port,
									Name:          name + "-port",
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

func includeDeploymentFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&name, "name", "n", "name", "The name of the deployment")
	cmd.Flags().Int32VarP(&replicas, "replicas", "r", 1, "The number of desired pod replicas")
	cmd.Flags().Int32VarP(&port, "port", "p", 80, "The main container port")

	// cmd.MarkFlagRequired("name")
	// cmd.MarkFlagRequired("replicas")
}

func init() {
	includeDeploymentFlags(deploymentCmd)

}
