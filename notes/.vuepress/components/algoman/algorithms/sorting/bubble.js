import Sorter from './sorter'
export default class BubbleSorter extends Sorter {
    constructor() {
        super()
    }

    sort(array) {
        this.array = array
        this.commands = []
        for (var i = 0; i < this.array.length; i++) {
            for (var j = i + 1; j < this.array.length; j++) {
                if (this.array[i] > this.array[j]) {
                    this.swap(i, j)
                }
            }
        }
    }
}
