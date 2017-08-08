package web

const instancesCanvas string = `
<script type="text/javascript">

function InstancesCanvas($el, canvasRouter) {
  var table = null;

  function setUp() {
    Canvas($el);
    table = InstancesTable(newDiv($el));
    canvasRouter.Apply($el);
  }

  setUp();

  return {
    Load: function(deployment) { table.Load({"deployment": deployment}); },
  };
}

</script>
`
