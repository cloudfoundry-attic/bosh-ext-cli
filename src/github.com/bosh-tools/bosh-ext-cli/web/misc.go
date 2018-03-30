package web

const misc string = `
<script type="text/javascript">

function newDiv($el) {
  return $el.append("<div></div>").find("div:last")
}

function newDivPrepended($el) {
  return $el.prepend("<div></div>").find("div:first")
}

function Tmpl(html, keys) {
  return {
    Render: function(data) {
      var html2 = html;
      keys.forEach(function(key) {
        html2 = html2.replace(new RegExp('{' + key + '}', 'g'), data[key]);
      });
      return html2;
    }
  };
}

function Tmpl1(html) {
  return {
    Render: function(data) {
      var html2 = html;
      return html2.replace(new RegExp('{_}', 'g'), data);
    }
  };
}

</script>

<style>
.tmpl { display: none; }
button { cursor: pointer; }
input[type="text"], button { font-size: 18px; }
input::placeholder { color: #ccc; }
</style>
`
