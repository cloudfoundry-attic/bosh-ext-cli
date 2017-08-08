package web

const table string = `
<script type="text/javascript">

function TableDataSource(command, table, lastItemFunc) {
  var path = "/api/command";
  var lastQuery = null;
  var lastDataItem = null;

  function noteData(data) {
    if (data.Tables && data.Tables.length > 0) {
      var rows = data.Tables[0].Rows || [];
      if (rows.length > 0) {
        lastDataItem = rows[rows.length-1];  
        return rows;
      }
    }

    var blocks = data.Blocks || [];
    if (blocks.length > 0) {
      return blocks;
    }

    return [];
  }

  function hasMoreData() {
    return lastItemFunc && lastDataItem;
  }

  return {
    Load: function(query) {
      lastQuery = query;
      var query_ = $.extend({"command": command}, query);

      $.post(path, query_)
        .done(function(data) { table.SetData(noteData(data), hasMoreData()); })
        .fail(table.ShowError);
    },

    More: function() {
      if (hasMoreData()) {
        var moreQuery = $.extend({"command": command}, lastQuery, lastItemFunc(lastDataItem));

        $.post(path, moreQuery)
          .done(function(data) { table.AddData(noteData(data)); })
          .fail(table.ShowError);  
      } else {
        table.AddData([]);
      }
    },
  };
}

function Table($el, moreCallback, tmpls) {
  var $table = null;
  var moreButton = null;

  function setUp() {
    $table = $(tmpls.table || '<table></table>').appendTo($el);
    moreButton = MoreButton(newDiv($el), moreCallback);
  }

  function setData(data, hasMoreData) {
    if (data.length == 0) {
      $table.html(tmpls.empty.Render());
      moreButton.Hide();
    } else {
      addData(data);
    }
    if (!hasMoreData) {
      moreButton.Hide();
    }
  }

  function addData(data) {
    if (data.length == 0) {
      moreButton.Hide();
    } else {
      var html = '';
      data.forEach(function(apiEvent) {
        html += tmpls.dataItem.Render(apiEvent);
      });
      $table.append(html);
      moreButton.Show();
    }
  }

  function showError() {
    $table.append(tmpls.error.Render());
    moreButton.Hide();
  }

  setUp();

  return {
    SetData: setData,
    AddData: addData,
    ShowError: showError,
  };
}

function MoreButton($el, clickCallback) {
  var $button = null;

  function setUp() {
    $el.addClass("table-more-button");
    $button = $el.html("<button>More...</button>").find("button");
    $button.click(clickCallback);
    $button.hide(); // default to hide
  }

  setUp();

  return {
    Show: function() { $button.show(); },
    Hide: function() { $button.hide(); },
  };
}

</script>

<style>
table {
  border-spacing: 0;
  border-collapse: collapse;
}
td {
  border: 1px solid #f1f1f1;
  vertical-align: top;
  padding: 0 5px;
}
table tr.hover { background: yellow; }
</style>
`
