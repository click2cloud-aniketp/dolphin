angular
.module('dockm')
.directive('rdHeader', function rdHeader() {
  var directive = {
    scope: {
      'ngModel': '='
    },
    transclude: true,
    /*template: '<div class="row header"><div class="col-xs-12"><div class="meta" ng-transclude></div></div></div>',*/
    template: '<div class="topBar"><div class="container" ng-transclude></div></div>',
    restrict: 'EA'
  };
  return directive;
});
