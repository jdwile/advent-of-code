import * as solutions from './solutions';
import { NotImplementedError } from './solutions/errors';

export async function execute() {
    console.log(solutions);
    Object.keys(solutions).forEach(async (className) => {
        const sol = new solutions[className]();
        await sol.solve()
            .catch((error: Error) => {
                if (!(error instanceof NotImplementedError)) throw error;
                console.log(error.message);
            });
    })
}

execute();