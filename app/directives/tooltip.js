angular
.module('dockm')
.directive('dockmTooltip', [function dockmTooltip() {
  var directive = {
    scope: {
      message: '@',
      position: '@'
    },
    template: '<span class="interactive" tooltip-append-to-body="true" tooltip-placement="{{position}}" tooltip-class="dockm-tooltip" uib-tooltip="{{message}}"><i class="fa fa-question-circle tooltip-icon" aria-hidden="true"></i></span>',
    restrict: 'E'
  };
  return directive;
}]);
