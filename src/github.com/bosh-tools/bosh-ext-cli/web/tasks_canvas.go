package web

const tasksCanvas string = `
<script type="text/javascript">

function TasksCanvas($el, canvasRouter) {
  var table = null;

  function setUp() {
    Canvas($el, function() { table.Load(); });
    table = TasksTable(newDiv($el));
    canvasRouter.Apply($el);
  }

  setUp();

  return {
    Load: function() { table.Load(); },
  };
}

</script>
`
