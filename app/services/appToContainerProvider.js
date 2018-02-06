angular.module('dockm.services')
    .factory('AppToContainerProvider', ['LocalStorage', function AppToContainerProviderFactory(LocalStorage) {
        'use strict';
        var service = {};
        var appToContainerProvider = {};

        service.initialize = function() {
            var BaseImage = LocalStorage.getBaseImage();
            var GitUrl = LocalStorage.getGitUrl();
            var ImageName = LocalStorage.getImageName();
            var EndPointId = LocalStorage.getEndPointId();
            var EndPointUrl = LocalStorage.getEndPointUrl();
            if (BaseImage) {
                appToContainerProvider.BaseImage = BaseImage;
            }
            if (GitUrl) {
                appToContainerProvider.GitUrl = GitUrl;
            }
            if (ImageName) {
                appToContainerProvider.ImageName = ImageName;
            }
            if (EndPointId) {
                appToContainerProvider.EndPointId = EndPointId;
            }
            if (EndPointUrl) {
                appToContainerProvider.EndPointUrl = EndPointUrl;
            }
        };

        service.clean = function() {
            appToContainerProvider = {};
        };

        service.BaseImage = function() {
            return appToContainerProvider.BaseImage;
        };

        service.setBaseImage = function(baseimage) {
            appToContainerProvider.BaseImage = baseimage;
            LocalStorage.storeEndpointID(baseimage);
        };

        service.GitUrl = function() {
            return appToContainerProvider.GitUrl;
        };

        service.setGitUrl = function(gitURL) {
            appToContainerProvider.GitUrl = gitURL;
            LocalStorage.storeEndpointPublicURL(gitURL);
        };
        service.ImageName = function() {
            return appToContainerProvider.ImageName;
        };

        service.setImageName = function(imagename) {
            appToContainerProvider.ImageName = imagename;
            LocalStorage.storeEndpointID(imagename);
        };
        service.endpointID = function() {
            return endpoint.ID;
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
        // service.Output = function() {
        //     return Output;
        // };
        //
        // service.setOutput = function(output) {
        //     appToContainer.Output = output;
        //     LocalStorage.storeEndpointID(output);
        // };

        return service;
    }]);
