angular.module('dockm.rest')
.factory('Network', ['$resource', 'API_ENDPOINT_ENDPOINTS', 'EndpointProvider', function NetworkFactory($resource, API_ENDPOINT_ENDPOINTS, EndpointProvider) {
  'use strict';
  return $resource(API_ENDPOINT_ENDPOINTS + '/:endpointId/docker/networks/:id/:action', {
    id: '@id',
    endpointId: EndpointProvider.endpointID
  },
  {
    query: {method: 'GET', isArray: true},
    get: {method: 'GET'},
    create: {method: 'POST', params: {action: 'create'}, transformResponse: genericHandler},
    remove: { method: 'DELETE', transformResponse: genericHandler },
    connect: {method: 'POST', params: {action: 'connect'}},
    disconnect: {method: 'POST', params: {action: 'disconnect'}}
  });
}]);
