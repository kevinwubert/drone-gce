package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	log "github.com/sirupsen/logrus"
)

func main() {
	googleCredentials := os.Getenv("PLUGIN_JSON_KEY")

	// write google creds to a file
	buf := []byte(googleCredentials)
	err := ioutil.WriteFile("/certs/svc_account.json", buf, 0644)
	if err != nil {
		log.WithError(err).Fatalf("could not write google credentials to file")
	}

	projectID := os.Getenv("PLUGIN_PROJECT_ID")
	zone := os.Getenv("PLUGIN_ZONE")
	instanceName := os.Getenv("PLUGIN_INSTANCE_NAME")
	sshCommandsStr := os.Getenv("PLUGIN_COMMANDS")
	fmt.Printf("Running drone-gce plugin on instance %s in zone %s in project %s\n\n", instanceName, zone, projectID)

	sshCommands := strings.Split(sshCommandsStr, ",")
	for _, sshCmd := range sshCommands {
		cmdStr := fmt.Sprintf("compute ssh --quiet --project %s --zone %s %s --command \"%s\"", projectID, zone, instanceName, sshCmd)
		cmdSplit := strings.Split(cmdStr, " ")

		fmt.Printf("Executing %s...\n", sshCmd)
		cmd := exec.Command("gcloud", cmdSplit...)

		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			fmt.Println("Exiting...")
			return
		}

		fmt.Println(string(out))
	}
}
