import { readFileSync } from "fs";

const input = readFileSync("input.txt").toString().trim();
const regex = /mul\([0-9]+,[0-9]+\)/g;
const instructions = input.match(regex) || [];
let total = 0;

instructions.forEach((instruction) => {
  const a = parseInt(instruction.split(",")[0].split("(")[1]);
  const b = parseInt(instruction.split(",")[1].split(")")[0]);
  total += a * b;
});

console.log(total);
