package web

const deploymentsTable string = `
<table id="deployment-tmpl" class="tmpl">
  <tr>
    <td class="deployment">
      <a href="#" data-query="deployment" data-value="{name}">{name}</a>
      <a href="#" data-query="instances-canvas" data-value="{name}">...</a>
    </td>
  </tr>
</table>

<script type="text/javascript">

function DeploymentsTable($el) {
  var dataSource = null;

  function setUp() {
    var moreCallback = function() { dataSource.More(); }
    var tmpls = {
      empty: Tmpl('<tr><td colspan="1">No deployments</td></tr>', []),
      error: Tmpl('<tr><td colspan="1">Error fetching deployments</td></tr>', []),
      dataItem: Tmpl($("#deployment-tmpl").html(), ["name"]),
    };
    var table = Table($el, moreCallback, tmpls);
    dataSource = TableDataSource("deployments", table, null);
  }

  setUp();

  return {
    Load: function() { dataSource.Load(); }
  };
}

</script>
`
