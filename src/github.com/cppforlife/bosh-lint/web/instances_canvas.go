package web

const instancesCanvas string = `
<script type="text/javascript">

function InstancesCanvas($el, canvasRouter) {
  var table = null;
  var loadedDeployment = null;

  function setUp() {
    Canvas($el, function() { table.Load({"deployment": loadedDeployment}); });
    table = InstancesTable(newDiv($el));
    canvasRouter.Apply($el);
  }

  setUp();

  return {
    Load: function(deployment) {
      loadedDeployment = deployment;
      table.Load({"deployment": deployment});
    },
  };
}

</script>
`
