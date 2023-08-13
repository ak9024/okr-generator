const fs = require("fs");

fs.writeFile(
  ".env",
  `REACT_APP_BACKEND=${process.env.BACKEND}`,
  function(err) {
    if (err) throw err;
    console.log(`Generate env`);
  },
);
