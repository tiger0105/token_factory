const TokenFactory = require("../../build/contracts/TokenFactory.json"),
      fs = require("fs");

module.exports = function(done) {
  fs.writeFile("TokenFactory.abi", JSON.stringify(TokenFactory.abi), err => {
      if (err) {
          console.error(err);
      } else { 
          console.log("Token Factory ABI created");
          done();
      }
  });
};