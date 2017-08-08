package web

const Layout string = `
<!DOCTYPE html>
<html>
<head>
  <script
  src="http://code.jquery.com/jquery-3.2.1.min.js"
  integrity="sha256-hwg4gsxgFZhOsEEamdOYGBf13FyQuiTwlAQgxVSNgt4="
  crossorigin="anonymous"></script>
</head>
<body>
` +
	misc +
	inputs +
	table +
	canvas +
	canvasRouter +
	deploymentsTable +
	deploymentsCanvas +
	instancesTable +
	instancesCanvas +
	tasksTable +
	tasksCanvas +
	taskOutputTable +
	taskOutputCanvas +
	eventsSearchCriteria +
	eventsSearchForm +
	eventsTable +
	eventsCanvas +
	main +
	`
</body>
</html>
`
