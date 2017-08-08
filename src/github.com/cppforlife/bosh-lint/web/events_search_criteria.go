package web

const eventsSearchCriteria string = `
<script type="text/javascript">

function EventsSearchCriteria($el) {
  var data = {};
  var focusedInputName = null;

  var keys = ["after", "before", "event-user", "action",
    "object-type", "object-name", "task", "instance", "deployment"];

  function setUp() {
    if ($el) { // or empty critiria
      keys.forEach(function(key) {
        data[key] = $el.find("input[name='"+key+"']").val();
      });

      var $focused = $el.find("input:focus");
      if ($focused.length > 0) {
        focusedInputName = $focused.attr("name");
      }
    }
  }

  function AsQuery() { return data; }

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

  function CopyFrom(criteria2) {
    var query2 = criteria2.AsQuery();
    keys.forEach(function(key) {
      data[key] = query2[key];
    });
  }

  function SetKV(key, value) {
    data[key] = value;
    focusedInputName = key;
  }

  setUp();

  return {
    AsQuery: AsQuery,
    ApplyToForm: ApplyToForm,
    ApplyFocusToForm: ApplyFocusToForm,
    CopyFrom: CopyFrom,
    SetKV: SetKV,
  }
}

</script>
`
