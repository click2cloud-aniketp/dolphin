package bolt//click2cloud-apptocontainer
import (
	"dolphin/api"
	"os/exec"
	"bytes"
	 "strconv"

	"strings"
	"errors"

)
var Output string
type AppToContainerService struct {
	store *Store
}

// CreateEndpoint assign an ID to a new endpoint and saves it.

func (service *AppToContainerService) BuildAppToContainer(atoc *dockm.AToC, endpoint *dockm.Endpoint) ( error , string ) {

	//endpoint, endpointErr = service.store.EndpointService.Endpoint(dockm.EndpointID(atoc.EndPointId))

	/*var TLSCertPath="C:\\data"+file.TLSStorePath+strconv.Itoa(atoc.EndPointId)+"cert.pem"
	var TLSCaPath  ="C:\\data"+file.TLSStorePath+strconv.Itoa(atoc.EndPointId)+"ca.pem"
	var TLkeyPath="C:\\data"+file.TLSStorePath+strconv.Itoa(atoc.EndPointId)+"key.pem"
	var DockerURL="tcp://"+atoc.EndPointUrl+":2376"
	var TLS="true"*/
	command:="s2i"
	var comarg []string

	if strings.HasPrefix(endpoint.URL, "tcp://") &&  endpoint.TLS{
		comarg = []string{"build",atoc.GitUrl,atoc.BaseImage,atoc.ImageName,"--ca",endpoint.TLSCACertPath,"--cert",endpoint.TLSCertPath,"--key",endpoint.TLSKeyPath,"--tls",strconv.FormatBool(endpoint.TLS), "--url",endpoint.URL}
	}else /*if strings.HasPrefix(endpoint.URL, "unix://")*/ {
		comarg = []string{"build",atoc.GitUrl,atoc.BaseImage,atoc.ImageName}
	}

    Jenkins_CICD()
	//comarg := []string{"build",atoc.GitUrl,atoc.BaseImage,atoc.ImageName,"--ca",TLSCaPath,"--cert",TLSCertPath,"--key",TLkeyPath,"--tls",TLS, "--url",DockerURL}
	//comarg := []string{"build",atoc.GitUrl,atoc.BaseImage,atoc.ImageName,"--ca",endpoint.TLSCACertPath,"--cert",endpoint.TLSCertPath,"--key",endpoint.TLSKeyPath,"--tls",strconv.FormatBool(endpoint.TLS), "--url",endpoint.URL}
	cmd := exec.Command(command,comarg...)
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb

	err := cmd.Run()
	stdout:=outb.String()
	stderr:=errb.String()
	Output=stdout+stderr
	var customError error = nil
	if err != nil {
		//Output=stdout+stderr
		customError = errors.New(Output)
	}
	return customError ,Output
}
