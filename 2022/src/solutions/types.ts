export interface SolveParams {
    part1?: boolean;
    part2?: boolean;
}

export abstract class Solution {
    abstract solve(params?: SolveParams): Promise<void>;
}