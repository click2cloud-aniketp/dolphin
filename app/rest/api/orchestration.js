/*eslint linebreak-style: [2, "windows"]*/
angular.module('dockm.rest')
    .factory('Orchestration', ['$resource', 'API_ENDPOINT_ORCHESTRATION', function OrchestrationFactory($resource, API_ENDPOINT_ORCHESTRATION) {
        'use strict';
        return $resource(API_ENDPOINT_ORCHESTRATION + '/:id/:action', {}, {
            create: { method: 'POST' }
            // get: { method: 'GET' }
        });
    }]);