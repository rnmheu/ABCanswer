//              //
//  🐶 WAO 🐶  //
//             //

const isTrial = true;
setTimeout(() => Deno.exit(0), isTrial ? 100000 : 2500);

type Answer = {
  count: number;
  coordinate: number[];
  isFinalAnswer: boolean;
};

//           //
//  UTILITY  //
//           //

async function readLine(): Promise<string | null> {
  const buf = new Uint8Array(1024);
  const n = await Deno.stdin.read(buf);
  return n !== null
    ? new TextDecoder().decode(buf.subarray(0, n)).trim()
    : null;
}

async function initLog() {
  await Deno.writeTextFile("./log/output.txt", "");
  await Deno.writeTextFile("./log/input.txt", "");
}

async function writeOutput(output: Answer) {
  const prefix = output.isFinalAnswer ? "a" : "q";
  const outputStr = `${prefix} ${output.count} ${
    output.coordinate.join(" ")
  }\n`;
  if (!isTrial) {
    console.log(outputStr);
    return;
  }

  await Deno.writeTextFile("./log/output.txt", outputStr, { append: true });

  const answer = hasOilField(output);
  await Deno.writeTextFile(
    "./log/input.txt",
    `${String(answer)}\n`,
    { append: true },
  );
  console.log(outputStr);
  return;
}

async function readInput() {
  if (!isTrial) {
    const read = await readLine();
    return parseInt(read!, 10);
  }
  const decoder = new TextDecoder();
  const readFile = decoder.decode(
    await Deno.readFile("./log/input.txt"),
  ).split("\n");
  return parseInt(readFile[readFile.length - 2], 10);
}

async function readOilMap(fieldSize: number) {
  const oilMap: number[][] = [];
  for (let i = 0; i < fieldSize; i++) {
    const line = (await readLine())!.split(" ");
    oilMap[i] = line.map((v) => parseInt(v, 10));
  }
  return oilMap;
}

function hasOilField(answer: Answer) {
  let findOilField = 0;
  for (let i = 0; i < answer.count * 2; i += 2) {
    findOilField += fields[answer.coordinate[i]][answer.coordinate[i + 1]];
  }
  return findOilField;
}

//                       //
//  📝 READ INFORMATION  //
//                       //

const line = (await readLine())!.split(" ");
const fieldSize = parseInt(line[0], 10);
const oilFieldCount = parseInt(line[1], 10);
let fields: number[][] = [];
if (isTrial) {
  fields = await readOilMap(fieldSize);
  await initLog();
} else {
  for (let i = 0; i < oilFieldCount; i++) {
    await readLine();
  }
}

//                        //
//  🔎 OIL FIELD SEARCH!  //
//                        //

const hasOil: number[][] = [];
const SPLIT_SIZE = 3;

const splitCount = Math.ceil(fieldSize / SPLIT_SIZE);

// TODO: ここを再帰的な処理に変えたい
for (let i = 0; i < splitCount; i++) {
  for (let j = 0; j < splitCount; j++) {
    // SPLIT_SIZE * SPLIT_SIZE の範囲を占うよ～
    const divided: number[] = [];

    for (let m = i * SPLIT_SIZE; m < i * SPLIT_SIZE + SPLIT_SIZE; m++) {
      for (let n = j * SPLIT_SIZE; n < j * SPLIT_SIZE + SPLIT_SIZE; n++) {
        if (m >= fieldSize || n >= fieldSize) {
          continue;
        }
        divided.push(m, n);
      }
    }

    await writeOutput({
      count: divided.length / 2,
      coordinate: divided,
      isFinalAnswer: false,
    });
    // PROBLEM: よく考えたら、ここで得られる油田の数は「近似値」であり…
    // それを使って探索した場合、結局全探索と変わらないケースが出てくる
    const oilFieldCount = await readInput();

    if (oilFieldCount === 0) {
      continue;
    }

    // 1 マスずつ掘る
    let findOilFieldInDivided = 0;
    for (let m = 0; m < divided.length; m += 2) {
      await writeOutput({
        count: 1,
        coordinate: [divided[m], divided[m + 1]],
        isFinalAnswer: false,
      });
      const resp = await readInput();

      if (resp !== 0) {
        findOilFieldInDivided++;
        hasOil.push([divided[m], divided[m + 1]]);
      }
      if (findOilFieldInDivided >= oilFieldCount!) {
        break;
      }
    }
  }
}

await writeOutput({
  count: hasOil.length,
  coordinate: hasOil.flat(),
  isFinalAnswer: true,
});

const resp: string = (await readLine())!;
if (resp !== "1") {
  throw new Error("Assertion failed");
}
Deno.exit(0);
