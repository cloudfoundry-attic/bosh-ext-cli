package web

const eventsCanvas string = `
<script type="text/javascript">

function EventsCanvas($el, canvasRouter) {
  var obj = {};
  var form = null;
  var table = null;
  var currCriteria = new EventsSearchCriteria();

  function setUp() {
    Canvas($el);

    form = EventsSearchForm(newDiv($el), function() {
      canvasRouter.NewEventsCanvas(form.Criteria());
      form.SetCriteria(currCriteria);
    });

    table = EventsTable(newDiv($el));

    canvasRouter.ApplyWithCustomEvents($el, function(criteria) {
      criteria.CopyFrom(form.Criteria());
    });
  }

  function search(criteria) {
    form.SetCriteria(criteria);
    form.SetFocus(criteria);
    table.Load(criteria);
    currCriteria = criteria;
  }

  setUp();

  obj.SearchCriteria = function() { return form.Criteria(); };
  obj.Search = search;

  return obj;
}

</script>
`
