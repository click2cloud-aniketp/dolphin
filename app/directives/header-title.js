angular
.module('dockm')
.directive('rdHeaderTitle', ['Authentication', function rdHeaderTitle(Authentication) {
  var directive = {
    requires: '^rdHeader',
    scope: {
      title: '@'
    },
    link: function (scope, iElement, iAttrs ) {
      scope.username = Authentication.getUserDetails().username;
    },
    transclude: true,
    // template: '<div class="page white-space-normal">{{title}}<span class="header_title_content" ng-transclude></span><span class="pull-right user-box" ng-if="username"><i class="fa fa-user-circle" aria-hidden="true"></i> {{username}}</span></div>',
    template: '<div class="row">' +
    '            <div class="col-md-12">' +
    '                <b class="topTile">{{title}}</b> ' +
    '                <span class="header_title_content" ng-transclude></span> '+
    '                <a href="" title="Menu" class="hamburger animated fadeInUp delay fa fa-bars"> </a> ' +
    '                <span class="pull-right user-box" style="padding-right: 7%; padding-top: 2px" ng-if="username"><i class="fa fa-user-circle" aria-hidden="true"></i> {{username}}</span> ' +
    '            </div>' +
    '        </div>',
    restrict: 'E'
  };
  return directive;
}]);