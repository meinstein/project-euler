function main() {
  let a = 0
  let b = 0
  let c = 0

  for (let i = 1; i < 300; i++) {
    a = i * i
    for (let j = i + 1; j < i + 300; j++) {
      b = j * j

      for (let x = j + 1; x < j + 300; x++) {
        c = x * x

        if (a + b === c && i + j + x === 1000) {
          return i * j * x
        }
      }
    }
  }
}

const solution = main()
console.log(solution)
