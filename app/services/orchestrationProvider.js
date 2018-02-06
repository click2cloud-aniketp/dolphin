angular.module('dockm.services')
    .factory('OrchestrationProvider', ['LocalStorage', function OrchestrationProviderFactory(LocalStorage) {
        'use strict';
        var service = {};
        var orchestration = {};

        service.initialize = function() {
            var EndPointId = LocalStorage.getEndPointId();
            var EndPointUrl = LocalStorage.getEndPointUrl();
            if (EndPointId) {
                orchestration.EndPointId = EndPointId;
            }
            if (EndPointUrl) {
                orchestration.EndPointUrl = EndPointUrl;
            }
        };

        service.clean = function() {
            orchestration = {};
        };
        service.endpointID = function() {
            return orchestration.ID;
        };
        service.setEndpointID = function(id) {
            endpoint.ID = id;
            LocalStorage.storeEndpointID(id);
        };
        service.EndpointUrl = function() {
            return endpoint.ID;
        };
        service.setEndpointUrl = function(id) {
            endpoint.ID = id;
            LocalStorage.storeEndpointID(id);
        };

        return service;
    }]);
