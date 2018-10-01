export default class Sorter {
    constructor() {
    }

    sort(array) {
        this.array = array
        this.commands = []
    }

    swap(ia, ib) {
        var temp = this.array[ia]
        this.array[ia] = this.array[ib]
        this.array[ib] = temp
        this.commands.push({a: ia, b: ib})
    }

    pass (ia, ib) {

    }
}