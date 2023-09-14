const express = require("express"),
	app = express();

const port = 5000;

const answer = [];

app.get("/fizz-buz-server/run", (req, res) => {
	let num = req.query.num;
	if (num % 3 === 0 && num % 5 === 0) {
		answer[num] = "FizzBuzz";
	} else if (num % 3 === 0) {
		answer[num] = "Fizz";
	} else if (num % 5 === 0) {
		answer[num] = "Buzz";
	} else {
		answer[num] = num;
	}
	res.send(JSON.stringify(answer, null, 2));
});

app.listen(port, () => console.log(`Server is lintening on port ${port}`));
