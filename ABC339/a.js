function main(input) {
    input = input.split('.');
    console.log(input[input.length - 1])
}

main(require('fs').readFileSync('/dev/stdin', 'utf8'));
