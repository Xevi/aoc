import { loadInputFile } from '../utils/fs.mjs'

const getData = (file) =>
  loadInputFile(file).map((entries) => {
    const [action, value] = entries.split(' ')
    return [action, Number(value)]
  })

const part1 = (file = 'input.txt') => {
  const data = getData(file)

  const { xPos, yPos } = data.reduce(
    (acc, [action, value]) => {
      switch (action) {
        case 'forward': {
          acc.xPos += value
          break
        }
        case 'up': {
          acc.yPos -= value
          break
        }
        case 'down': {
          acc.yPos += value
          break
        }
      }

      return acc
    },
    { xPos: 0, yPos: 0 }
  )

  return xPos * yPos
}

const part2 = (file = 'input.txt') => {
  const data = getData(file)

  const { xPos, yPos } = data.reduce(
    (acc, [action, value]) => {
      switch (action) {
        case 'forward': {
          acc.xPos += value
          acc.yPos += value * acc.aim
          break
        }
        case 'up': {
          acc.aim -= value
          break
        }
        case 'down': {
          acc.aim += value
          break
        }
      }

      return acc
    },
    { xPos: 0, yPos: 0, aim: 0 }
  )

  return xPos * yPos
}

console.log('Part 1:', part1('input.txt'))
console.log('Part 2:', part2('input.txt'))
