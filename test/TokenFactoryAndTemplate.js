const TokenFactory = artifacts.require("./TokenFactory.sol");
const TokenTemplate = artifacts.require("./TokenTemplate.sol");

contract("TokenFactory", async accounts => {
  const mintableToken = {
      name: "Mintable Token",
      symbol: "MNT",
      decimals: 18,
      logoURL:
        "https://apollo-uploads-las.s3.amazonaws.com/1442324623/atlanta-hawks-logo-944556.png",
      logoHash: web3.utils.toHex(
        "0x4D021B157A49F472A48AB02A1F2F6E2986C169A7C78CC94179EDAEBD5E96E8E4"
      ), // sha256 hash
      contractHash: web3.utils.toHex(
        "0x4D021B157A49F472A48AB02A1F2F6E2986C169A7C78CC94179EDAEBD5E96E8E4"
      ),
      supply: 0
    },
    cappedToken = {
      name: "Capped Token",
      symbol: "CAP",
      decimals: 18,
      logoURL:
        "https://apollo-uploads-las.s3.amazonaws.com/1442324623/atlanta-hawks-logo-944556.png",
      logoHash: web3.utils.toHex(
        "0x4D021B157A49F472A48AB02A1F2F6E2986C169A7C78CC94179EDAEBD5E96E8E4"
      ),
      contractHash: web3.utils.toHex(
        "0x4D021B157A49F472A48AB02A1F2F6E2986C169A7C78CC94179EDAEBD5E96E8E4"
      ),
      supply: 100
    };

  const owner = accounts[0];
  const otherUser = accounts[1];

  before(async function() {
    try {
      await web3.eth.personal.unlockAccount(owner, "", 10000);
      await web3.eth.personal.unlockAccount(otherUser, "", 10000);
    } catch(error) {
      console.warn(`error in unlocking wallet: ${JSON.stringify(error)}`);
    }
  });

  describe("Token Creation Test", _ => {
    it("Legit Create Mintable Token", async () => {
      const TokenFactoryInstance = await TokenFactory.deployed();
      const {
        symbol,
        name,
        decimals,
        logoURL,
        supply,
        logoHash,
        contractHash
      } = mintableToken;

      let token = await TokenFactoryInstance.getToken(symbol, { from: owner });
      assert.equal(
        token[5],
        0,
        "Must have no owner, since it does not exist yet"
      );
      await TokenFactoryInstance.createToken(
        name,
        symbol,
        decimals,
        logoURL,
        logoHash,
        supply,
        contractHash,
        { from: owner }
      );
      token = await TokenFactoryInstance.getToken(symbol, { from: owner });
      assert.equal(token[5], owner, "Must have owner who sent the transaction");
    });

    it("Legit Create Capped Token", async () => {
      const TokenFactoryInstance = await TokenFactory.deployed();
      const {
        symbol,
        name,
        decimals,
        logoURL,
        supply,
        logoHash,
        contractHash
      } = cappedToken;

      let token = await TokenFactoryInstance.getToken(symbol, { from: owner });
      assert.equal(
        token[5],
        0,
        "Must have no owner, since it does not exist yet"
      );
      await TokenFactoryInstance.createToken(
        name,
        symbol,
        decimals,
        logoURL,
        logoHash,
        supply,
        contractHash,
        { from: owner }
      );
      token = await TokenFactoryInstance.getToken(symbol, { from: owner });
      assert.equal(token[5], owner, "Must have owner who sent the transaction");
    });
  });

  describe("Transfer test", _ => {
    it("Legit transfer token (capped)", async () => {
      const TokenFactoryInstance = await TokenFactory.deployed();
      const tokenData = await TokenFactoryInstance.getToken(
        cappedToken.symbol,
        { from: owner }
      );
      const tokenInstance = await TokenTemplate.at(tokenData[0]);

      let ownerBalance = await tokenInstance.balanceOf(owner, { from: owner });
      let otherUserBalance = await tokenInstance.balanceOf(otherUser, {
        from: otherUser
      });

      assert.equal(
        ownerBalance,
        cappedToken.supply,
        "Owner should have token equal to total supply before the transfer"
      );
      assert.equal(
        otherUserBalance,
        0,
        "Other member should have token equal to 0 before the transfer"
      );

      await tokenInstance.transfer(otherUser, cappedToken.supply / 2, {
        from: owner
      });

      ownerBalance = await tokenInstance.balanceOf(owner, { from: owner });
      otherUserBalance = await tokenInstance.balanceOf(otherUser, {
        from: otherUser
      });

      assert.equal(
        ownerBalance,
        cappedToken.supply / 2,
        "Owner should have token equal to half total supply before the transfer"
      );
      assert.equal(
        otherUserBalance,
        cappedToken.supply / 2,
        "Other member should have token equal to half total supply before the transfer"
      );
    });
    it("Legit transfer token (mintable)", async () => {
      const TokenFactoryInstance = await TokenFactory.deployed();
      //await TokenFactoryInstance.mintToken(mintableToken.symbol, owner, cappedToken.supply, { from : owner })

      const tokenData = await TokenFactoryInstance.getToken(
        mintableToken.symbol,
        { from: owner }
      );
      const tokenInstance = await TokenTemplate.at(tokenData[0]);

      await tokenInstance.mint(owner, cappedToken.supply, { from: owner });

      let ownerBalance = await tokenInstance.balanceOf(owner, { from: owner });
      let otherUserBalance = await tokenInstance.balanceOf(otherUser, {
        from: otherUser
      });

      assert.equal(
        ownerBalance,
        cappedToken.supply,
        "Owner should have token equal to total supply before the transfer"
      );
      assert.equal(
        otherUserBalance,
        0,
        "Other member should have token equal to 0 before the transfer"
      );

      await tokenInstance.transfer(otherUser, cappedToken.supply / 2, {
        from: owner
      });

      ownerBalance = await tokenInstance.balanceOf(owner, { from: owner });
      otherUserBalance = await tokenInstance.balanceOf(otherUser, {
        from: otherUser
      });

      assert.equal(
        ownerBalance,
        cappedToken.supply / 2,
        "Owner should have token equal to half total supply after the transfer"
      );
      assert.equal(
        otherUserBalance,
        cappedToken.supply / 2,
        "Other member should have token equal to half total supply after the transfer"
      );
    });
  });
});
