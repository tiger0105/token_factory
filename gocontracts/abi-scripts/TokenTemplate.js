const TokenTemplate = require("../../build/contracts/TokenTemplate.json"),
      fs = require("fs");

module.exports = function(done) {
  fs.writeFile("TokenTemplate.abi", JSON.stringify(TokenTemplate.abi), err => {
      if (err) {
          console.error(err);
      } else { 
          console.log("Token Template ABI created");
          done();
      }
  });
};