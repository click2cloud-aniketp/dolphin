angular
.module('dockm')
.directive('rdHeaderContent', ['Authentication', function rdHeaderContent(Authentication) {
  var directive = {
    requires: '^rdHeader',
    transclude: true,
    link: function (scope, iElement, iAttrs) {
      scope.username = Authentication.getUserDetails().username;
    },
    /*template: '<div class="breadcrumb-links"><div class="pull-left" ng-transclude></div><div class="pull-right" ng-if="username"><a ui-sref="userSettings" class="text-my-account" style="margin-right: 5px;"><u><i class="fa fa-wrench" aria-hidden="true"></i> my account </u></a><a ui-sref="auth({logout: true})" class="text-logout" style="margin-right: 25px;"><u><i class="fa fa-sign-out" aria-hidden="true"></i> log out</u></a></div></div>',*/
    template: '<div class="row">' +
    '            <div class="col-md-12">' +
    '                <b class="breadcrumb-links" ng-transclude></b>' +
    '            </div>' +
    '        </div>',
    restrict: 'E'
  };
  return directive;
}]);
