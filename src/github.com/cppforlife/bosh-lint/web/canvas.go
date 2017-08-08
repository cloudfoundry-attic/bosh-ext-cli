package web

const canvas string = `
<script type="text/javascript">

function Canvas($el) {
  function setUp() {
    $el.addClass("canvas");
    CanvasDeleteButton(newDivPrepended($el), function() { $el.remove(); });
  }
  setUp();
  return {};
}

function CanvasAddButton($el, title, clickCallback) {
  function setUp() {
    $el.addClass("canvas-add-button");
    $el.html("<button>+ "+title+"</button>").find("button").click(clickCallback);
  }
  setUp();
  return {};
}

function CanvasDeleteButton($el, clickCallback) {
  function setUp() {
    $el.addClass("canvas-delete-button");
    $el.html("<button>-</button>").find("button").click(clickCallback);
  }
  setUp();
  return {};
}

</script>

<style>
.canvas {
  padding-top: 10px;
  border-top: 2px solid #efefef;
}
.canvas-add-button { 
  display: inline-block;
  margin-right: 10px;
}
.canvas-delete-button { float: right; }
</style>
`
