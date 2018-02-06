package bolt
import (
	"dolphin/api"
	"os/exec"
	"bytes"
	"errors"
	"time"
)
var Outputval string
var chk_certs_created bool = false;
type OrchestrationService struct {
	store *Store
}

// CreateEndpoint assign an ID to a new endpoint and saves it.
func (service *OrchestrationService) BuildOrchestration(otos *dockm.OToS, endpoint *dockm.Endpoint) ( error , string ) {
	// Check whether certificates are already created or not
	if (chk_certs_created == false) {
		var kube_dash_yml_url string = "https://raw.githubusercontent.com/kubernetes/dashboard/master/src/deploy/recommended/kubernetes-dashboard.yaml"

		command:="kubectl"
		var kube_create []string
		var kube_delete []string

		kube_delete = []string{"delete","-f",kube_dash_yml_url}
		kube_create = []string{"create","-f",kube_dash_yml_url}

		cmd2 := exec.Command(command,kube_delete...)

		cmd := exec.Command(command,kube_create...)

		var outb, errb bytes.Buffer
		cmd.Stdout = &outb
		cmd.Stderr = &errb

		cmd1 := exec.Command(command,"proxy")

		cmd2.Run()
		err := cmd.Run()
		err1 := cmd1.Start()

		stdout:=outb.String()
		stderr:=errb.String()

		Outputval:=stdout+stderr
		chk_certs_created = true;
		var customError error = nil
		if err != nil {
			chk_certs_created = false
			customError = errors.New(Outputval)
		}
		if err1 != nil {
			chk_certs_created = false
			customError = errors.New(Outputval)
		}
		time.Sleep(50 * time.Second)
		return customError ,Outputval

	} else {
		Outputval = "Kubernetes dashboard configuration already setup."
		var customError error = nil
		return customError ,Outputval
	}

}
