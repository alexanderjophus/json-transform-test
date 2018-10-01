var express = require('express');
var router = express.Router();
var faker = require('faker')

var idKey = ["id", "uuid"]
var fixtureJoiners = [" vs ", " VS ", " + ", "+", ""]
var sports = ["boxing", "golf", "blinking", "high fiving", ""]
var marketType = ["craziest-victory-dance", "fastest-sports-person", "slowest-sports-person", ""]

/* GET home page. */
router.get('/', function(req, res, next) {

  // possible QPs?
  var markets = 3

  var randomName1 = faker.name.findName();
  var randomName2 = faker.name.findName();

  var joiner = fixtureJoiners[Math.floor(Math.random()*fixtureJoiners.length)];
  var fixture = randomName1 + joiner + randomName2 

  var text = { 
    fixture: listify(fixture),
    startDate: faker.date.future(1),
    sport: listify(sports[Math.floor(Math.random()*sports.length)]),
    markets: Array.from({length: Math.random()*markets}, () => createMarket(randomName1, randomName2)),
  }

  var fixtureID = createId()
  text[idKey[Math.floor(Math.random()*idKey.length)]] = listify(fixtureID)

  res.setHeader('Content-Type', 'application/json');
  res.setHeader('Fixture-ID', fixtureID)
  res.end(JSON.stringify(text, null, 3));
});

module.exports = router;

function createMarket(n1, n2) {
  var options = [
    {
        id: createId(),
        name: n1,
        odds: createOdds()
    },
    {
        id: createId(),
        name: n2,
        odds: createOdds()
    }
  ]

  if (Math.random() < 0.5) {
    options.push({
      id: createId(),
      name: "Draw",
      odds: createOdds(),
    });
  }

  return {
    id: createId(),
    type: marketType[Math.floor(Math.random()*marketType.length)],
    options: options,
  }
}

function createId() {
  if (Math.random() < 0.5) {
    return Math.floor(Math.random()*1000)
  } else {
    return Math.floor(Math.random()*1000).toString()
  }
}

function createOdds() {
  var r = Math.random()
  if (r < 0.2) {
    return createFractionalOdds()
  } else if (r < 0.4) {
    return createDecimalOdds()
  } else if (r < 0.6) {
    return createMoneyLineOdds()
  } else {
    return
  }
}

function createFractionalOdds() {
  return Math.floor(Math.random()*10+1).toString() + "/" + Math.floor(Math.random()*10+1).toString()
}

function createDecimalOdds() {
  return (Math.floor(Math.random()*10)/5)+1
}

function createMoneyLineOdds() {
  if (Math.random() < 0.5) {
    return "+" + (Math.floor((Math.random()*5)+1)*50).toString()
  } else {
    return "-" + (Math.floor((Math.random()*5)+1)*50).toString()
  }
}

function listify(value) {
  if (Math.random() < 0.2) {
    return [value]
  }
  return value
}
