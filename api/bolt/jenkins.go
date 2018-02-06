package bolt

import (
	"github.com/bndr/gojenkins"

)

func Jenkins_CICD()  {

	jenkins := gojenkins.CreateJenkins(nil, "http://115.249.182.102:10000/", "sapan", "72799493e63a7bebc7cde035c626b06f")
	_, err := jenkins.Init()


	if err != nil {
	panic("Something Went Wrong")
	}


	configString := `<?xml version='1.0' encoding='UTF-8'?>
<project>
  <actions/>
  <description></description>
  <keepDependencies>false</keepDependencies>
  <properties/>
  <scm class="hudson.plugins.git.GitSCM" plugin="git@3.5.1">
    <configVersion>2</configVersion>
    <userRemoteConfigs>
      <hudson.plugins.git.UserRemoteConfig>
        <url>https://github.com/Click2Cloud/nodejs-example</url>
      </hudson.plugins.git.UserRemoteConfig>
    </userRemoteConfigs>
    <branches>
      <hudson.plugins.git.BranchSpec>
        <name>*/master</name>
      </hudson.plugins.git.BranchSpec>
    </branches>
    <doGenerateSubmoduleConfigurations>false</doGenerateSubmoduleConfigurations>
    <submoduleCfg class="list"/>
    <extensions/>
  </scm>
  <canRoam>true</canRoam>
  <disabled>false</disabled>
  <blockBuildWhenDownstreamBuilding>false</blockBuildWhenDownstreamBuilding>
  <blockBuildWhenUpstreamBuilding>false</blockBuildWhenUpstreamBuilding>
  <triggers>
    <com.cloudbees.jenkins.GitHubPushTrigger plugin="github@1.28.0">
      <spec></spec>
    </com.cloudbees.jenkins.GitHubPushTrigger>
  </triggers>
  <concurrentBuild>false</concurrentBuild>
  <builders>
    <jenkins.plugins.http__request.HttpRequest plugin="http_request@1.8.20">
      <url>http://115.249.182.102:20000/api/apptocontainer</url>
      <ignoreSslErrors>false</ignoreSslErrors>
      <httpMode>POST</httpMode>
      <passBuildParameters>false</passBuildParameters>
      <validResponseCodes>100:399</validResponseCodes>
      <validResponseContent></validResponseContent>
      <acceptType>APPLICATION_JSON</acceptType>
      <contentType>APPLICATION_JSON</contentType>
      <outputFile></outputFile>
      <timeout>0</timeout>
      <consoleLogResponseBody>true</consoleLogResponseBody>
      <authentication></authentication>
      <requestBody>{ &quot;BaseImage&quot;: &quot;click2cloud/nodejs-010-centos7&quot;, &quot;GitUrl&quot;: &quot;https://github.com/Click2Cloud/nodejs-example&quot;, &quot;ImageName&quot;:&quot;hello-node&quot;, &quot;EndPointId&quot;: 2}</requestBody>
      <customHeaders>
        <pair>
          <name>Authorization</name>
          <value>Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcm5hbWUiOiJhZG1pbiIsInJvbGUiOjEsImV4cCI6MTUwNTc0MDg4M30.Cda977lVHNZIL1CGziNYNsM4uCirW4WA_rAKBqGEOYU</value>
          <maskValue>true</maskValue>
        </pair>
      </customHeaders>
    </jenkins.plugins.http__request.HttpRequest>
  </builders>
  <publishers/>
  <buildWrappers/>
</project>`

	jenkins.CreateJob(configString, "ApptoContainer")
}