import fs from 'fs'
import path from 'path'
import os from 'os'

export const loadInputFile = (name, dir = process.cwd()) => {
  return fs
    .readFileSync(path.join(dir, name), { encoding: 'utf8' })
    .split(os.EOL)
}
