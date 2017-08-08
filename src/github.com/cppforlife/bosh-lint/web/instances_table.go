package web

const instancesTable string = `
<table id="instance-tmpl" class="tmpl">
  <tr class="instance-table-row">
    <td class="instance">
      <a href="#" data-query="instance" data-value="{instance}">{instance}</a>
    </td>
    <td class="process_state">{process_state} [{state}]</td>
    <td class="az">{az}</td>
    <td class="ips">{ips}</td>
    <!-- td class="index">{index}</td -->
    <td class="vm_type">{vm_type}</td>
    <td class="vm_cid">
      <a href="#" data-query="object-name" data-value="{vm_cid}">{vm_cid}</a>
    </td>
    <td class="disk_cids">{disk_cids}</td>
    <!-- td class="agent_id">{agent_id}</td -->
    <!-- td class="resurrection_paused">{resurrection_paused}</td -->
    <!-- td class="ignore">{ignore}</td -->
    <!-- td class="bootstrap">{bootstrap}</td -->
  </tr>
</table>

<script type="text/javascript">

function InstancesTable($el) {
  var dataSource = null;

  function setUp() {
    var moreCallback = function() { dataSource.More(); }
    var tmpls = {
      empty: Tmpl('<tr><td colspan="7">No matching instances</td></tr>', []),
      error: Tmpl('<tr><td colspan="7">Error fetching instances</td></tr>', []),
      dataItem: Tmpl($("#instance-tmpl").html(), 
        ["agent_id", "az", "bootstrap", "disk_cids", "ignore", "index", 
        "instance", "ips", "process_state", "resurrection_paused", "state", 
        "vm_cid", "vm_type"]),
    };

    var table = Table($el, moreCallback, tmpls);
    dataSource = TableDataSource("instances", table, null);
  }

  setUp();

  return {
    Load: function(query) { dataSource.Load($.extend({"details": ""}, query)); }
  };
}

</script>
`
