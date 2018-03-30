package web

const canvasRouter string = `
<script type="text/javascript">

function CanvasRouter(collection) {
  function buildRouter(customizeCriteriaFunc) {
    return function(key, val) {
      if (key == "instances-canvas") {
        collection.NewInstancesCanvas(val);
      }
      else if (key == "task-output-canvas") {
        collection.NewTaskOutputCanvas(val);
      }
      else {
        var criteria = EventsSearchCriteria();
        if (customizeCriteriaFunc) {
          customizeCriteriaFunc(criteria);
        }
        criteria.SetKV(key, val);
        var canvas = collection.NewEventsCanvas();
        canvas.Search(criteria);
      }
    }
  }

  return {
    Apply: function($el) {
      FilteringKVs($el, buildRouter(null));
    },

    ApplyWithCustomEvents: function($el, criteriaFunc) {
      FilteringKVs($el, buildRouter(criteriaFunc));
    },

    NewEventsCanvas: function(criteria) {
      var canvas = collection.NewEventsCanvas();
      canvas.Search(criteria);
    },

    NewTaskOutputCanvas: function(id) {
      collection.NewTaskOutputCanvas(id);
    }
  }
}

function FilteringKVs($el, clickCallback) {
  $el.on("click", "a[data-query]", function(event) {
    event.preventDefault();
    var query = $(event.target).data("query");
    var val = $(event.target).data("value");
    clickCallback(query, val);
  });
}

</script>
`
