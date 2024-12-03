import { readFileSync } from "fs";

const input = readFileSync("input.txt").toString().trim();
const regex = /(mul\([0-9]+,[0-9]+\))|(do(n't)?\(\))/g;
const instructions = input.match(regex) || [];
let mulEnabled = true;
let total = 0;

instructions.forEach((rawInstruction) => {
  const instruction = rawInstruction.split("(")[0];
  switch (instruction) {
    case "mul":
      if (!mulEnabled) return;

      const a = parseInt(rawInstruction.split("(")[1].split(",")[0]);
      const b = parseInt(rawInstruction.split(",")[1].split(")")[0]);
      total += a * b;
      break;

    case "do":
      mulEnabled = true;
      break;

    case "don't":
      mulEnabled = false;
      break;
  }
});

console.log(total);
