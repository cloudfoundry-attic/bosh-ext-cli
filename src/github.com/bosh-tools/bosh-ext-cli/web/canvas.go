package web

const canvas string = `
<script type="text/javascript">

function Canvas($el, reloadCallback) {
  function setUp() {
    $el.addClass("canvas");

    var reloadInterval = null;

    CanvasDeleteButton(newDivPrepended($el), function() {
      $el.remove();
      window.clearInterval(reloadInterval);
    });

    if (reloadCallback) {
      CanvasReloadButton(newDivPrepended($el), reloadCallback);  
      reloadInterval = setInterval(reloadCallback, 5000);
    }
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
    $el.html("<button title='remove'>-</button>").find("button").click(clickCallback);
  }
  setUp();
  return {};
}

function CanvasReloadButton($el, clickCallback) {
  function setUp() {
    $el.addClass("canvas-reload-button");
    $el.html("<button title='reload'>c</button>").find("button").click(clickCallback);
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
.canvas-delete-button,
.canvas-reload-button {
  margin-left: 10px;
  float: right;
}
</style>
`
