package web

const taskOutputCanvas string = `
<script type="text/javascript">

function TaskOutputCanvas($el, canvasRouter) {
  var form = null;
  var table = null;
  var loadedID = null;

  function setUp() {
    Canvas($el);

    form = TaskOutputForm(newDiv($el), function(id) {
      canvasRouter.NewTaskOutputCanvas(id);
      form.Set(loadedID);
    });

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
      form.Set(id);
      form.Focus();
      table.Load(id);
    },
  };
}

</script>
`
