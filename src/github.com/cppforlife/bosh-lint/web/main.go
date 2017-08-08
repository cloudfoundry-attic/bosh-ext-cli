package web

const main string = `
<div id="root"></div>

<script type="text/javascript">

function CanvasCollection($el) {
  var obj = {};
  var $canvases = null;
  var router = null;

  function setUp() {
    CanvasAddButton(newDiv($el), "deployments", NewDeploymentsCanvas);
    CanvasAddButton(newDiv($el), "tasks", NewTasksCanvas);
    CanvasAddButton(newDiv($el), "events", NewEventsCanvas);

    $canvases = newDiv($el);
    router = CanvasRouter(obj);
  }

  function NewDeploymentsCanvas() {
    var canvas = DeploymentsCanvas(newDivPrepended($canvases), router);
    canvas.Load();
    return canvas;
  }

  function NewInstancesCanvas(deployment) {
    var canvas = InstancesCanvas(newDivPrepended($canvases), router);
    canvas.Load(deployment);
    return canvas;
  }

  function NewTasksCanvas() {
    var canvas = TasksCanvas(newDivPrepended($canvases), router);
    canvas.Load();
    return canvas;
  }

  function NewTaskOutputCanvas(id) {
    var canvas = TaskOutputCanvas(newDivPrepended($canvases), router);
    canvas.Load(id);
    return canvas;
  }

  function NewEventsCanvas() {
    return EventsCanvas(newDivPrepended($canvases), router);
  }

  setUp();

  obj.NewEventsCanvas = NewEventsCanvas;
  obj.NewTasksCanvas = NewTasksCanvas;
  obj.NewTaskOutputCanvas = NewTaskOutputCanvas;
  obj.NewInstancesCanvas = NewInstancesCanvas;

  return obj;
}

function main() {
  var collection = CanvasCollection(newDiv($("#root")));

  // start by default with new canvas with all results
  var firstCanvas = collection.NewEventsCanvas();
  firstCanvas.Search(EventsSearchCriteria());
}

window.addEventListener("load", function load(event){
  window.removeEventListener("load", load, false);
  main();
}, false);

</script>

<style>
.canvas-add-button,
form,
.canvas table,
.table-more-button { margin-bottom: 10px; }
</style>
`
