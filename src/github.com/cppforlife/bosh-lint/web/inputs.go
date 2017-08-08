package web

const inputs string = `
<script type="text/javascript">

function NewSearchCriteriaExpandingInput($input) {
  function setUp() {
    $input
      .focus(function() { $(this).addClass("expanded"); })
      .blur(function() {
        var $el = $(this);
        setTimeout(function() { $el.removeClass("expanded"); }, 100);
      });
  }

  setUp();

  return {};
}

function NewSearchCriteriaClearButton($input) {
  function setUp() {
    var $button = $("<a class='search-criteria-clear-button'>x</a>").click(function(event) {
      event.preventDefault();
      $input.val("");
      $input.focus();
    });
    $input.after($button);
  }

  setUp();

  return {};
}

</script>

<style>
.canvas input { width: 100px; }
.canvas input.expanded { width: 300px; }

.search-criteria-clear-button {
  background: none;
  border: none;
  vertical-align: top;
  padding: 5px;
  font-size: 12px;
  font-family: system-ui;
  cursor: pointer;
}
</style>
`
