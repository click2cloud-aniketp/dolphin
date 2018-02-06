/*eslint linebreak-style: [2, "windows"]*/
angular.module('appToContainer', [])
    .controller('appToContainerController', ['$scope', '$state', 'appToContainerService','AppToContainerProvider', 'Notifications', 'Pagination',
        function ($scope, $state, appToContainerService,AppToContainerProvider, Notifications, Pagination) {
            $scope.appToContainerlogs = '';
            $scope.formValues = {
                BaseImage: '',
                GitUrl: '',
                ImageName: ''
            };
            $scope.builderImages = {
                'Click2Cloud Python Builder Image' :  'click2cloud/python-33-centos7',
                'Click2Cloud NodeJs Builder Image'  :   'click2cloud/nodejs-010-centos7',
                'Click2Cloud Ruby Builder Image'  :   'click2cloud/ruby-20-centos7',
                'Click2Cloud PHP Builder Image'  :   'click2cloud/php-55-centos7',
                'Click2Cloud Perl Builder Image'  :   'click2cloud/perl-516-centos7',
                'Click2Cloud .NET Core Builder Image' : 'click2cloud/aspnet-core-centos7',

                'Click2Cloud .NET legacy Builder Image' : 'click2cloud/aspnet-4.5-centos7',
                'Click2Cloud J2EE Builder Image'   :   'dishawani/jadu',
                'Click2Cloud nodejs-010 Builder Image' : '',
                'Click2Cloud nodejs-4 Builder Image' : '',
                'Click2Cloud nodejs-latest Builder Image' : '',
                'Click2Cloud nodejs-example Builder Image' : '',
                'Click2Cloud nodejs-mongodb-example Builder Image' : '',
                'Click2Cloud ruby-20 Builder Image' : '',
                'Click2Cloud ruby-22 Builder Image' : '',
                'Click2Cloud ruby-23 Builder Image' : '',
                'Click2Cloud ruby-latest Builder Image' : '',
                'Click2Cloud perl-516 Builder Image' : '',
                'Click2Cloud perl-520 Builder Image' : '',
                'Click2Cloud perl-524 Builder Image' : '',
                'Click2Cloud php-55 Builder Image' : '',
                'Click2Cloud php-56 Builder Image' : '',
                'Click2Cloud php-70 Builder Image' : '',
                'Click2Cloud cakephp-example Builder Image' : '',
                'Click2Cloud cakephp-mysql-example Builder Image' : '',
                'Click2Cloud python-27 Builder Image' : '',
                'Click2Cloud python-33 Builder Image' : '',
                'Click2Cloud python-34 Builder Image' : '',
                'Click2Cloud python-35 Builder Image' : '',
                'Click2Cloud mysql-55 Builder Image' : '',
                'Click2Cloud mysql-56 Builder Image' : '',
                'Click2Cloud mysql-57 Builder Image' : '',
                'Click2Cloud mysql-ephemeral Builder Image' : '',
                'Click2Cloud mysql-persistent Builder Image' : '',
                'Click2Cloud postgresql-92 Builder Image' : '',
                'Click2Cloud postgresql-94 Builder Image' : '',
                'Click2Cloud postgresql-95 Builder Image' : '',
                'Click2Cloud postgresql-ephemeral Builder Image' : '',
                'Click2Cloud postgresql-persistent Builder Image' : '',
                'Click2Cloud mongodb-24 Builder Image' : '',
                'Click2Cloud mongodb-26 Builder Image' : '',
                'Click2Cloud mongodb-32 Builder Image' : '',
                'Click2Cloud mongodb-ephemeral Builder Image' : '',
                'Click2Cloud mongodb-persistent Builder Image' : '',
                'Click2Cloud jenkins-1 Builder Image' : '',
                'Click2Cloud jenkins-2 Builder Image' : '',
                'Click2Cloud jenkins-persistent Builder Image' : '',
                'Click2Cloud jenkins-ephemeral Builder Image' : '',
                'Click2Cloud openjdk-18 Builder Image' : '',
                'Click2Cloud dotnet/dotnetcore-10 Builder Image' : '',
                'Click2Cloud dotnet/dotnetcore-11 Builder Image' : '',
                'Click2Cloud fis-java Builder Image' : '',
                'Click2Cloud fis-karaf Builder Image' : '',
                'Click2Cloud jboss-decisionserver62 Builder Image' : '',
                'Click2Cloud jboss-eap64 / UAT-1.1 Builder Image' : '',
                'Click2Cloud jboss-eap64 / UAT-1.2 Builder Image' : '',
                'Click2Cloud jboss-eap64 / UAT-1.3 Builder Image' : '',
                'Click2Cloud jboss-eap70 Builder Image' : '',
                'Click2Cloud jboss-webserver30-tomcat7 / UAT-1.1 Builder Image' : '',
                'Click2Cloud jboss-webserver30-tomcat7 / UAT-1.2 Builder Image' : '',
                'Click2Cloud jboss-webserver30-tomcat8 / UAT-1.1 Builder Image' : '',
                'Click2Cloud jboss-webserver30-tomcat8 / UAT-1.2 Builder Image' : '',
                'Click2Cloud asp.net 4.5 Builder Image' : '',
                'Click2Cloud aspnet-45-mongodb-ex Builder Image' : '',
                'Click2Cloud aspnet-45-mssqlextdb Builder Image' : '',
                'Click2Cloud aspnet-45-externaldb-ex Builder Image' : '',
                'Click2Cloud aspnet-45-postgresql-ex Builder Image' : '',
                'Click2Cloud aspnet-45-mysql-ex Builder Image' : '',
                'Click2Cloud aspnet-core-mysql-ex Builder Image' : '',
                'Click2Cloud aspnet-core-postgresql Builder Image' : '',
                'Click2Cloud aspnet-core-mongodb-ex Builder Image' : '',
                'Click2Cloud aspnet-core-external-ex Builder Image' : '',
                'Click2Cloud aspnet-core-example Builder Image' : '',
                'Click2Cloud dancer-example Builder Image' : '',
                'Click2Cloud dancer-mysql-example Builder Image' : '',
                'Click2Cloud django-example Builder Image' : '',
                'Click2Cloud django-psql-example Builder Image' : '',
                'Click2Cloud amq62-basic Builder Image' : '',
                'Click2Cloud amq62-persistent Builder Image' : '',
                'Click2Cloud amq62-persistent-ssl Builder Image' : '',
                'Click2Cloud amq62-ssl Builder Image' : '',
                'Click2Cloud datagrid65-basic Builder Image' : '',
                'Click2Cloud datagrid65-https Builder Image' : '',
                'Click2Cloud datagrid65-mysql Builder Image' : '',
                'Click2Cloud datagrid65-mysql-persistent Builder Image' : '',
                'Click2Cloud datagrid65-postgresql Builder Image' : '',
                'Click2Cloud datagrid65-postgresql-persistent Builder Image' : '',
                'Click2Cloud decisionserver62-amq-s2i Builder Image' : '',
                'Click2Cloud decisionserver62-https-s2i Builder Image' : '',
                'Click2Cloud decisionserver62-basic-s2i Builder Image' : '',
                'Click2Cloud eap64-amq-persistent-s2i Builder Image' : '',
                'Click2Cloud eap64-amq-s2i Builder Image' : '',
                'Click2Cloud eap64-basic-s2i Builder Image' : '',
                'Click2Cloud eap64-https-s2i Builder Image' : '',
                'Click2Cloud eap64-mongodb-persistent-s2i Builder Image' : '',
                'Click2Cloud eap64-mongodb-s2i Builder Image' : '',
                'Click2Cloud eap64-mysql-persistent-s2i Builder Image' : '',
                'Click2Cloud eap64-mysql-s2i Builder Image' : '',
                'Click2Cloud eap64-postgresql-persistent-s2i Builder Image' : '',
                'Click2Cloud eap64-postgresql-s2i Builder Image' : '',
                'Click2Cloud eap64-sso-s2i Builder Image' : '',
                'Click2Cloud eap70-amq-persistent-s2i Builder Image' : '',
                'Click2Cloud eap70-amq-s2i Builder Image' : '',
                'Click2Cloud eap70-basic-s2i Builder Image' : '',
                'Click2Cloud eap70-https-s2i Builder Image' : '',
                'Click2Cloud eap70-mongodb-persistent-s2i Builder Image' : '',
                'Click2Cloud eap70-mongodb-s2i Builder Image' : '',
                'Click2Cloud eap70-mysql-persistent-s2i Builder Image' : '',
                'Click2Cloud eap70-mysql-s2i Builder Image' : '',
                'Click2Cloud eap70-postgresql-persistent-s2i Builder Image' : '',
                'Click2Cloud eap70-postgresql-s2i Builder Image' : '',
                'Click2Cloud jws30-tomcat7-basic-s2i Builder Image' : '',
                'Click2Cloud jws30-tomcat7-https-s2i Builder Image' : '',
                'Click2Cloud jws30-tomcat7-mongodb-persistent-s2i Builder Image' : '',
                'Click2Cloud jws30-tomcat7-mongodb-s2i Builder Image' : '',
                'Click2Cloud jws30-tomcat7-mysql-persistent-s2i Builder Image' : '',
                'Click2Cloud jws30-tomcat7-mysql-s2i Builder Image' : '',
                'Click2Cloud jws30-tomcat7-postgresql-persistent-s2i Builder Image' : '',
                'Click2Cloud jws30-tomcat7-postgresql-s2i Builder Image' : '',
                'Click2Cloud jws30-tomcat8-basic-s2i Builder Image' : '',
                'Click2Cloud jws30-tomcat8-https-s2i Builder Image' : '',
                'Click2Cloud jws30-tomcat8-mongodb-persistent-s2i Builder Image' : '',
                'Click2Cloud jws30-tomcat8-mongodb-s2i Builder Image' : '',
                'Click2Cloud jws30-tomcat8-mysql-persistent-s2i Builder Image' : '',
                'Click2Cloud jws30-tomcat8-mysql-s2i Builder Image' : '',
                'Click2Cloud jws30-tomcat8-postgresql-persistent-s2i Builder Image' : '',
                'Click2Cloud jws30-tomcat8-postgresql-s2i Builder Image' : '',
                'Click2Cloud sso70-basic Builder Image' : '',
                'Click2Cloud sso70-mysql Builder Image' : '',
                'Click2Cloud sso70-mysql-persistent Builder Image' : '',
                'Click2Cloud sso70-postgresql Builder Image' : '',
                'Click2Cloud sso70-postgresql-persistent Builder Image' : '',
                'Click2Cloud rails-postgresql-example Builder Image' : '',
                'Click2Cloud logging-deployer-template Builder Image' : '',
                'Click2Cloud metrics-deployer-template Builder Image' : ''

            };
            $scope.isbisabled=false;

            //$scope.builderImage = ['centos/python-35-centos7','click2cloud/perl-516-centos7','click2cloud/nodejs-010-centos7','click2cloud/ruby-20-centos7','click2cloud/php-55-centos7'];

            $scope.buildApptocontainer = function() {
                $('#atocBuildSpinner').show();
                $('#notify').show();
                $scope.isbisabled=true;
                var BaseImage = $scope.formValues.BaseImage;
                if (BaseImage in $scope.builderImages) {
                    BaseImage = $scope.builderImages[BaseImage];
                }
                var GitUrl =$scope.formValues.GitUrl;
                var ImageName =$scope.formValues.ImageName;
                appToContainerService.appToContainer(BaseImage, GitUrl, ImageName).then(function success(data) {
                    $scope.appToContainerlogs=data.Output;
                    console.log('controller function calling');
                    console.log($scope.appToContainerlogs);
                    Notifications.success('Image created', name);
                    //$state.reload();
                })
                    .catch(function error(err) {
                        Notifications.error('Failure','', 'App to Container Image Build failed');
                        $scope.appToContainerlogs = err.data.err;

                    })
                    .finally(function final() {
                        $('#atocBuildSpinner').hide();
                        $('#notify').hide();
                        $scope.isbisabled=false;
                    });
                $scope.appToContainerlogs = '';
            };
        }]);
