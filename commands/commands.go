package commands

import (
	"github.com/codegangsta/cli"
)

func GetCommands() []cli.Command {
	return []cli.Command{
		{
			Name:  "cli",
			Usage: "Starts a shell for executing Kubernetes commands",
		},
		{
			Name:    "set-namespace",
			Aliases: []string{"ns"},
			Usage:   "Sets the working namespace for all commands",
			Action:  set_namespace,
		},
		{
			Name:   "get",
			Usage:  "Display one or many resources",
			Action: wrap_command,
		},
		{
			Name:   "set",
			Usage:  "Set specific features on objects",
			Action: wrap_command,
		},
		{
			Name:   "describe",
			Usage:  "Show details of a specific resource or group of resources",
			Action: wrap_command,
		},
		{
			Name:   "create",
			Usage:  "Create a resource by filename or stdin",
			Action: wrap_command,
		},
		{
			Name:   "replace",
			Usage:  "Replace a resource by filename or stdin.",
			Action: wrap_command,
		},
		{
			Name:   "patch",
			Usage:  "Update field(s) of a resource using strategic merge patch.",
			Action: wrap_command,
		},
		{
			Name:   "delete",
			Usage:  "Delete resources by filenames, stdin, resources and names, or by resources and label selector.",
			Action: wrap_command,
		},
		{
			Name:   "edit",
			Usage:  "Edit a resource on the server",
			Action: wrap_command,
		},
		{
			Name:   "apply",
			Usage:  "Apply a configuration to a resource by filename or stdin",
			Action: wrap_command,
		},
		{
			Name:   "namespace",
			Usage:  "SUPERSEDED: Set and view the current Kubernetes namespace",
			Action: wrap_command,
		},
		{
			Name:   "logs",
			Usage:  "Print the logs for a container in a pod.",
			Action: wrap_command,
		},
		{
			Name:   "rolling-update",
			Usage:  "Perform a rolling update of the given ReplicationController.",
			Action: wrap_command,
		},
		{
			Name:   "scale",
			Usage:  "Set a new size for a Deployment, ReplicaSet, Replication Controller, or Job.",
			Action: wrap_command,
		},
		{
			Name:   "cordon",
			Usage:  "Mark node as unschedulable",
			Action: wrap_command,
		},
		{
			Name:   "drain",
			Usage:  "Drain node in preparation for maintenance",
			Action: wrap_command,
		},
		{
			Name:   "uncordon",
			Usage:  "Mark node as schedulable",
			Action: wrap_command,
		},
		{
			Name:   "attach",
			Usage:  "Attach to a running container.",
			Action: wrap_command,
		},
		{
			Name:   "exec",
			Usage:  "Execute a command in a container.",
			Action: wrap_command,
		},
		{
			Name:   "port-forward",
			Usage:  "Forward one or more local ports to a pod.",
			Action: wrap_command,
		},
		{
			Name:   "proxy",
			Usage:  "Run a proxy to the Kubernetes API server",
			Action: wrap_command,
		},
		{
			Name:   "run",
			Usage:  "Run a particular image on the cluster.",
			Action: wrap_command,
		},
		{
			Name:   "expose",
			Usage:  "Take a replication controller, service, deployment or pod and expose it as a new Kubernetes Service",
			Action: wrap_command,
		},
		{
			Name:   "autoscale",
			Usage:  "Auto-scale a Deployment, ReplicaSet, or ReplicationController",
			Action: wrap_command,
		},
		{
			Name:   "rollout",
			Usage:  "rollout manages a deployment",
			Action: wrap_command,
		},
		{
			Name:   "label",
			Usage:  "Update the labels on a resource",
			Action: wrap_command,
		},
		{
			Name:   "annotate",
			Usage:  "Update the annotations on a resource",
			Action: wrap_command,
		},
		{
			Name:   "taint",
			Usage:  "Update the taints on one or more nodes",
			Action: wrap_command,
		},
		{
			Name:   "config",
			Usage:  "config modifies kubeconfig files",
			Action: wrap_command,
		},
		{
			Name:   "cluster-info",
			Usage:  "Display cluster info",
			Action: wrap_command,
		},
		{
			Name:   "api-versions",
			Usage:  "Print the supported API versions on the server, in the form of \"group/version\".",
			Action: wrap_command,
		},
		{
			Name:   "version",
			Usage:  "Print the client and server version information.",
			Action: wrap_command,
		},
		{
			Name:   "explain",
			Usage:  "Documentation of resources.",
			Action: wrap_command,
		},
		{
			Name:   "convert",
			Usage:  "Convert config files between different API versions",
			Action: wrap_command,
		},
		{
			Name:   "completion",
			Usage:  "Output shell completion code for the given shell (bash or zsh)",
			Action: wrap_command,
		},
		{
			Name:    "help",
			Aliases: []string{"h"},
			Usage:   "Shows a list of commands or help for one command",
			Action:  wrap_command,
		},
	}
}
