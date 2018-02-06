angular.module('dockm.rest')
.factory('Template', ['$resource', 'API_ENDPOINT_TEMPLATES', function TemplateFactory($resource, API_ENDPOINT_TEMPLATES) {
  return $resource(API_ENDPOINT_TEMPLATES, {}, {
    get: {method: 'GET', isArray: true}
  });
}]);
