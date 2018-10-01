import BubbleSortor from './bubble'

var algorithms = Object.freeze({"bubble": 1, "quick": 2, "merge": 3})

export default class SorterFactory {
    newSorter(algorithm) {
        if (algorithm === algorithms.bubble) {
            return new BubbleSortor()
        }
    }
}

export const SortingAlgorithms = algorithms
