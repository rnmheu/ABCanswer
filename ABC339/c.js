function main(input) {
    input = input.split('\n');
    const N = parseInt(input[0]);
    const A = input[1].split(' ').map(x => parseInt(x));

    let minusCount = 0;
    let current = 0;

    for (let i = 0; i < N; i++) {
        current += A[i];
        if (current < 0) {
            minusCount += Math.abs(current);
            current = 0;
        }

    }
    console.log(current)
}

main(require('fs').readFileSync('/dev/stdin', 'utf8'));
// main("4\n-1 1000000000 1000000000 1000000000")