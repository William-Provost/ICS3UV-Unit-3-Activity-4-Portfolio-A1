/**
 * @author William Provost
 * @version 1.0.0
 * @date 2025-11-13
 * @fileoverview This program tells you if you should eat based on hunger level.
 */

// variables
let hungerAsString: string = "";
let hungerAsNumber: number = 0;

// input hunger level
hungerAsString = prompt(
  "Hello, how hungry are you on a scale of 1 to 10, where 1 is not hungry and 10 means you must eat?"
) || "0";
hungerAsNumber = parseInt(hungerAsString);

// check hunger level
if (hungerAsNumber > 5) {
  console.log("Please eat!");
} else {
  console.log("You are not really that hungry.");
}

console.log("\nDone.");
