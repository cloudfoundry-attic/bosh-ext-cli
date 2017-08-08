package web

const taskOutputCanvas string = `
<script type="text/javascript">

function TaskOutputCanvas($el, canvasRouter) {
  var table = null;
  var loadedID = null;

  function setUp() {
    Canvas($el);

    table = TaskOutputTable(newDiv($el));

    canvasRouter.ApplyWithCustomEvents($el, function(criteria) {
      if (loadedID) {
        criteria.SetKV("task", loadedID);
      }
    });
  }

  setUp();

  return {
    Load: function(id) {
      loadedID = id;
      table.Load(id);
    },
  };
}

</script>
`
