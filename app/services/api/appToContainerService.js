/*eslint linebreak-style: [2, "windows"]*/
angular.module('dockm.services')
.factory('appToContainerService', ['$q', 'AppToContainer', 'FileUploadService','LocalStorage', function appToContainerServiceFactory($q, appToContainer, FileUploadService, LocalStorage) {
  'use strict';
    var service = {};
    service.appToContainer = function (BaseImage, GitUrl, ImageName) {
        var atoc ={
            BaseImage: BaseImage,
            GitUrl:GitUrl,
            ImageName:ImageName,
            EndPointId:''
            //EndPointUrl:''
        };
        atoc.EndPointId = LocalStorage.getEndpointID()
        //atoc.EndPointUrl = LocalStorage.getEndpointPublicURL()
        return appToContainer.create({}, atoc).$promise;
    };
    // service.appToContainerOutput = function() {
    //
    //     return appToContainer.get({}).$promise;
    // };


  return service;
}]);
