angular.module('dockm').component('porServiceList', {
  templateUrl: 'app/directives/serviceList/porServiceList.html',
  controller: 'porServiceListController',
  bindings: {
    'services': '<',
    'nodes': '<'
  }
});
