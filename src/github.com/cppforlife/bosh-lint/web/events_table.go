package web

const eventsTable string = `
<table id="event-tmpl" class="tmpl">
  <tr data-event-id="{id}" class="event-table-row">
    <td class="id">{id}</td>
    <td class="time">
      <a href="#" data-query="after" data-value="{time}">{time}</a>
    </td>
    <td class="user">
      <a href="#" data-query="event-user" data-value="{user}">{user}</a>
    </td>
    <td class="action">
      <a href="#" data-query="action" data-value="{action}">{action}</a>
    </td>
    <td class="object_type">
      <a href="#" data-query="object-type" data-value="{object_type}">{object_type}</a>
    </td>
    <td class="object_name">
      <a href="#" data-query="object-name" data-value="{object_name}">{object_name}</a>
    </td>
    <td class="task_id">
      <a href="#" data-query="task" data-value="{task_id}">{task_id}</a>
      <a href="#" data-query="task-output-canvas" data-value="{task_id}">...</a>
    </td>
    <td class="deployment">
      <a href="#" data-query="deployment" data-value="{deployment}">{deployment}</a>
      <a href="#" data-query="instances-canvas" data-value="{deployment}">...</a>
    </td>
    <td class="instance">
      <a href="#" data-query="instance" data-value="{instance}">{instance}</a>
    </td>
  </tr>
</table>

<table id="event-context-tmpl" class="tmpl">
  <tr class="event-context-table-row">
    <td colspan="9"><pre>{_}</pre></td>
  </tr>
</table>

<table id="event-error-tmpl" class="tmpl">
  <tr class="event-error-table-row">
    <td colspan="9"><pre>{_}</pre></td>
  </tr>
</table>

<script type="text/javascript">

function EventsTable($el) {
  var dataSource = null;

  function setUp() {    
    var moreCallback = function() { dataSource.More(); }
    var tmpls = {
      empty: Tmpl('<tr><td colspan="9">No matching events</td></tr>', []),
      error: Tmpl('<tr><td colspan="9">Error fetching events</td></tr>', []),
      dataItem: EventTmpl(),
    };

    var table = Table($el, moreCallback, tmpls);
    dataSource = TableDataSource(
      "events", table, function(i) { return {"before-id": i.id}; });

    HoverableEvents($el);
  }

  setUp();

  return {
    Load: function(criteria) { dataSource.Load(criteria.AsQuery()); }
  };
}

function EventTmpl() {
  var regularTmpl = Tmpl($("#event-tmpl").html(), 
    ["action", "context", "deployment", "error", "id",
      "instance", "object_name", "object_type", "task_id", "time", "user"]);

  var contextTmpl = Tmpl1($("#event-context-tmpl").html());
  var errorTmpl = Tmpl1($("#event-error-tmpl").html());

  return {
    Render: function(data) {
      var result = regularTmpl.Render(data);
      if (data.context.length > 0) {
        result += contextTmpl.Render(data.context);
      }
      if (data.error.length > 0) {
        result += errorTmpl.Render(data.error);
      }
      return result;
    }
  };
}

function HoverableEvents($el) {
  var $selected = $el;
  var className = "hover";

  $el.on("mouseover", "table tr", function(event) {
    var $tr = $(event.target).parent("tr");
    if ($tr.length == 0) return;
    if (!$tr.data("event-id")) return;

    $selected.removeClass(className);
    $selected = $tr;

    // Example: "200 <- 199" hovering over 200
    var ids = String($tr.data("event-id")).split(" <- ");
    if (ids.length == 2) {
      $selected = $selected.add($el.find("tr[data-event-id='"+ids[1]+"']"));
    }

    $selected.addClass(className);
  });

  return {};
}

</script>

<style>
.event-table-row td.time { width: 230px; }
.event-table-row td.context,
.event-table-row td.error { width: 30px; }
.event-table-row td.context span,
.event-table-row td.error span {
  display: inline-block;
  width: 20px;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}
.event-context-table-row pre,
.event-error-table-row pre { margin: 7px 0; }
.event-error-table-row { background: pink; }
table tr.hover { background: yellow; }
</style>
`
