package web

const eventsSearchForm string = `
<div id="events-search-form-tmpl" class="tmpl">
  <form autocomplete="off">
    <input type="text" name="after" placeholder="after" />
    <input type="text" name="before" placeholder="before" />
    <input type="text" name="event-user" placeholder="user" />
    <input type="text" name="action" placeholder="action" />
    <input type="text" name="object-type" placeholder="obj. type" />
    <input type="text" name="object-name" placeholder="obj. name" />
    <input type="text" name="task" placeholder="task" />
    <input type="text" name="deployment" placeholder="deployment" />
    <input type="text" name="instance" placeholder="instance" />
    <button type="submit">Go</button>
  </form>
</div>

<script type="text/javascript">

function EventsSearchForm($el, searchCallback) {
  function setUp() {
    $el.html($("#events-search-form-tmpl").html());

    $el.find("form").submit(function(event) {
      event.preventDefault();
      searchCallback();
    });

    $el.find("form input").each(function() {
      NewSearchCriteriaExpandingInput($(this));
      NewSearchCriteriaClearButton($(this));
    });
  }

  setUp();

  return {
    Criteria: function() {
      return EventsSearchCriteria($el.find("form"));
    },
    SetCriteria: function(criteria) {
      criteria.ApplyToForm($el.find("form"));
    },
    SetFocus: function(criteria) {
      criteria.ApplyFocusToForm($el.find("form"));
    }
  };
}

</script>
`
