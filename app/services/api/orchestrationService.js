angular.module('dockm.services')
    .factory('orchestrationService', ['$q', 'Orchestration', 'FileUploadService','LocalStorage', function orchestrationServiceFactory($q, orchestration, FileUploadService, LocalStorage) {
        'use strict';
        var service = {};
        service.orchestration = function () {
            var otos ={
                EndPointId:''
            };
            otos.EndPointId = LocalStorage.getEndpointID()

            return orchestration.create({}, otos).$promise;
        };
        return service;
    }]);
