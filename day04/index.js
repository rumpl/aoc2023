const { log } = require("console");
const fs = require("fs");

const data = fs.readFileSync("input.txt", "utf8");

const lines = data.split("\n");

let total = 0;
let wins = {};

for (let i = 0; i < lines.length; i++) {
  const line = lines[i];
  const parts = line.split(":");
  const game = parts[1];
  const pp = game.split("|");
  const a = pp[0];
  const b = pp[1];

  const win = a.trim().replaceAll("  ", " ").trim().split(" ");
  const mine = b.trim().replaceAll("  ", " ").trim().split(" ");

  wins[i] = { count: 1, wins: 0 };

  for (let j = 0; j < mine.length; j++) {
    m = mine[j];
    if (win.includes(m)) {
      wins[i].wins += 1;
    }
  }
}

for (p in wins) {
  for (i = 0; i < wins[p].wins; i++) {
    const n = i + parseInt(p) + 1;
    if (wins[n]) {
      wins[i + parseInt(p) + 1].count += wins[p].count;
    }
  }
}

for (p in wins) {
  total += wins[p].count;
}

console.log(total);
