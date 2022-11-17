import * as solutions from './solutions';

export async function execute() {
    console.log(solutions);
    Object.keys(solutions).forEach(async (className) => {
        const sol = new solutions[className]();
        await sol.solve();
    })
}

execute();