angular.module('dockm').component('porContainerList', {
  templateUrl: 'app/directives/containerList/porContainerList.html',
  controller: 'porContainerListController',
  bindings: {
    'containers': '<'
  }
});
