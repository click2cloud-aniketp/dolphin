/*eslint linebreak-style: [2, "windows"]*/
angular.module('dockm.rest')
.factory('AppToContainer', ['$resource', 'API_ENDPOINT_APPTOCONTAINER', function appToContainerFactory($resource, API_ENDPOINT_APPTOCONTAINER) {
  'use strict';
  return $resource(API_ENDPOINT_APPTOCONTAINER + '/:id/:action', {}, {
      create: { method: 'POST' },
      // get: { method: 'GET' }
  });
}]);

