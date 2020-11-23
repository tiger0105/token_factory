const CrowdsaleFactory = artifacts.require("./CrowdsaleFactoryTester.sol");
const TokenCrowdsale = artifacts.require("./TokenCrowdsaleTester.sol");
const TokenFactory = artifacts.require("./TokenFactory.sol");
const TokenTemplate = artifacts.require("./TokenTemplate.sol");

contract("CrowdsaleFactory", async accounts => {
  const token = {
    name: "Crowdsale Mintable Token",
    symbol: "CRMNT",
    decimals: 18,
    logoURL:
      "https://apollo-uploads-las.s3.amazonaws.com/1442324623/atlanta-hawks-logo-944556.png",
    logoHash: web3.utils.toHex(
      "0x4D021B157A49F472A48AB02A1F2F6E2986C169A7C78CC94179EDAEBD5E96E8E4"
    ), // sha256 hash
    contractHash: web3.utils.toHex(
      "0x4D021B157A49F472A48AB02A1F2F6E2986C169A7C78CC94179EDAEBD5E96E8E4"
    ),
    supply: 100000000000
  };

  const span = 864000;

  const crowdsale = {
    crowdsaleID: "test",
    acceptRatio: 1,
    giveRatio: 1,
    maxCap: 100
  };

  const owner = accounts[0];
  const otherUser = accounts[1];

  before(async function() {
    try {
      await web3.eth.personal.unlockAccount(owner, "", 10000);
      await web3.eth.personal.unlockAccount(otherUser, "", 10000);
    } catch (error) {
      console.warn(`error in unlocking wallet: ${JSON.stringify(error)}`);
    }
  });

  describe("Crowdsale Setup", _ => {
    it("Create Crowdsale", async () => {
      const CrowdsaleFactoryInstance = await CrowdsaleFactory.deployed();
      const TokenFactoryInstance = await TokenFactory.deployed();

      const {
        symbol,
        name,
        decimals,
        logoURL,
        supply,
        logoHash,
        contractHash
      } = token;
      const { crowdsaleID, acceptRatio, giveRatio, maxCap } = crowdsale;

      try {
        await TokenFactoryInstance.createToken(
          name + "1",
          symbol + "1",
          decimals,
          logoURL,
          logoHash,
          supply,
          contractHash,
          { from: owner }
        );
        await TokenFactoryInstance.createToken(
          name + "2",
          symbol + "2",
          decimals,
          logoURL,
          logoHash,
          supply,
          contractHash,
          { from: owner }
        );

        const tokenToGive = await TokenFactoryInstance.getToken(symbol + "1", {
          from: owner
        });
        const tokenToAccept = await TokenFactoryInstance.getToken(
          symbol + "2",
          { from: owner }
        );

        const start = Math.floor(new Date() / 1000);
        const end = Math.floor(new Date() / 1000 + 86400000); // for some reason it works

        await CrowdsaleFactoryInstance.createCrowdsale(
          crowdsaleID,
          tokenToGive[0],
          tokenToAccept[0],
          start,
          end,
          acceptRatio,
          giveRatio,
          maxCap,
          { from: owner }
        );

        const crowdsaleArr = await CrowdsaleFactoryInstance.getCrowdSale(
          crowdsaleID,
          { from: owner }
        );

        const tokenToAcceptInstance = await TokenTemplate.at(tokenToAccept[0]);
        await tokenToAcceptInstance.transfer(otherUser, supply / 2, {
          from: owner
        });

        const tokenToGiveInstance = await TokenTemplate.at(tokenToGive[0]);
        await tokenToGiveInstance.transfer(crowdsaleArr[0], supply, {
          from: owner
        });

        const crowdsaleInstance = await TokenCrowdsale.at(crowdsaleArr[0]);
        await crowdsaleInstance.unlockCrowdsale({ from: owner });

        const tokenToGiveTest = await TokenTemplate.at(crowdsaleArr[1], {
          from: owner
        });
        const symbolTest = await tokenToGiveTest.symbol({ from: owner });

        assert.equal(
          symbolTest,
          symbol + "1",
          "Expecting crowdsale to be created with correct tokenToGive symbol"
        );
      } catch (error) {
        assertFailError(error);
      }
    });
  });
  describe("Crowdsale Tests", () => {
    it("Time Travel", async () => {
      const { crowdsaleID } = crowdsale;

      try {
        const CrowdsaleFactoryInstance = await CrowdsaleFactory.deployed();

        const crowdsaleData = await CrowdsaleFactoryInstance.getCrowdSale(
          crowdsaleID,
          { from: owner }
        );

        const crowdsaleAddress = crowdsaleData[0];
        const crowdsaleInstance = await TokenCrowdsale.at(crowdsaleAddress, {
          from: owner
        });

        const startBN = await crowdsaleInstance.start({ from: owner });
        const start = await startBN.toNumber();

        await crowdsaleInstance.timeTravel(-span, { from: owner });

        let newStart = await crowdsaleInstance.start({ from: owner });
        assert.equal(
          newStart.toNumber(),
          start + span,
          "Should have got back in time (start must be forward)"
        );

        await crowdsaleInstance.timeTravel(span, { from: owner });

        newStart = await crowdsaleInstance.start({ from: owner });
        assert.equal(
          newStart.toNumber(),
          start,
          "Should have got forth in time (start must be backward)"
        );
      } catch (error) {
        assertFailError(error);
      }
    });
    it("Join not started Crowdsale", async () => {
      const { crowdsaleID } = crowdsale;
      const { symbol } = token;

      try {
        const CrowdsaleFactoryInstance = await CrowdsaleFactory.deployed();
        const TokenFactoryInstance = await TokenFactory.deployed();

        const crowdsaleData = await CrowdsaleFactoryInstance.getCrowdSale(
          crowdsaleID,
          { from: owner }
        );

        const crowdsaleAddress = crowdsaleData[0];
        const crowdsaleInstance = await TokenCrowdsale.at(crowdsaleAddress, {
          from: owner
        });

        await crowdsaleInstance.timeTravel(-span, { from: owner });

        const tokenToAccept = await TokenFactoryInstance.getToken(
          symbol + "2",
          {
            from: owner
          }
        );

        const tokenInstance = await TokenTemplate.at(tokenToAccept[0], {
          from: owner
        });

        await tokenInstance.approve(crowdsaleAddress, 100, { from: otherUser });

        try {
          await crowdsaleInstance.joinCrowdsale(crowdsaleID, { from: otherUser });
        } catch (_) {
          // success
        }

        await crowdsaleInstance.timeTravel(span, { from: owner });
      } catch (error) {
        assertFailError(error);
      }
    });
    it("Join ended Crowdsale", async () => {
      const { crowdsaleID } = crowdsale;
      const { symbol } = token;

      try {
        const CrowdsaleFactoryInstance = await CrowdsaleFactory.deployed();
        const TokenFactoryInstance = await TokenFactory.deployed();

        const crowdsaleData = await CrowdsaleFactoryInstance.getCrowdSale(
          crowdsaleID
        );

        const crowdsaleAddress = crowdsaleData[0];
        const crowdsaleInstance = await TokenCrowdsale.at(crowdsaleAddress);

        await crowdsaleInstance.timeTravel(span * 2);

        const tokenToGive = await TokenFactoryInstance.getToken(symbol + "2", {
          from: owner
        });

        const tokenInstance = await TokenTemplate.at(tokenToGive[0]);
        await tokenInstance.approve(crowdsaleAddress, 100, { from: otherUser });

        try {
          await crowdsaleInstance.joinCrowdsale(crowdsaleID, { from: otherUser });
        } catch (_) {
          // success
        }
        await crowdsaleInstance.timeTravel(span * -2);
      } catch (error) {
        assertFailError(error);
      }
    });
    it("Join crowdsale", async () => {
      const { crowdsaleID } = crowdsale;
      const { symbol } = token;

      try {
        const CrowdsaleFactoryInstance = await CrowdsaleFactory.deployed();
        const TokenFactoryInstance = await TokenFactory.deployed();

        const crowdsaleData = await CrowdsaleFactoryInstance.getCrowdSale(
          crowdsaleID,
          { from: owner }
        );

        const crowdsaleAddress = crowdsaleData[0];
        const crowdsaleInstance = await TokenCrowdsale.at(crowdsaleAddress);

        const tokenToAccept = await TokenFactoryInstance.getToken(
          symbol + "2",
          {
            from: owner
          }
        );

        const tokenInstance = await TokenTemplate.at(tokenToAccept[0]);
        await tokenInstance.approve(crowdsaleAddress, 1, { from: otherUser });

        await crowdsaleInstance.timeTravel(span / 2);

        await crowdsaleInstance.joinCrowdsale(1, { from : otherUser });

        const myReservation = await crowdsaleInstance.getMyReservation({ from: otherUser });
        assert.equal(myReservation.toNumber(), 1, "Must have 1 reservation");

        const reservationsData = await crowdsaleInstance.getReservationsData({from : owner});
        assert.equal(reservationsData[0][0].toString(), otherUser, "Owner must be in reservations data");
        assert.equal(reservationsData[1][0].toNumber(), 1, "Owner must have 1 reservation");

      } catch (error) {
        assertFailError(error);
      }
    });
    it("Refund me", async () => {
      const { crowdsaleID } = crowdsale;
      const { symbol } = token;

      try {
        const CrowdsaleFactoryInstance = await CrowdsaleFactory.deployed();
        const TokenFactoryInstance = await TokenFactory.deployed();

        const crowdsaleData = await CrowdsaleFactoryInstance.getCrowdSale(
          crowdsaleID,
          { from: owner }
        );

        const crowdsaleAddress = crowdsaleData[0];
        const crowdsaleInstance = await TokenCrowdsale.at(crowdsaleAddress);

        const tokenToAccept = await TokenFactoryInstance.getToken(
          symbol + "2", { from: owner }
        );

        const tokenInstance = await TokenTemplate.at(tokenToAccept[0]);

        let balBN = await tokenInstance.balanceOf(otherUser, { from: otherUser });
        const balBefore = balBN.toNumber();

        await crowdsaleInstance.refundMe(1, { from : otherUser });

        balBN = await tokenInstance.balanceOf(otherUser, { from: otherUser });
        const balAfter = balBN.toNumber();

        assert.equal(balAfter, balBefore + 1, "Must have refunded 1 coin");
      } catch (error) {
        assertFailError(error);
      }
    });
    it("Join crowdsale (over maxCap) and obtain reservation", async () => {
      const { crowdsaleID, maxCap } = crowdsale;
      const { symbol } = token;
      const extra = 10;

      try {
        const CrowdsaleFactoryInstance = await CrowdsaleFactory.deployed();
        const TokenFactoryInstance = await TokenFactory.deployed();

        const crowdsaleData = await CrowdsaleFactoryInstance.getCrowdSale(
          crowdsaleID,
          { from: owner }
        );

        const crowdsaleAddress = crowdsaleData[0];
        const crowdsaleInstance = await TokenCrowdsale.at(crowdsaleAddress);

        const tokenToGive = await TokenFactoryInstance.getToken(symbol + "1", {
          from: owner
        });
        const tokenToAccept = await TokenFactoryInstance.getToken(
          symbol + "2", { from: owner }
        );

        let tokenInstance = await TokenTemplate.at(tokenToAccept[0]);
        await tokenInstance.approve(crowdsaleAddress, 3 * maxCap, { from: otherUser });

        tokenInstance = await TokenTemplate.at(tokenToGive[0]);

        let balBN = await tokenInstance.balanceOf(otherUser, { from: otherUser });
        const balBefore = await balBN.toNumber();

        await crowdsaleInstance.joinCrowdsale(2 * maxCap + extra, { from: otherUser });

        balBN = await tokenInstance.balanceOf(otherUser, { from: otherUser });
        const balAfter = await balBN.toNumber();

        const myNewReservation = await crowdsaleInstance.getMyReservation({ from: otherUser });
        assert.equal(myNewReservation.toNumber(), 0, "Must have no reservation");

        assert.equal(balAfter, balBefore + maxCap, "Expecting to have added difference to maxCap after cap reached");
      } catch (error) {
        assertFailError(error);
      }
    });
  });
});

const assertFailError = error => {
  assert.fail(`
    Should always be able to call smart contracts, got instead this error:
    caused by: 
    ${error.stack}
    Reason: "${error.reason}"`);
};

const send = (method, params = []) =>
  web3.currentProvider.send(
    { id: 0, jsonrpc: "2.0", method, params },
    () => {}
  );

// timeTravel travels in the blockchain time.
const timeTravel = async seconds => {
  await send("evm_increaseTime", [seconds]);
  await send("evm_mine");
};
