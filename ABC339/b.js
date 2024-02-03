function main(input) {
    input = input.split(' ');
    const H = parseInt(input[0]);
    const W = parseInt(input[1]);
    const N = parseInt(input[2]);

    // 0 = "."
    // 1 = "#"
    // TOP -> RIGHT -> BOTTOM -> LEFT -> ...
    const DIRECTION = ["TOP", "RIGHT", "BOTTOM", "LEFT"]
    const torusGrid = Array.from(Array(H), () => new Array(W).fill(0));
    
    let count = 0;
    let currentPoint = [0, 0];
    let currentDirection = 0;

    while (count < N) {
        if (torusGrid[currentPoint[0]][currentPoint[1]] === 0) {
            torusGrid[currentPoint[0]][currentPoint[1]] = 1;
            
            currentDirection = direction("R", currentDirection);
            currentPoint = move(DIRECTION[currentDirection], currentPoint);
            count ++;
        } else {
            torusGrid[currentPoint[0]][currentPoint[1]] = 0;
            currentDirection = direction("L", currentDirection);
            currentPoint = move(DIRECTION[currentDirection], currentPoint);
            count ++
        }
    }

    oekaki(torusGrid)

    function move(direction, currentPoint) {
        switch (direction) {
            case "TOP":
                currentPoint[0] -= 1;
                if (currentPoint[0] < 0) {
                    currentPoint[0] = H - 1;
                }
                return currentPoint;
            case "RIGHT":
                currentPoint[1] += 1;
                if (currentPoint[1] >= W) {
                    currentPoint[1] = 0;
                }
                return currentPoint;
            case "BOTTOM":
                currentPoint[0] += 1;
                if (currentPoint[0] >= H) {
                    currentPoint[0] = 0;
                }
                return currentPoint;
            case "LEFT":
                currentPoint[1] -= 1;
                if (currentPoint[1] < 0) {
                    currentPoint[1] = W - 1;
                }
                return currentPoint;
        }
    }
    
    function direction(rotate, currentDirection) {
        if (currentDirection === 0 && rotate === "L") {
            return 3
        }
        if (currentDirection === 3 && rotate === "R") {
            return 0
        }
        if (rotate === "R") {
            return currentDirection + 1
        }
        return currentDirection - 1
    }
}

function oekaki(array) {
    for (let i = 0; i < array.length; i++) {
        for (let j = 0; j < array[i].length; j++) {
            if (array[i][j] === 0) {
                array[i][j] = "."
            } else {
                array[i][j] = "#"
            }
        }
        console.log(array[i].join(''));
    }
}

// main(require('fs').readFileSync('/dev/stdin', 'utf8'));
main("3 4 5");

