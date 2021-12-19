import { loadInputFile } from "../utils/fs.mjs";

const getData = (file) =>
  loadInputFile(file).map((entries) => {
    return entries.split("");
  });

const getCommon = (data, i) => {
  const count = data.reduce((acc, num) => {
    const key = num[i];
    acc[key] = acc[key] || 0;
    acc[key]++;
    return acc;
  }, {});

  return {
    lessCommon: count["1"] >= count["0"] ? "0" : "1",
    mostCommon: count["1"] >= count["0"] ? "1" : "0",
  };
};

const getElementsFrom = (data, pos, el) =>
  data.filter((num) => {
    return num[pos] === el;
  });

const part1 = (file) => {
  const data = getData(file);

  const num = data[0];
  const gamma = [];
  const epsilon = [];

  for (let i = 0; i < num.length; i++) {
    const { lessCommon, mostCommon } = getCommon(data, i);
    gamma.push(mostCommon);
    epsilon.push(lessCommon);
  }

  return (
    Number.parseInt(gamma.join(""), 2) * Number.parseInt(epsilon.join(""), 2)
  );
};

const part2 = (file) => {
  const data = getData(file);
  let oxygenGenerationRating = data;
  let co2Scrubber = data;
  let pos = 0;

  while (oxygenGenerationRating.length > 1 || co2Scrubber.length > 1) {
    if (oxygenGenerationRating.length > 1)
      oxygenGenerationRating = getElementsFrom(
        oxygenGenerationRating,
        pos,
        getCommon(oxygenGenerationRating, pos).mostCommon
      );

    if (co2Scrubber.length > 1)
      co2Scrubber = getElementsFrom(
        co2Scrubber,
        pos,
        getCommon(co2Scrubber, pos).lessCommon
      );

    pos++;
  }

  return (
    Number.parseInt(oxygenGenerationRating[0].join(""), 2) *
    Number.parseInt(co2Scrubber[0].join(""), 2)
  );
};

console.log("Part 1", part1("input.txt"));
console.log("Part 2", part2("input.txt"));
