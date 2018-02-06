angular.module('dockm.helpers')
.factory('UserHelper', [function UserHelperFactory() {
  'use strict';
  var helper = {};

  helper.filterNonAdministratorUsers = function(users) {
    return users.filter(function (user) {
      if (user.Role !== 1) {
        return user;
      }
    });
  };

  /*Added by sapan to filter out org users*/ 
  helper.filterOrgUsers = function(users) {
    return users.filter(function (user) {
      if (user.Username.indexOf('click2clouduser') == -1) {
        return user;
      }
    });
  };
  
  return helper;
}]);
