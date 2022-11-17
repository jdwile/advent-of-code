#! /usr/bin/env node
const chalk = require('chalk');
const clear = require('clear');
const figlet = require('figlet');
const path = require('path');
const program = require('commander');
const fs = require('fs').promises;
const { existsSync } = require('fs');
require('dotenv').config({ path: require('find-config')('.env') })

async function run() {
    clear();
    console.log(
        chalk.red(
            figlet.textSync('aoc-cli', { horizontalLayout: 'full' })
        )
    );

    program
        .version('0.0.1')
        .description("CLI tool to generate Typescript project files for Advent of Code")
        .option('-g, --generate')
        .parse();

    const options = program.opts();

    if (options.generate) {
        const today = new Date().getDate();
        const dir = process.cwd();
        const yearStr = dir.split('\\')[dir.split('\\').length - 1];

        const inYearDirectory = yearStr.match(/\d\d\d\d/)?.length
        if (!inYearDirectory) {
            console.log('Not in a valid subdirectory - please run this in the directory for a year');
            return;
        }

        const year = parseInt(yearStr);

        await generateInputFiles(year, today);
        await generateBoilerplateSolutionFiles(today);
        await addSolutionFilesToExports(today);
    }
}

async function generateInputFiles(year: number, day: number) {
    const inputData = await getInputFileFromAOC(year, day);

    await fs.writeFile(`src/solutions/inputs/day${day}.txt`, inputData);

    console.log(`Created Day ${day} input file`);
}

async function generateBoilerplateSolutionFiles(day: number) {
    if (!await existsSync(`src/solutions/day${day}.ts`)) {
        const boilerplate = await fs.readFile('src/solutions/boilerplate/solution.txt', 'utf8');
        const adjustedBoilerplate = boilerplate.replace("DAY_PLACEHOLDER", `${day}`)
            .replace("CLASS_NAME_PLACEHOLDER", `Day${day}Solution`)

        await fs.writeFile(`src/solutions/day${day}.ts`, adjustedBoilerplate);

        console.log(`Created Day ${day} boilerplate solution file`);
    }
}

async function addSolutionFilesToExports(day: number) {
    if (await existsSync(`src/solutions/index.ts`)) {
        const currentContents: string = await fs.readFile('src/solutions/index.ts', 'utf8');
        const exportStatement = `export * from './day${day}';\n`;

        if (!currentContents.includes(exportStatement)) {
            const adjustedContents = currentContents + exportStatement;

            await fs.writeFile(`src/solutions/index.ts`, adjustedContents);

            console.log(`Added Day ${day} to solutions exports`);
        }
    }
}

async function getInputFileFromAOC(year: number, day: number): Promise<string> {
    const { got } = await import('got');

    const response = await got.get(`https://adventofcode.com/2021/day/${day}/input`, {
        headers: {
            Cookie: `session=${process.env.AOC_TOKEN}`
        }
    }).text();

    return response;
}

run();
