package web

const taskOutputTable string = `
<script type="text/javascript">

function TaskOutputTable($el) {
  var dataSource = null;

  function setUp() {
    var moreCallback = function() { dataSource.More(); }
    var tmpls = {
      table: '<div></div>',
      empty: Tmpl('No matching task ouptut', []),
      error: Tmpl('Error fetching task output', []),
      dataItem: TaskOutputLineTmpl(), // Tmpl1('<pre>{_}</pre>'),
    };
    var table = Table($el, moreCallback, tmpls);
    dataSource = TableDataSource("task", table, null);
  }

  setUp();

  return {
    Load: function(id) { dataSource.Load({"id": id}); }
  };
}

function TaskOutputLineTmpl() {
  // Example: "Updating instance zookeeper: zookeeper/d3bfdf1e-57f0-4b1c-8f3e-c11a6759ee24 (0) (canary)"
  // Example: "Creating missing vms: zookeeper/d3bfdf1e-57f0-4b1c-8f3e-c11a6759ee24 (0)"
  // todo catches jobs and packages
  var instanceReg = /^(.+)\s+(\S+?\/\S+?)\s+(.+)$/;
  var line = 0;

  return {
    Render: function(data) {
      if ((line++)==0) {
        data = data.replace(/^\s+/, "");
      }
      var pieces = data.split("\n");
      for (var i=0; i<pieces.length; i++) {
        if (instanceReg.test(pieces[i])) {
          pieces[i] = pieces[i].replace(instanceReg, 
            "$1 <a href='#' data-query='instance' data-value='$2'>$2</a> $3")
        }
      }
      return pieces.join("<br/>");
    }
  };
}

</script>
`
