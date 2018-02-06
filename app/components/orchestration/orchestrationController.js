angular.module('orchestration', [])
    .controller('orchestrationController', ['$scope', '$state', '$window', 'orchestrationService', 'OrchestrationProvider', 'Notifications', 'Pagination',
        function ($scope, $state, $window, orchestrationService, OrchestrationProvider, Notifications, Pagination) {
            $scope.Orchestrationlogs = '';
            $scope.isbisabled=false;

            $scope.selectOrchestration = function() {
                $('#otosBuildSpinner').show();
                $('#notify').show();
                $scope.isbisabled=true;
                orchestrationService.orchestration().then(function success(data) {
                    $scope.Orchestrationlogs=data.Output;

                    Notifications.success('Kubernetes Dashboard setup complete', name);

                    $window.open('http://127.0.0.1:8001/ui/');
                })
                    .catch(function error(err) {
                        $scope.Orchestrationlogs = err.data.err;
                    })
                    .finally(function final() {
                        $('#otosBuildSpinner').hide();
                        $('#notify').hide();
                        $scope.isbisabled=false;
                    });
                $scope.Orchestrationlogs = '';
            }
        }]);
