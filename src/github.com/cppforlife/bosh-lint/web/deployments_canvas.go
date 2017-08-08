package web

const deploymentsCanvas string = `
<script type="text/javascript">

function DeploymentsCanvas($el, canvasRouter) {
  var table = null;

  function setUp() {
    Canvas($el);
    table = DeploymentsTable(newDiv($el));
    canvasRouter.Apply($el);
  }

  setUp();

  return {
    Load: function() { table.Load(); },
  };
}

</script>
`
