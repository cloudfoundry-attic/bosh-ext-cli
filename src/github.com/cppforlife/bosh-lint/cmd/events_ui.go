package cmd

/* todo
- pagination
- style?
- api endpoints for:
	- t 8 --json

*/

const eventsUI = `
<!DOCTYPE html>
<html>
<head>
	<script
  src="http://code.jquery.com/jquery-3.2.1.min.js"
  integrity="sha256-hwg4gsxgFZhOsEEamdOYGBf13FyQuiTwlAQgxVSNgt4="
  crossorigin="anonymous"></script>
</head>
<body>

<div id="root"></div>

<div id="canvas-tmpl" class="tmpl">
	<form autocomplete="off">
	  <input type="text" name="after" placeholder="after" />
		<input type="text" name="before" placeholder="before" />
		<input type="text" name="event-user" placeholder="user" />
		<input type="text" name="action" placeholder="action" />
		<input type="text" name="object-type" placeholder="object-type" />
		<input type="text" name="object-name" placeholder="object-name" />
		<input type="text" name="task" placeholder="task" />
		<input type="text" name="deployment" placeholder="deployment" />
		<input type="text" name="instance" placeholder="instance" />
		<button type="submit">Submit</button>
	</form>
	<table></table>
</div>

<table id="event-tmpl" class="tmpl">
	<tr>
		<td class="id">{id}</td>
		<td class="time">{time}</td>
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
		</td>
		<td class="deployment">
			<a href="#" data-query="deployment" data-value="{deployment}">{deployment}</a>
		</td>
		<td class="instance">
			<a href="#" data-query="instance" data-value="{instance}">{instance}</a>
		</td>
		<td class="context"><span>{context}</span></td>
		<td class="error"><span>{error}</span></td>
	<tr>
</table>

<table id="no-events-tmpl" class="tmpl">
	<tr>
		<td colspan="11">No matching events</td>
	<tr>
</table>

<script type="text/javascript">

function CanvasCollection($el) {
	var $canvases = null;

	function setUp() {
		$canvases = newNamedDiv($el, "canvases")
		NewCanvasAddButton(newNamedDivPrepended($el, "add-button"), NewCanvas);
	}

	function NewCanvas() {
		return new Canvas(newNamedDivPrepended($canvases, "canvas"), searchCallback);
	}

	function searchCallback(canvas) {
		var canvas2 = NewCanvas();
		canvas2.Search(canvas.SearchCriteria());
		canvas.ResetCriteria();
	}

	setUp();

	return {
		NewCanvas: NewCanvas
	};
}

function NewCanvasAddButton($el, clickCallback) {
	function setUp() {
		$el.html("<button>+</button>").find("button").click(clickCallback);
	}

	setUp();

	return {};
}

function NewCanvasDeleteButton($el, clickCallback) {
	function setUp() {
		$el.html("<button>-</button>").find("button").click(clickCallback);
	}

	setUp();

	return {};
}

function Canvas($el, searchCallback) {
	var obj = {};
	var currCriteria = new EmptySearchCriteria();

	function setUp() {
		$el.html($("#canvas-tmpl").html());

		$el.find("form").submit(function search(event) {
			event.preventDefault();
			searchCallback(obj);
		});

		$el.on("click", "a[data-query]", function(event) {
			event.preventDefault();
      // todo represent as a object
			var query = $(event.target).data("query");
			var val = $(event.target).data("value");
			$el.find("form").find("input[name='"+query+"']").val(val).focus();
			searchCallback(obj);
		});

		NewCanvasDeleteButton(newNamedDivPrepended($el, "delete-button"), function() {
			$el.remove();
		});
	}

	obj.SearchCriteria = function() {
		return new SearchCriteria($el.find("form"));
	};

	obj.Search = function(criteria) {
		criteria.ApplyToForm($el.find("form"));
		criteria.ApplyFocusToForm($el.find("form"));
		currCriteria = criteria;

		$.post("/api/events", criteria.AsQuery())
			.done(addEvents)
			.fail(function() {
		    console.log("error");
		  });
	};

	obj.ResetCriteria = function() {
		currCriteria.ApplyToForm($el.find("form"));
	};

	var eventHtml = $('#event-tmpl').html();
	var eventKeys = ["action", "context", "deployment", "error", "id",
		"instance", "object_name", "object_type", "task_id", "time", "user"];

	function addEvents(data) {
		var eventsHtml = '';

		if (data.Tables[0].Rows.length == 0) {
			eventsHtml = $("#no-events-tmpl").html();
		} else {
			data.Tables[0].Rows.forEach(function(apiEvent) {
				eventsHtml += buildEventTmpl(apiEvent);
			});
		}

		$el.find('table').html(eventsHtml);
	}

	function buildEventTmpl(apiEvent) {
		var eventHtml2 = eventHtml;
		eventKeys.forEach(function(key) {
			eventHtml2 = eventHtml2.replace(new RegExp('{' + key + '}', 'g'), apiEvent[key]);
		});
		return eventHtml2;
	}

	setUp();

	return obj;
}

function EmptySearchCriteria() {
	return {
		AsQuery: function() { return "" },
		ApplyToForm: function($el) { $el[0].reset(); },
		ApplyFocusToForm: function($el) {},
	}
}

function SearchCriteria($el) {
	var data = {};
	var serializedData = "";
	var focusedInputName = null;

	var keys = ["after", "before", "event-user", "action",
		"object-type", "object-name", "task", "instance", "deployment"];

	function setUp() {
		keys.forEach(function(key) {
			data[key] = $el.find("input[name='"+key+"']").val();
		});

		serializedData = $el.serialize();

		var $focused = $el.find("input:focus");
		if ($focused.length > 0) {
			focusedInputName = $focused.attr("name");
		}
	}

	function AsQuery() {
		return serializedData;
	}

	function ApplyToForm($el2) {
		Object.keys(data).forEach(function(key) {
			$el2.find("input[name='"+key+"']").val(data[key]);
		});
	}

	function ApplyFocusToForm($el2) {
		if (focusedInputName) {
			$el2.find("input[name='"+focusedInputName+"']").focus();
		}
	}

	setUp();

	return {
		AsQuery: AsQuery,
		ApplyToForm: ApplyToForm,
		ApplyFocusToForm: ApplyFocusToForm,
	}
}

function newNamedDiv($el, className) {
	return $el.append("<div class='"+className+"'></div>").find("div:last")
}

function newNamedDivPrepended($el, className) {
	return $el.prepend("<div class='"+className+"'></div>").find("div:first")
}

function main() {
  var collection = new CanvasCollection(newNamedDiv($("#root"), "canvas-collection"));

  // start by default with new canvas with all results
  var firstCanvas = collection.NewCanvas();
  firstCanvas.Search(new EmptySearchCriteria());
}

window.addEventListener("load", function load(event){
  window.removeEventListener("load", load, false);
  main();
}, false);

</script>

<style>
.tmpl { display: none; }
button { cursor: pointer; }
form { margin-bottom: 10px; }
input[type="text"], button { font-size: 18px; }
input::placeholder { color: #ccc; }
input[name="action"],
input[name="event-user"],
input[name="object-type"],
input[name="task"] { width: 70px; }
.delete-button { float: right; }
.canvas {
	margin-top: 10px;
	padding-top: 10px;
	border-top: 2px solid #efefef;
}
table {
  border-spacing: 0;
  border-collapse: collapse;
}
td {
	border: 1px solid #f1f1f1;
	vertical-align: top;
	padding: 0 5px;
}
td.time { width: 230px; }
td.context, td.error { width: 30px; }
td.context span,
td.error span {
	display: inline-block;
	width: 20px;
	white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}
</style>

</body>
</html>
`
