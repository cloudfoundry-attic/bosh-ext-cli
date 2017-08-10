package web

const taskOutputForm string = `
<div id="task-output-form-tmpl" class="tmpl">
  <form autocomplete="off">
    <input type="text" name="id" placeholder="id" />
    <button type="submit">Go</button>
  </form>
</div>

<script type="text/javascript">

function TaskOutputForm($el, showCallback) {
  function setUp() {
    $el.html($("#task-output-form-tmpl").html());

    $el.find("form").submit(function(event) {
      event.preventDefault();
      showCallback($el.find("form input").val());
    });

    $el.find("form input").each(function() {
      NewSearchCriteriaExpandingInput($(this));
    });
  }

  setUp();

  return {
    Set: function(id) { $el.find("form input").val(id); },
    Focus: function() { $el.find("form input").focus(); },
  };
}

</script>
`
