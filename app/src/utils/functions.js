
export function range(start, end) {
    return Array(end - start + 1).fill().map((_, idx) => start + idx)
}

export const deleteTile = (tiles, setTile) => {
    setTile(tiles.slice(0, -1))
}

export const clearTiles = (setTile, setResult) => {
    setTile([])
    setResult([])
}

export const checkPong = arr => arr.every( v => v === arr[0] )

export function checkChi(array) {
    var i = 2, d;
    while (i < array.length) {
        d = array[i - 1] - array[i - 2];
        if (Math.abs(d) === 1 && d === array[i] - array[i - 1]) {
            return true;
        }
        i++;
    }
    return false;
}

function onlyUnique(value, index, self) {
    return self.indexOf(value) === index;
}

export function checkPattern(arr) {
    // remove duplicates in arr
    let removed = arr.filter(onlyUnique)

    // check which pattern class
    let pattern1 = range(1,9)
    let pattern2 = range(10,18)
    let pattern3 = range(19, 27)
    let pattern4 = range(28, 31)
    let pattern5 = range(32,34)

    if(removed.every(r => pattern1.includes(r))){
        return 1
    }else if(removed.every(r => pattern2.includes(r))){
        return 2
    }else if(removed.every(r => pattern3.includes(r))){
        return 3
    }else if(removed.every(r => pattern4.includes(r))){
        return 4
    }else if(removed.every(r => pattern5.includes(r))){
        return 5
    }else{
        return 6
    }
}
