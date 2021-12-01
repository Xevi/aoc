import { loadInputFile } from '../utils/fs.mjs'

const data = loadInputFile('input.txt').map((el) => Number(el))

const part1 = (list = data) =>
  list.reduce(
    (count, el, i, array) =>
      el > (array[i - 1] || Infinity) ? count + 1 : count,
    0
  )

const part2 = () => {
  if (data.length < 3) return 0

  const threeWindows = []
  for (let i = 0; i < data.length - 2; i++) {
    threeWindows.push(data[i] + data[i + 1] + data[i + 2])
  }

  return part1(threeWindows)
}

console.log(part2())
